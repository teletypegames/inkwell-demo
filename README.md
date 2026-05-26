# pncdsl-demo — Morning Coffee

A short point-and-click adventure built on the
[`pncdsl`](https://git.teletypegames.org/games/pncdsl) framework. Two scenes,
three items, one NPC dialog and an ending cutscene — small on purpose,
designed as an end-to-end reference for someone learning the library.

## Running it

The library is fetched as a regular Go module. Prerequisites: Go 1.24+
and SSH access to `git.teletypegames.org` (the private host that serves the
`pncdsl` repo).

```bash
git clone ssh://git@git.teletypegames.org:2222/games/pncdsl-demo
cd pncdsl-demo
go mod tidy        # fetches git.teletypegames.org/games/pncdsl
go run .
```

A 1280×800 window opens.

### Pointing Go at the private host

If `go mod tidy` can't reach `git.teletypegames.org` over HTTPS, ask Go to
clone via SSH instead. One-time setup:

```bash
git config --global \
    url."ssh://git@git.teletypegames.org:2222/".insteadOf "https://git.teletypegames.org/"
export GOPRIVATE=git.teletypegames.org/*
```

(Add `GOPRIVATE=git.teletypegames.org/*` to your shell profile to make it
permanent.)

### Local-development override

While iterating on the framework and the demo in parallel, point Go at a
local checkout of `pncdsl` with a temporary `replace` directive:

```bash
go mod edit -replace=git.teletypegames.org/games/pncdsl=../pncdsl
go mod tidy
```

Remove the `replace` line before pushing.

## Project layout

Game content lives in `lib/` as package-level `var`s — one struct value
per file, no constructor functions. `main.go` only calls `lib.Run()`.

```
main.go                       func main() { lib.Run() }
lib/
  setup.go                    Run() and Build(); the only wiring code
  assets.go                   var Assets []pncdsl.Asset
  flags.go                    flag-name constants
  character.player.go         var Player
  character.cat.go            var Cat
  item.key.go / beans / mug   var Key, Beans, Mug
  scene.bedroom.go            var Bedroom
  scene.kitchen.go            var Kitchen
  script.intro.go             var IntroScript
  script.victory.go           var VictoryScript
  dialog.cat_morning.go       var CatMorning
  theme.go                    var MorningCoffeeTheme + var Chat
  widget.gamechat.go          GameChat — demo-only chat widget
  widget.savekeys.go          SaveKeys — F5/F9 save/load
```

Adding new content means creating one file with one `var`, and adding
one `Register(...)` line in `lib/setup.go`. No registrar functions, no
indirection.

## Custom theme and chat widget

The demo ships with a game-specific theme (`morning-coffee`, warm coffee
browns over deep espresso) registered alongside the framework's preset
themes and activated in `Build()`. It is registered the same way as any
other entity:

```go
g.ThemeManager.Register(MorningCoffeeTheme)
g.UseTheme("morning-coffee")
```

The chat strip at the bottom of the screen is **not** the framework's
built-in `ChatLog`. It is `GameChat`, a demo-only widget defined in
`lib/widget.gamechat.go`: click to focus (border highlights), type a
message, press Enter to send. The widget responds with a random canned
line — independent of the game state, just a playful HUD element. The
framework's `ChatLog` is removed and `GameChat` is registered in its
place during `Build()`.

`GameChat` is intentionally kept inside the demo. If another game wants
the same panel, copy the file in; the framework stays free of any
gameplay-flavoured chatter.

## Controls

| Action | Effect |
|--------|--------|
| left click on a hotspot/UI button | run the active verb |
| left click on an inventory item   | select it (cursor picks it up) |
| left click on a hotspot with item selected | use-with |
| right click on a hotspot          | open the verb coin |
| right click elsewhere             | deselect held item, or reset verb to Look |
| left click on the chat panel      | focus chat input |
| type + Enter (chat focused)       | send message, receive a random reply |
| Esc                               | unfocus chat |
| click during dialog               | advance line / pick choice |
| **F1**                            | toggle hotspot debug overlay |
| **F5** / **F9**                   | save / load slot 0 |

## License

MIT — see [`LICENSE.md`](LICENSE.md). Copyright © 2026 Teletype Games.
