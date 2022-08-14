package tic_tac_toe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Title           = "Tic Tac Toe"
	ScreenWidth     = 300
	ScreenHeight    = 300
	backgroundColor = color.White
)

// Game represents a game state.
type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard()
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(ScreenWidth, ScreenWidth)
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)

	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
