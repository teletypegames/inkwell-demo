package lib

import (
	"image/color"

	p "git.teletypegames.org/games/pncdsl"
)

// MorningCoffeeTheme is the game-specific palette: warm coffee browns
// over deep espresso black, paired with the GameChat colors below so the
// stand-alone chat panel reads as part of the same visual set.
var MorningCoffeeTheme = p.Theme{
	Name: "morning-coffee",

	PanelBG:                 color.RGBA{30, 22, 18, 255},
	StatusText:              color.RGBA{245, 230, 200, 255},
	FlashText:               color.RGBA{255, 200, 110, 255},
	VerbButtonBG:            color.RGBA{60, 40, 30, 255},
	VerbButtonSelectedBG:    color.RGBA{180, 110, 50, 255},
	VerbButtonText:          color.RGBA{255, 240, 215, 255},
	InventorySlotBG:         color.RGBA{45, 32, 25, 255},
	InventorySlotSelectedBG: color.RGBA{180, 110, 50, 255},
	SpeechBubbleBG:          color.RGBA{20, 14, 10, 215},
	SpeechDefaultText:       color.RGBA{255, 240, 215, 255},
	DialogBG:                color.RGBA{25, 18, 14, 245},
	DialogBorder:            color.RGBA{180, 110, 50, 255},
	DialogChoiceBG:          color.RGBA{45, 30, 22, 255},
	DialogChoiceHover:       color.RGBA{95, 60, 35, 255},
	DialogSpeaker:           color.RGBA{255, 200, 110, 255},
	DialogText:              color.RGBA{255, 240, 215, 255},
	EndCardBG:               color.RGBA{10, 6, 4, 235},
	EndCardText:             color.RGBA{255, 200, 110, 255},
	CursorColor:             color.RGBA{255, 230, 180, 255},
	HotspotOutline:          color.RGBA{255, 200, 110, 200},
	SceneBackdrop:           color.RGBA{20, 14, 10, 255},

	TopBarBG:     color.RGBA{18, 12, 8, 255},
	TopBarText:   color.RGBA{255, 230, 200, 255},
	TopBarAccent: color.RGBA{255, 200, 110, 255},

	// Built-in ChatLog colors stay defined for completeness; the demo
	// replaces the widget with GameChat, but other rich-UI widgets may
	// still index these slots.
	ChatLogBG:       color.RGBA{18, 12, 8, 235},
	ChatLogPrompt:   color.RGBA{255, 200, 110, 255},
	ChatLogResponse: color.RGBA{255, 240, 215, 255},
	ChatLogSystem:   color.RGBA{160, 130, 100, 255},

	CharacterPanelBG:     color.RGBA{15, 10, 6, 180},
	CharacterPanelBorder: color.RGBA{180, 110, 50, 255},
	CharacterPanelTitle:  color.RGBA{255, 200, 110, 255},
}

// Chat is the GameChat instance used by the demo. Its colors stay in sync
// with MorningCoffeeTheme by reusing the same warm-brown palette.
var Chat = &GameChat{
	Name:        "chat",
	Bounds:      p.Rect(2, 148, 280, 38),
	LineHeight:  9,
	Padding:     3,
	InputHeight: 13,

	BG:          color.RGBA{18, 12, 8, 235},
	Border:      color.RGBA{120, 80, 40, 255},
	FocusBorder: color.RGBA{255, 200, 110, 255},
	InputBG:     color.RGBA{28, 18, 12, 240},
	UserText:    color.RGBA{255, 220, 140, 255},
	BotText:     color.RGBA{220, 200, 170, 255},
	PromptText:  color.RGBA{255, 240, 215, 255},
	CursorColor: color.RGBA{255, 230, 180, 255},

	UserPrefix: "Te: ",
	BotPrefix:  "Cirmos: ",
}
