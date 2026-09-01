package lib

import p "git.teletypegames.org/engines/inkwell"

var Kitchen = p.Scene{
	Name:       "kitchen",
	Title:      p.Text{{Lang: "hu", Text: "Konyha"}},
	Background: "bg/kitchen",
	Music:      "mus/calm",
	Actors: []p.SceneActor{
		{CharacterName: "player", At: p.Point{X: 60, Y: 145}},
		{CharacterName: "cat", At: p.Point{X: 215, Y: 142}},
	},
	// Floor strip. Same shape as the bedroom — keeps movement on a
	// consistent Y band and clips destinations away from the cupboard
	// and coffee machine.
	Walkboxes: []p.Polygon{
		p.PolyRect(10, 135, 300, 25),
	},
	// First time the player walks into the kitchen with the key in
	// hand, drop a hint. The trigger arms on scene enter; the
	// rising-edge fires once per arming because Once: true.
	Triggers: []p.Trigger{
		{
			Name: "kitchen_with_key_hint",
			Once: true,
			When: p.And(p.HasItem("key"), p.Not(p.Flag(FlagCupboardOpen))),
			Do:   p.Say("player", p.Text{{Lang: "hu", Text: "A kulcs jól jöhet a szekrényhez."}}),
		},
	},
	Hotspots: []p.Hotspot{
		{
			Name:   "cupboard",
			Area:   p.Rect(30, 40, 80, 80),
			Label:  p.Text{{Lang: "hu", Text: "szekrény"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "Konyhaszekrény. Zárva."}}),
			OnUse: p.If(p.Flag(FlagCupboardOpen),
				p.Say("player", p.Text{{Lang: "hu", Text: "Már nyitva van."}}),
				p.If(p.HasItem("key"),
					p.Seq(
						p.Walk("player", p.Point{X: 70, Y: 145}),
						p.Say("player", p.Text{{Lang: "hu", Text: "Klikk."}}),
						p.SetFlag(FlagCupboardOpen),
						p.Give("beans"),
						p.TakeAway("key"),
					),
					p.Say("player", p.Text{{Lang: "hu", Text: "Zárva. Kéne egy kulcs."}}),
				),
			),
			OnUseWith: map[string]p.Action{
				"key": p.Seq(
					p.Walk("player", p.Point{X: 70, Y: 145}),
					p.Say("player", p.Text{{Lang: "hu", Text: "Klikk."}}),
					p.SetFlag(FlagCupboardOpen),
					p.Give("beans"),
					p.TakeAway("key"),
				),
			},
		},
		{
			Name:   "shelf",
			Area:   p.Rect(115, 45, 35, 50),
			Label:  p.Text{{Lang: "hu", Text: "polc"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "A polc tele bögrékkel."}}),
			OnTake: p.If(p.Flag(FlagMugTaken),
				p.Say("player", p.Text{{Lang: "hu", Text: "Tiszta már mind elfogyott."}}),
				p.Seq(
					p.Walk("player", p.Point{X: 125, Y: 145}),
					p.Give("mug"),
					p.SetFlag(FlagMugTaken),
					p.Say("player", p.Text{{Lang: "hu", Text: "Egy bögre, kérem."}}),
				),
			),
			OnUse: p.If(p.Flag(FlagMugTaken),
				p.Say("player", p.Text{{Lang: "hu", Text: "Már elvettem egyet."}}),
				p.Seq(
					p.Walk("player", p.Point{X: 125, Y: 145}),
					p.Give("mug"),
					p.SetFlag(FlagMugTaken),
					p.Say("player", p.Text{{Lang: "hu", Text: "Egy bögre, kérem."}}),
				),
			),
		},
		{
			Name:   "coffee_machine",
			Area:   p.Rect(160, 75, 50, 55),
			Label:  p.Text{{Lang: "hu", Text: "kávéfőző"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "A jó öreg kávéfőző."}}),
			OnUse: p.If(p.And(p.Flag(FlagCoffeeHasBeans), p.Flag(FlagCoffeeHasMug)),
				p.RunScript("victory"),
				p.Say("player", p.Text{{Lang: "hu", Text: "Még hiányzik valami. Bab? Bögre?"}}),
			),
			OnUseWith: map[string]p.Action{
				"beans": p.Seq(
					p.Walk("player", p.Point{X: 180, Y: 145}),
					p.SetFlag(FlagCoffeeHasBeans),
					p.TakeAway("beans"),
					p.Say("player", p.Text{{Lang: "hu", Text: "Bab a helyén."}}),
					p.If(p.And(p.Flag(FlagCoffeeHasBeans), p.Flag(FlagCoffeeHasMug)),
						p.RunScript("victory")),
				),
				"mug": p.Seq(
					p.Walk("player", p.Point{X: 180, Y: 145}),
					p.SetFlag(FlagCoffeeHasMug),
					p.TakeAway("mug"),
					p.Say("player", p.Text{{Lang: "hu", Text: "Bögre a helyén."}}),
					p.If(p.And(p.Flag(FlagCoffeeHasBeans), p.Flag(FlagCoffeeHasMug)),
						p.RunScript("victory")),
				),
			},
		},
		{
			Name:   "cat",
			Area:   p.Rect(195, 118, 35, 22),
			Label:  p.Text{{Lang: "hu", Text: "Cirmos"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "Cirmos, a házikedvencem."}}),
			OnTalk: p.RunDialog("cat_morning"),
		},
		{
			Name:   "door",
			Area:   p.Rect(0, 40, 22, 90),
			Label:  p.Text{{Lang: "hu", Text: "ajtó"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "Vissza a hálóba."}}),
			OnUse: p.Seq(
				p.Walk("player", p.Point{X: 20, Y: 145}),
				p.GoTo("bedroom"),
			),
		},
	},
}
