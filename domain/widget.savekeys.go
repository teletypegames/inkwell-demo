package domain

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	p "git.teletypegames.org/games/pncdsl"
)

// saveKeys is a no-draw widget that binds F5 → Save(slot 0) and
// F9 → Load(slot 0). Errors are surfaced via FlashLine so the player
// sees them in the status strip.
type saveKeys struct{}

func (s *saveKeys) GetName() string { return "save_keys" }

func (s *saveKeys) Tick(ctx *p.UICtx) {
	g := ctx.Game
	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		if err := g.Save(0); err != nil {
			g.FlashLine("Mentés hiba: " + err.Error())
		} else {
			g.FlashLine("Mentve (F5).")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF9) {
		if err := g.Load(0); err != nil {
			g.FlashLine("Töltés hiba: " + err.Error())
		} else {
			g.FlashLine("Betöltve (F9).")
		}
	}
}

func (s *saveKeys) Draw(_ *ebiten.Image, _ *p.UICtx) {}
