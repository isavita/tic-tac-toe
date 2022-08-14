package tic_tac_toe

import (
	"image/color"
	"math"

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

func (g *Game) drawBoard(screen *ebiten.Image, clr color.Color) {
	step := ScreenWidth / 3
	for i := step; i < ScreenWidth; i += step {
		for j := 0; j < ScreenHeight; j++ {
			screen.Set(i, j, clr)
		}
	}

	for i := 0; i < ScreenWidth; i++ {
		for j := step; j < ScreenHeight; j += step {
			screen.Set(i, j, clr)
		}
	}
}

func (g *Game) drawX(screen *ebiten.Image, x, y, size int, clr color.Color, thickness int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j || j == size-i-1 {
				screen.Set(x+i, y+j, clr)
				if thickness > 1 {
					for thick := 1; thick < thickness; thick++ {
						if y+j-thick > 0 {
							screen.Set(x+i, y+j-thick, clr)
						}
						if j+thick < size {
							screen.Set(x+i, y+j+thick, clr)
						}
					}
				}
			}
		}
	}
}

func (g *Game) drawO(screen *ebiten.Image, x, y, r int, clr color.Color, thickness int) {
	radius := float64(r)
	minAngle := math.Acos(1 - 1/radius)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius * math.Cos(angle)
		yDelta := radius * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		if thickness > 2 {
			g.drawO(screen, x, y, r-1, clr, thickness-1)
		}

		screen.Set(x1, y1, clr)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(ScreenWidth, ScreenWidth)
	}
	screen.Fill(backgroundColor)

	boardFrameColor := color.Black
	g.drawBoard(screen, boardFrameColor)

	xColor := color.RGBA{64, 140, 242, 100}
	oColor := color.RGBA{242, 140, 64, 100}
	g.drawX(screen, 10, 10, 80, xColor, 3)
	g.drawO(screen, 150, 50, 40, oColor, 3)
}
