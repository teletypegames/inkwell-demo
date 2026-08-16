package lib

import p "git.teletypegames.org/engines/inkwell"

var Key = p.Item{
	Name:        "key",
	Sprite:      "spr/key",
	Description: "rozsdás kulcs",
	OnUseWith: map[string]p.Action{
		"cupboard": p.Seq(
			p.Say("player", "Klikk."),
			p.SetFlag(FlagCupboardOpen),
			p.Give("beans"),
			p.TakeAway("key"),
			p.Say("player", "Bab van benne. Pont, ami kell."),
		),
	},
}
