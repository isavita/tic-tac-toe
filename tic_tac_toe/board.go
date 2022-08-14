package tic_tac_toe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gridSize    = 9
	tileSize    = 96
	boarderSize = 3
)

// Board represents a board state.
type Board struct {
	grid [gridSize]int
}

// NewBoard generates a new Board.
func NewBoard() (*Board, error) {
	b := &Board{
		grid: [gridSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	return b, nil
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	rect := ebiten.NewImage(tileSize, tileSize)
	rect.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(float64(w*85+10), float64(h*85+10))
	boardImage.DrawImage(rect, op)
}
