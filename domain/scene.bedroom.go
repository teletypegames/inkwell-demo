package domain

import p "git.teletypegames.org/games/pncdsl"

func defineBedroom(g *p.Game) {
	g.SceneManager.Register(p.Scene{
		Name:       "bedroom",
		Title:      "Hálószoba",
		Background: "bg/bedroom",
		Music:      "mus/wakeup",
		Actors: []p.SceneActor{
			{CharacterName: "player", At: p.Point{X: 160, Y: 145}},
		},
		// Floor strip — Walk targets get clipped to this band, so clicks
		// on the bed or nightstand send the player to the floor in front
		// of them instead of letting them clip through furniture.
		Walkboxes: []p.Polygon{
			p.Poly(
				p.Point{X: 10, Y: 135},
				p.Point{X: 310, Y: 135},
				p.Point{X: 310, Y: 160},
				p.Point{X: 10, Y: 160},
			),
		},
		Hotspots: []p.Hotspot{
			{
				Name:   "bed",
				Area:   p.Rect(20, 70, 90, 60),
				Label:  "ágy",
				OnLook: p.Say("player", "A párnám még meleg."),
				OnUse:  p.Say("player", "Most keltem fel. Nem alszom vissza."),
			},
			{
				Name:  "nightstand",
				Area:  p.Rect(115, 80, 35, 50),
				Label: "éjjeliszekrény",
				OnLook: p.If(p.Flag(FlagKeyTaken),
					p.Say("player", "Üres a tetején."),
					p.Say("player", "Egy rozsdás kulcs az éjjeliszekrényemen."),
				),
				OnTake: p.If(p.Flag(FlagKeyTaken),
					p.Say("player", "Már elvittem."),
					p.Seq(
						p.Walk("player", p.Point{X: 130, Y: 145}),
						p.Give("key"),
						p.SetFlag(FlagKeyTaken),
						p.Say("player", "Pont jól jöhet."),
					),
				),
				OnUse: p.If(p.Flag(FlagKeyTaken),
					p.Say("player", "Már elvittem a kulcsot."),
					p.Seq(
						p.Walk("player", p.Point{X: 130, Y: 145}),
						p.Give("key"),
						p.SetFlag(FlagKeyTaken),
						p.Say("player", "Megvan a kulcs."),
					),
				),
			},
			{
				Name:   "door",
				Area:   p.Rect(260, 50, 40, 90),
				Label:  "ajtó",
				OnLook: p.Say("player", "Az ajtó a konyhába vezet."),
				OnUse: p.Seq(
					p.Walk("player", p.Point{X: 260, Y: 145}),
					p.GoTo("kitchen"),
				),
			},
		},
	})
}
