package lib

import p "git.teletypegames.org/engines/inkwell"

var Bedroom = p.Scene{
	Name:       "bedroom",
	Title:      p.Text{{Lang: "hu", Text: "Hálószoba"}},
	Background: "bg/bedroom",
	Music:      "mus/wakeup",
	Actors: []p.SceneActor{
		{CharacterName: "player", At: p.Point{X: 160, Y: 145}},
	},
	// Floor strip — Walk targets get clipped to this band, so clicks
	// on the bed or nightstand send the player to the floor in front
	// of them instead of letting them clip through furniture.
	Walkboxes: []p.Polygon{
		p.PolyRect(10, 135, 300, 25),
	},
	Hotspots: []p.Hotspot{
		{
			Name:   "bed",
			Area:   p.Rect(20, 70, 90, 60),
			Label:  p.Text{{Lang: "hu", Text: "ágy"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "A párnám még meleg."}}),
			OnUse:  p.Say("player", p.Text{{Lang: "hu", Text: "Most keltem fel. Nem alszom vissza."}}),
		},
		{
			Name:  "nightstand",
			Area:  p.Rect(115, 80, 35, 50),
			Label: p.Text{{Lang: "hu", Text: "éjjeliszekrény"}},
			OnLook: p.If(p.Flag(FlagKeyTaken),
				p.Say("player", p.Text{{Lang: "hu", Text: "Üres a tetején."}}),
				p.Say("player", p.Text{{Lang: "hu", Text: "Egy rozsdás kulcs az éjjeliszekrényemen."}}),
			),
			OnTake: p.If(p.Flag(FlagKeyTaken),
				p.Say("player", p.Text{{Lang: "hu", Text: "Már elvittem."}}),
				p.Seq(
					p.Walk("player", p.Point{X: 130, Y: 145}),
					p.Give("key"),
					p.SetFlag(FlagKeyTaken),
					p.Say("player", p.Text{{Lang: "hu", Text: "Pont jól jöhet."}}),
				),
			),
			OnUse: p.If(p.Flag(FlagKeyTaken),
				p.Say("player", p.Text{{Lang: "hu", Text: "Már elvittem a kulcsot."}}),
				p.Seq(
					p.Walk("player", p.Point{X: 130, Y: 145}),
					p.Give("key"),
					p.SetFlag(FlagKeyTaken),
					p.Say("player", p.Text{{Lang: "hu", Text: "Megvan a kulcs."}}),
				),
			),
		},
		{
			Name:   "door",
			Area:   p.Rect(260, 50, 40, 90),
			Label:  p.Text{{Lang: "hu", Text: "ajtó"}},
			OnLook: p.Say("player", p.Text{{Lang: "hu", Text: "Az ajtó a konyhába vezet."}}),
			OnUse: p.Seq(
				p.Walk("player", p.Point{X: 260, Y: 145}),
				p.GoTo("kitchen"),
			),
		},
	},
}
