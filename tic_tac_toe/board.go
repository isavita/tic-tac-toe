package tic_tac_toe

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gridSize      = 9
	gridLiniaSize = 1
	boardWidth    = 300
	boardHeight   = 300
	XSymbol       = 1
	OSymbol       = 2
	XSymbolSize   = 80
	OSymbolRadius = 40
)

var (
	XColor          = color.RGBA{64, 140, 242, 0xff}
	OColor          = color.RGBA{242, 140, 64, 0xff}
	boardFrameColor = color.Black
)

// Board represents a board state.
type Board struct {
	grid          [gridSize]int
	currentPlayer int
}

// Update updates the board state.
func (b *Board) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		index := posToIndex(x, y)
		if index >= 0 && b.grid[index] == 0 {
			b.grid[index] = b.currentPlayer
			if b.currentPlayer == XSymbol {
				b.currentPlayer = OSymbol
			} else {
				b.currentPlayer = XSymbol
			}
		}
	}
	return nil
}

// Layout implements ebiten.Game's Layout.
func (b *Board) Layout(outsideWidth, outsideHeight int) (boardWidth, boardHeight int) {
	return 300, 300
}

// NewBoard generates a new Board.
func NewBoard() (*Board, error) {
	b := &Board{
		grid:          [gridSize]int{},
		currentPlayer: XSymbol,
	}
	return b, nil
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Clear()
	for i := 0; i < gridSize; i++ {
		if b.grid[i] == XSymbol {
			posX, posY := indexToPos(i)
			b.drawX(boardImage, posX, posY, XSymbolSize, XColor, 4)
		} else if b.grid[i] == OSymbol {
			posX, posY := indexToPos(i)
			b.drawO(boardImage, posX, posY, OSymbolRadius, OColor, 5)
		}
	}

	b.drawBoard(boardImage, boardFrameColor)
}

func indexToPos(index int) (int, int) {
	i := index % 3
	x := i*100 + 50 + i*gridLiniaSize
	if index < 3 {
		return x, 50 + i*gridLiniaSize
	} else if index < 6 {
		return x, 150 + i*gridLiniaSize
	} else {
		return x, 250 + i*gridLiniaSize
	}
}

func posToIndex(x, y int) int {
	index := -1

	if x < 100 {
		index = 0
	} else if x < 200 {
		index = 1
	} else {
		index = 2
	}

	if y < 100 {
		index += 0
	} else if y < 200 {
		index += 3
	} else {
		index += 6
	}

	return index
}

func (b *Board) drawBoard(screen *ebiten.Image, clr color.Color) {
	step := boardWidth/3 + gridLiniaSize
	for i := step; i < boardWidth; i += step {
		for j := 0; j < boardHeight; j++ {
			for k := 0; k < gridLiniaSize; k++ {
				screen.Set(i+k, j+k, clr)
			}
		}
	}

	for i := 0; i < boardWidth; i++ {
		for j := step; j < boardHeight; j += step {
			screen.Set(i, j, clr)
		}
	}
}

func (b *Board) drawX(screen *ebiten.Image, x, y, size int, clr color.Color, thickness int) {
	x1 := x - size/2
	y1 := y - size/2
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j || j == size-i-1 {
				screen.Set(x1+i, y1+j, clr)
				if thickness > 1 {
					for thick := 1; thick < thickness; thick++ {
						if y1+j-thick > 0 {
							screen.Set(x1+i, y1+j-thick, clr)
						}
						if j+thick < size {
							screen.Set(x1+i, y1+j+thick, clr)
						}
					}
				}
			}
		}
	}
}

func (b *Board) drawO(screen *ebiten.Image, x, y, r int, clr color.Color, thickness int) {
	radius := float64(r)
	minAngle := math.Acos(1 - 1/radius)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius * math.Cos(angle)
		yDelta := radius * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		screen.Set(x1, y1, clr)
	}
	if thickness > 1 {
		b.drawO(screen, x, y, r-1, clr, thickness-1)
	}
}
