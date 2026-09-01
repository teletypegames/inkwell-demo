package lib

import p "git.teletypegames.org/engines/inkwell"

var VictoryScript = p.Script{
	Name: "victory",
	Actions: p.Seq(
		p.Say("player", p.Text{{Lang: "hu", Text: "Készen is van!"}}),
		p.PlaySound("snd/coffee_done"),
		p.Wait(0.6),
		p.Say("cat", p.Text{{Lang: "hu", Text: "Mióóóóu. *boldog macska*"}}),
		p.Say("player", p.Text{{Lang: "hu", Text: "Reggeli kávé: megmentve."}}),
		p.SetVar("score", 100),
		p.SetVar("player.mood", "Boldog"),
		p.SetVar("cat.mood", "Boldog"),
		p.ShowEnd(p.Text{{Lang: "hu", Text: "Vége — kösz, hogy játszottál!"}}),
	),
}
