package main

import (
	"log"

	"git.teletypegames.org/games/pncdsl-demo/domain"
	"git.teletypegames.org/games/pncdsl"
)

func main() {
	if err := pncdsl.Run(domain.Build()); err != nil {
		log.Fatal(err)
	}
}
