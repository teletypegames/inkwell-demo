package lib

import (
	"image/color"

	p "git.teletypegames.org/engines/inkwell"
)

var Cat = p.Character{
	Name:        "cat",
	Sprite:      "spr/cat",
	Speed:       40,
	SpeechColor: color.RGBA{200, 200, 255, 255},
	W:           34,
	H:           20,
}
