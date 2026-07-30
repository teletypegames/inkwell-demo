package lib

import p "git.teletypegames.org/games/inkwell"

var VictoryScript = p.Script{
	Name: "victory",
	Actions: p.Seq(
		p.Say("player", "Készen is van!"),
		p.PlaySound("snd/coffee_done"),
		p.Wait(0.6),
		p.Say("cat", "Mióóóóu. *boldog macska*"),
		p.Say("player", "Reggeli kávé: megmentve."),
		p.SetVar("score", 100),
		p.SetVar("player.mood", "Boldog"),
		p.SetVar("cat.mood", "Boldog"),
		p.ShowEnd("Vége — kösz, hogy játszottál!"),
	),
}
