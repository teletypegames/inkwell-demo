package lib

import (
	"log"

	p "git.teletypegames.org/engines/inkwell"
)

// Run builds the game and hands it to the inkwell engine. It is the
// single entry point referenced from main.go.
func Run() {
	if err := p.Run(Build()); err != nil {
		log.Fatal(err)
	}
}

// Build wires every entity declared in the sibling files into a fresh
// *Game. Validate() runs from p.Run() to cross-check name references
// between managers.
func Build() *p.Game {
	g := p.NewGame("Reggeli Kávé", 320, 200)

	g.ThemeManager.Register(MorningCoffeeTheme)
	g.UseTheme("morning-coffee")

	for _, a := range Assets {
		g.AssetManager.Register(a)
	}

	g.CharacterManager.Register(Player)
	g.CharacterManager.Register(Cat)

	g.ItemManager.Register(Key)
	g.ItemManager.Register(Beans)
	g.ItemManager.Register(Mug)

	g.DialogueManager.Register(CatMorning)

	g.ScriptManager.Register(IntroScript)
	g.ScriptManager.Register(VictoryScript)

	g.SceneManager.Register(Bedroom)
	g.SceneManager.Register(Kitchen)

	// Story-rich HUD: top bar + NPC character panel + chat log + permanent
	// verb wheel + small inventory. The player panel is skipped — the
	// player avatar is always visible in the scene, so the floating card
	// in the corner just gets in the way. Pass "player" as the second
	// arg to put it back.
	p.RegisterRichUI(g, "", "cat")
	// Replace the framework's in-game ChatLog with the demo's own
	// click-and-type GameChat panel — game-independent chatter that
	// does not belong in the inkwell library.
	g.UIManager.Remove("chat")
	g.UIManager.Register(Chat)
	// Swap the rich UI's always-visible verb wheel for a right-click,
	// hotspot-only verb coin. The default wheel sits over the bedroom
	// door, blocking that exit — the coin only appears on a right-click
	// landing on a hotspot, so the screen stays free between picks.
	// Drop the cursor too so we can re-register it after the new wheel,
	// keeping the crosshair painted on top of the open coin.
	g.UIManager.Remove("verbs")
	g.UIManager.Remove("cursor")
	g.UIManager.Register(&p.RadialVerbs{
		Name:        "verbs",
		Trigger:     p.MouseButtonRight,
		Radius:      36,
		HotspotOnly: true,
		Labels: map[string]string{
			"look": "Néz",
			"use":  "Tedd",
			"talk": "Szól",
			"take": "Vedd",
		},
	})
	g.UIManager.Register(&p.Cursor{Name: "cursor"})
	g.UIManager.Register(&SaveKeys{})
	g.MaxLogLines = 64

	// initial HUD vars
	g.State.SetVar("score", 0)
	g.State.SetVar("time", "Nap 1 - 07:23")
	g.State.SetVar("player.state", "Ébren")
	g.State.SetVar("player.mood", "Álmos")
	g.State.SetVar("cat.state", "Várakozik")
	g.State.SetVar("cat.mood", "Türelmetlen")

	g.StartAt("bedroom").OnStart(p.RunScript("intro"))
	return g
}
