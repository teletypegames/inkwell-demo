package lib

import p "git.teletypegames.org/games/pncdsl"

// Assets lists every asset the game references. Files don't need to
// exist on disk — the library substitutes deterministic colored
// placeholders for anything missing.
var Assets = []p.Asset{
	{Name: "bg/bedroom", Path: "assets/bg/bedroom.png", Kind: p.AssetImage},
	{Name: "bg/kitchen", Path: "assets/bg/kitchen.png", Kind: p.AssetImage},
	{Name: "spr/key", Path: "assets/spr/key.png", Kind: p.AssetImage},
	{Name: "spr/beans", Path: "assets/spr/beans.png", Kind: p.AssetImage},
	{Name: "spr/mug", Path: "assets/spr/mug.png", Kind: p.AssetImage},
	{Name: "spr/player", Path: "assets/spr/player.png", Kind: p.AssetImage},
	{Name: "spr/cat", Path: "assets/spr/cat.png", Kind: p.AssetImage},

	{Name: "mus/wakeup", Path: "assets/mus/wakeup.ogg", Kind: p.AssetAudio},
	{Name: "mus/calm", Path: "assets/mus/calm.ogg", Kind: p.AssetAudio},
	{Name: "snd/coffee_done", Path: "assets/snd/coffee_done.ogg", Kind: p.AssetAudio},
}
