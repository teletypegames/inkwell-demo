# pncdsl-demo — Morning Coffee

A short point-and-click adventure built on the
[`pncdsl`](https://git.teletypegames.org/games/pncdsl) framework. Two scenes,
three items, one NPC dialog and an ending cutscene — small on purpose,
designed as an end-to-end reference for someone learning the library.

Full walkthrough and file layout: [`DEMO.md`](DEMO.md).

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

## Controls

| Action | Effect |
|--------|--------|
| left click on a hotspot/UI button | run the active verb |
| left click on an inventory item   | select it (cursor picks it up) |
| left click on a hotspot with item selected | use-with |
| right click                       | deselect held item, or reset verb to Look |
| click during dialog               | advance line / pick choice |
| **F1**                            | toggle hotspot debug overlay |

## License

MIT — see [`LICENSE.md`](LICENSE.md). Copyright © 2026 Teletype Games.
