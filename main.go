package main

import (
	"embed"

	p "git.teletypegames.org/games/inkwell"
	"git.teletypegames.org/games/inkwell-demo/lib"
)

// The asset tree is compiled into the binary so the js/wasm build can
// load it too — in the browser there is no OS filesystem to read from.
//
//go:embed assets
var assetsFS embed.FS

func main() {
	p.AssetFS = assetsFS
	lib.Run()
}
