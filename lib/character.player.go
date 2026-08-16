package lib

import (
	"image/color"

	p "git.teletypegames.org/engines/inkwell"
)

var Player = p.Character{
	Name:        "player",
	Sprite:      "spr/player",
	Speed:       80,
	SpeechColor: color.RGBA{255, 230, 160, 255},
	W:           28,
	H:           62,
}
