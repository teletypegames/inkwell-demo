package lib

import p "git.teletypegames.org/engines/inkwell"

var IntroScript = p.Script{
	Name: "intro",
	Actions: p.Seq(
		p.Wait(0.4),
		p.Say("player", p.Text{{Lang: "hu", Text: "Brr, hideg van."}}),
		p.Say("player", p.Text{{Lang: "hu", Text: "Egy kávé kéne, mielőtt szétszakad a fejem."}}),
		p.Walk("player", p.Point{X: 200, Y: 145}),
		p.SetVar("player.mood", "Eltökélt"),
	),
}
