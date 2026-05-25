package domain

import p "git.teletypegames.org/games/pncdsl"

func defineMug(g *p.Game) {
	g.ItemManager.Register(p.Item{
		Name:        "mug",
		Sprite:      "spr/mug",
		Description: "kedvenc bögréje",
	})
}
