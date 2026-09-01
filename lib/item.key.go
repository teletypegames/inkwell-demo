package lib

import p "git.teletypegames.org/engines/inkwell"

var Key = p.Item{
	Name:        "key",
	Sprite:      "spr/key",
	Description: p.Text{{Lang: "hu", Text: "rozsdás kulcs"}},
	OnUseWith: map[string]p.Action{
		"cupboard": p.Seq(
			p.Say("player", p.Text{{Lang: "hu", Text: "Klikk."}}),
			p.SetFlag(FlagCupboardOpen),
			p.Give("beans"),
			p.TakeAway("key"),
			p.Say("player", p.Text{{Lang: "hu", Text: "Bab van benne. Pont, ami kell."}}),
		),
	},
}
