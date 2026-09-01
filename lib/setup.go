package lib

import (
	"log"

	p "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

// Run builds the game and hands it to the inkwell engine. It is the
// single entry point referenced from main.go.
func Run() {
	if err := p.Run(Build()); err != nil {
		log.Fatal(err)
	}
}

// Build wires every entity declared in the sibling files into a fresh
// *Game. Everything the game *is* goes in the p.Game literal; the
// entities it is made of are registered into the managers that literal
// carries. Validate() runs from p.Run() to cross-check name references
// between managers.
func Build() *p.Game {
	g := p.New(p.Game{
		Title:       "Reggeli Kávé",
		Width:       320,
		Height:      200,
		MaxLogLines: 64,

		StartStage:  "game",
		StartScene:  "bedroom",
		StartScript: "intro",
		StartTheme:  "morning-coffee",
	})

	g.ThemeManager.Register(MorningCoffeeTheme)

	for _, asset := range Assets {
		g.AssetManager.Register(asset)
	}

	g.CharacterManager.Register(Player)
	g.CharacterManager.Register(Cat)

	g.ItemManager.Register(Key)
	g.ItemManager.Register(Beans)
	g.ItemManager.Register(Mug)

	g.DialogManager.Register(CatMorning)

	g.ScriptManager.Register(IntroScript)
	g.ScriptManager.Register(VictoryScript)

	g.SceneManager.Register(Bedroom)
	g.SceneManager.Register(Kitchen)

	// One Stage, and it is the world. A game with a front end registers
	// more of them and opens on one that is not the world; this one goes
	// straight in, which is what StartStage naming the world Stage means.
	g.StageManager.Register(p.Stage{
		Name:  "game",
		World: true,
	})

	registerUI(g)

	// initial HUD vars
	g.State.SetVar("score", 0)
	g.State.SetVar("time", "Nap 1 - 07:23")
	g.State.SetVar("player.state", "Ébren")
	g.State.SetVar("player.mood", "Álmos")
	g.State.SetVar("cat.state", "Várakozik")
	g.State.SetVar("cat.mood", "Türelmetlen")

	return g
}

// registerUI puts up the story-rich HUD by hand rather than through
// RegisterRichUI. The preset would be shorter, but the demo wants three
// things it does not offer — its own chat panel instead of the built-in
// ChatLog, a right-click verb coin instead of the always-visible wheel
// (the wheel sits over the bedroom door and blocks that exit), and no
// player panel, because the avatar is already on screen. Registering
// what you want beats registering a preset and subtracting from it.
func registerUI(g *p.Game) {
	g.WidgetManager.Register(&p.TopBar{
		Name:     "topbar",
		Height:   12,
		ScoreVar: "score",
		ScoreMax: 100,
		TimeVar:  "time",
	})
	g.WidgetManager.Register(&p.HotspotDebug{
		Name: "hotspot_debug",
	})
	g.WidgetManager.Register(&p.CharacterPanel{
		Name:      "panel_cat",
		Bounds:    p.Rect(180, 14, 92, 36),
		Character: "cat",
		Title: p.Text{
			{Lang: "hu", Text: "MACSKA"},
		},
		Stats: []p.CharStat{
			{
				Label:  p.Text{{Lang: "hu", Text: "Állapot"}},
				VarKey: "cat.state",
			},
			{
				Label:  p.Text{{Lang: "hu", Text: "Kedv"}},
				VarKey: "cat.mood",
			},
		},
	})
	g.WidgetManager.Register(&p.InventoryBar{
		Name:     "inventory",
		Origin:   p.Point{X: 4, Y: 188},
		Slots:    8,
		Cols:     8,
		SlotSize: 10,
		Gap:      1,
	})
	g.WidgetManager.Register(&p.RadialVerbs{
		Name:        "verbs",
		Trigger:     p.MouseButtonRight,
		Radius:      36,
		HotspotOnly: true,
		Labels: map[string]p.Text{
			"look": {{Lang: "hu", Text: "Néz"}},
			"use":  {{Lang: "hu", Text: "Tedd"}},
			"talk": {{Lang: "hu", Text: "Szól"}},
			"take": {{Lang: "hu", Text: "Vedd"}},
		},
	})
	g.WidgetManager.Register(&p.SpeechBubble{
		Name:      "speech",
		MaxWidth:  200,
		Padding:   3,
		OffsetY:   4,
		FallbackY: 18,
	})
	g.WidgetManager.Register(&p.DialogBox{
		Name:       "dialog",
		Bounds:     p.Rect(0, 90, float64(g.Width), 56),
		LineHeight: 12,
		Padding:    6,
	})
	g.WidgetManager.Register(&p.EndCard{
		Name: "endcard",
	})
	g.WidgetManager.Register(&p.Cursor{
		Name: "cursor",
	})

	// The demo's own chat panel, in place of the built-in ChatLog: game
	// -independent chatter that does not belong in the library.
	g.WidgetManager.Register(Chat)

	// F5 saves, F9 loads. Both are a KeyBinding — a key and what it does —
	// so neither needs a widget of its own. The flash only fires when the
	// action before it succeeded, because a Seq stops at the first failure.
	g.WidgetManager.Register(&p.KeyBinding{
		Name: "save_key",
		Key:  ebiten.KeyF5,
		Do: p.Seq(
			p.SaveSlot(1),
			p.Flash(p.Text{{Lang: "hu", Text: "Mentve (F5)."}}),
		),
	})
	// F9 flashes nothing: a load replaces State wholesale on the next
	// frame, and the flash would be the first thing it wiped. The world
	// changing under the player is the feedback.
	g.WidgetManager.Register(&p.KeyBinding{
		Name: "load_key",
		Key:  ebiten.KeyF9,
		Do:   p.LoadSlot(1),
	})
}
