package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pieti/gofigure/game_of_life/gol"
)

func main() {
	g := gol.NewGame()

	ebiten.SetWindowSize(gol.ScreenWidth, gol.ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
