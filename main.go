package main

import (
	"embed"

	"git.teletypegames.org/demos/inkwell-demo/lib"
	p "git.teletypegames.org/engines/inkwell"
)

// The asset tree is compiled into the binary so the js/wasm build can
// load it too — in the browser there is no OS filesystem to read from.
//
//go:embed assets
var assetsFS embed.FS

func main() {
	p.AssetFS = assetsFS
	// The demo is written in Hungarian and nothing else, so both the
	// spoken language and the fallback are "hu". Adding a second one is
	// adding a line to each Text — see the engine manual on languages.
	p.Lang = "hu"
	p.Fallback = "hu"
	lib.Run()
}
