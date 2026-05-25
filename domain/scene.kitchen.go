package domain

import p "git.teletypegames.org/games/pncdsl"

func defineKitchen(g *p.Game) {
	g.SceneManager.Register(p.Scene{
		Name:       "kitchen",
		Title:      "Konyha",
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
			p.Poly(
				p.Point{X: 10, Y: 135},
				p.Point{X: 310, Y: 135},
				p.Point{X: 310, Y: 160},
				p.Point{X: 10, Y: 160},
			),
		},
		// First time the player walks into the kitchen with the key in
		// hand, drop a hint. The trigger arms on scene enter; the
		// rising-edge fires once per arming because Once: true.
		Triggers: []p.Trigger{
			{
				Name: "kitchen_with_key_hint",
				Once: true,
				When: p.And(p.HasItem("key"), p.Not(p.Flag(FlagCupboardOpen))),
				Do:   p.Say("player", "A kulcs jól jöhet a szekrényhez."),
			},
		},
		Hotspots: []p.Hotspot{
			{
				Name:   "cupboard",
				Area:   p.Rect(30, 40, 80, 80),
				Label:  "szekrény",
				OnLook: p.Say("player", "Konyhaszekrény. Zárva."),
				OnUse: p.If(p.Flag(FlagCupboardOpen),
					p.Say("player", "Már nyitva van."),
					p.If(p.HasItem("key"),
						p.Seq(
							p.Walk("player", p.Point{X: 70, Y: 145}),
							p.Say("player", "Klikk."),
							p.SetFlag(FlagCupboardOpen),
							p.Give("beans"),
							p.TakeAway("key"),
						),
						p.Say("player", "Zárva. Kéne egy kulcs."),
					),
				),
				OnUseWith: map[string]p.Action{
					"key": p.Seq(
						p.Walk("player", p.Point{X: 70, Y: 145}),
						p.Say("player", "Klikk."),
						p.SetFlag(FlagCupboardOpen),
						p.Give("beans"),
						p.TakeAway("key"),
					),
				},
			},
			{
				Name:   "shelf",
				Area:   p.Rect(115, 45, 35, 50),
				Label:  "polc",
				OnLook: p.Say("player", "A polc tele bögrékkel."),
				OnTake: p.If(p.Flag(FlagMugTaken),
					p.Say("player", "Tiszta már mind elfogyott."),
					p.Seq(
						p.Walk("player", p.Point{X: 125, Y: 145}),
						p.Give("mug"),
						p.SetFlag(FlagMugTaken),
						p.Say("player", "Egy bögre, kérem."),
					),
				),
				OnUse: p.If(p.Flag(FlagMugTaken),
					p.Say("player", "Már elvettem egyet."),
					p.Seq(
						p.Walk("player", p.Point{X: 125, Y: 145}),
						p.Give("mug"),
						p.SetFlag(FlagMugTaken),
						p.Say("player", "Egy bögre, kérem."),
					),
				),
			},
			{
				Name:   "coffee_machine",
				Area:   p.Rect(160, 75, 50, 55),
				Label:  "kávéfőző",
				OnLook: p.Say("player", "A jó öreg kávéfőző."),
				OnUse: p.If(p.And(p.Flag(FlagCoffeeHasBeans), p.Flag(FlagCoffeeHasMug)),
					p.RunScript("victory"),
					p.Say("player", "Még hiányzik valami. Bab? Bögre?"),
				),
				OnUseWith: map[string]p.Action{
					"beans": p.Seq(
						p.Walk("player", p.Point{X: 180, Y: 145}),
						p.SetFlag(FlagCoffeeHasBeans),
						p.TakeAway("beans"),
						p.Say("player", "Bab a helyén."),
						p.If(p.And(p.Flag(FlagCoffeeHasBeans), p.Flag(FlagCoffeeHasMug)),
							p.RunScript("victory")),
					),
					"mug": p.Seq(
						p.Walk("player", p.Point{X: 180, Y: 145}),
						p.SetFlag(FlagCoffeeHasMug),
						p.TakeAway("mug"),
						p.Say("player", "Bögre a helyén."),
						p.If(p.And(p.Flag(FlagCoffeeHasBeans), p.Flag(FlagCoffeeHasMug)),
							p.RunScript("victory")),
					),
				},
			},
			{
				Name:   "cat",
				Area:   p.Rect(195, 118, 35, 22),
				Label:  "Cirmos",
				OnLook: p.Say("player", "Cirmos, a házikedvencem."),
				OnTalk: p.RunDialogue("cat_morning"),
			},
			{
				Name:   "door",
				Area:   p.Rect(0, 40, 22, 90),
				Label:  "ajtó",
				OnLook: p.Say("player", "Vissza a hálóba."),
				OnUse: p.Seq(
					p.Walk("player", p.Point{X: 20, Y: 145}),
					p.GoTo("bedroom"),
				),
			},
		},
	})
}
