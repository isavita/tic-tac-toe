package tic_tac_toe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Title        = "Tic Tac Toe"
	ScreenWidth  = 300
	ScreenHeight = 400
)

// Game represents a game state.
type Game struct {
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	return &Game{}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
}
