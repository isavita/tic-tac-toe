package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/isavita/tic-tac-toe/tic_tac_toe"
)

func main() {
	game, err := tic_tac_toe.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(tic_tac_toe.ScreenWidth*2, tic_tac_toe.ScreenHeight*2)
	ebiten.SetWindowTitle(tic_tac_toe.Title)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
