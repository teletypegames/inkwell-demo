package lib

import (
	"image/color"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	p "git.teletypegames.org/games/inkwell"
)

// GameChat is a stand-alone chat panel: click to focus, type a message,
// press Enter to send, and the widget answers with a random canned line.
// It does not talk to any game system — the in-game action log stays
// decoupled from this gameplay-flavoured chatter, so the widget can be
// reused across games without dragging the demo's content with it.
type GameChat struct {
	Name        string
	Bounds      p.Rectangle
	LineHeight  int
	Padding     int
	InputHeight int

	BG          color.Color
	Border      color.Color
	FocusBorder color.Color
	InputBG     color.Color
	UserText    color.Color
	BotText     color.Color
	PromptText  color.Color
	CursorColor color.Color

	UserPrefix string
	BotPrefix  string
	Replies    []string

	history []chatEntry
	buffer  []rune
	focused bool
	blink   float64
	scratch []rune
}

type chatEntry struct {
	text string
	self bool
}

func (c *GameChat) GetName() string { return c.Name }

func (c *GameChat) Tick(ctx *p.UICtx) {
	in := ctx.Game.Input
	c.blink += ctx.DT

	if in.LeftClicked() {
		if c.Bounds.Contains(in.Point()) {
			c.focused = true
			in.ConsumeLeft()
		} else {
			c.focused = false
		}
	}
	if !c.focused {
		return
	}

	c.scratch = ebiten.AppendInputChars(c.scratch[:0])
	if len(c.scratch) > 0 {
		c.buffer = append(c.buffer, c.scratch...)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(c.buffer) > 0 {
		c.buffer = c.buffer[:len(c.buffer)-1]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		c.focused = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
		text := strings.TrimSpace(string(c.buffer))
		c.buffer = c.buffer[:0]
		if text == "" {
			return
		}
		c.history = append(c.history, chatEntry{text: text, self: true})
		c.history = append(c.history, chatEntry{text: c.randomReply(), self: false})
	}
}

func (c *GameChat) randomReply() string {
	src := c.Replies
	if len(src) == 0 {
		src = defaultChatReplies
	}
	return src[rand.Intn(len(src))]
}

var defaultChatReplies = []string{
	"Hmm…",
	"Aha.",
	"Mióóóu.",
	"Komolyan?",
	"Talán.",
	"Igen, igen.",
	"Hagyj békén, álmos vagyok.",
	"Most főzöm a kávét.",
	"Persze, persze.",
}

func (c *GameChat) Draw(dst *ebiten.Image, ctx *p.UICtx) {
	g := ctx.Game
	b := c.Bounds
	if b.W <= 0 || b.H <= 0 {
		return
	}
	lh := c.LineHeight
	if lh <= 0 {
		lh = 10
	}
	pad := c.Padding
	if pad <= 0 {
		pad = 3
	}
	inputH := c.InputHeight
	if inputH <= 0 {
		inputH = lh + 4
	}
	userPrefix := c.UserPrefix
	if userPrefix == "" {
		userPrefix = "Te: "
	}
	botPrefix := c.BotPrefix
	if botPrefix == "" {
		botPrefix = "Cirmos: "
	}

	historyH := int(b.H) - inputH
	if historyH < lh {
		historyH = lh
	}

	vector.DrawFilledRect(dst, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), c.BG, false)

	wrapW := int(b.W) - pad*2
	type line struct {
		text string
		col  color.Color
	}
	var rendered []line
	for _, m := range c.history {
		col := c.BotText
		prefix := botPrefix
		if m.self {
			col = c.UserText
			prefix = userPrefix
		}
		for _, w := range chatWrap(prefix+m.text, wrapW) {
			rendered = append(rendered, line{text: w, col: col})
		}
	}
	maxLines := (historyH - pad*2) / lh
	if maxLines < 0 {
		maxLines = 0
	}
	start := len(rendered) - maxLines
	if start < 0 {
		start = 0
	}
	for i, ln := range rendered[start:] {
		y := int(b.Y) + pad + i*lh
		g.DrawText(dst, ln.text, int(b.X)+pad, y-1, ln.col)
	}

	inY := int(b.Y) + historyH
	vector.DrawFilledRect(dst, float32(b.X)+1, float32(inY)+1, float32(int(b.W)-2), float32(inputH-2), c.InputBG, false)
	txt := "> " + string(c.buffer)
	g.DrawText(dst, txt, int(b.X)+pad, inY+pad-1, c.PromptText)
	if c.focused && int(c.blink*2)%2 == 0 {
		cx := int(b.X) + pad + chatTextWidth(txt)
		cy := inY + pad - 1
		vector.DrawFilledRect(dst, float32(cx), float32(cy), 1, float32(lh-1), c.CursorColor, false)
	}

	border := c.Border
	if c.focused {
		border = c.FocusBorder
	}
	vector.StrokeRect(dst, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), 1, border, false)
	vector.DrawFilledRect(dst, float32(b.X), float32(inY), float32(b.W), 1, border, false)
}

// BlocksClickAt implements the inkwell clickBlocker contract so the
// verb-coin doesn't pop up over the chat panel.
func (c *GameChat) BlocksClickAt(pt p.Point) bool { return c.Bounds.Contains(pt) }

// chatTextWidth mirrors the inkwell text width estimate for the bundled
// debug font (~6 px / glyph). Kept private to the demo.
func chatTextWidth(s string) int { return len(s) * 6 }

func chatWrap(s string, maxPx int) []string {
	if maxPx <= 0 || chatTextWidth(s) <= maxPx {
		return []string{s}
	}
	var out []string
	var line string
	flush := func() {
		if line != "" {
			out = append(out, line)
			line = ""
		}
	}
	word := ""
	for _, r := range s {
		if r == ' ' || r == '\n' {
			if line == "" {
				line = word
			} else if chatTextWidth(line+" "+word) <= maxPx {
				line += " " + word
			} else {
				flush()
				line = word
			}
			word = ""
			if r == '\n' {
				flush()
			}
			continue
		}
		word += string(r)
	}
	if word != "" {
		if line == "" {
			line = word
		} else if chatTextWidth(line+" "+word) <= maxPx {
			line += " " + word
		} else {
			flush()
			line = word
		}
	}
	flush()
	return out
}
