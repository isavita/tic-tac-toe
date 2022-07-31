package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	title        = "Tic Tac Toe"
	screenWidth  = 300
	screenHeight = 400
)

var (
	fontSize        int = 32
	mplusNormalFont font.Face
	bkg             = color.White
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bkg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
