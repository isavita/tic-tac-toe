package tic_tac_toe

import (
	"image/color"
	"log"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	gridSize      = 9
	gridLiniaSize = 1
	boardWidth    = 300
	boardHeight   = 300
	XSize         = 80
	ORadius       = 40
)

var (
	XColor          = color.RGBA{64, 140, 242, 0xff}
	OColor          = color.RGBA{242, 140, 64, 0xff}
	gameOverColor   = color.RGBA{165, 60, 30, 0xff}
	boardFrameColor = color.Black
	mplusBigFont    font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Board represents a board state.
type Board struct {
	gameState *GameState
	winner    int
}

// Update updates the board state.
func (b *Board) Update() error {
	if b.gameState.HasWinner() {
		b.winner = GetOponent(b.gameState.currentPlayer)
	} else if b.gameState.IsDraw() {
		b.winner = Draw
	} else if b.gameState.currentPlayer == b.gameState.player {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			move := posToIndex(x, y)
			if b.gameState.Play(move) {
				b.gameState.NextTurn()
			}
		}
	} else {
		b.gameState.Play(b.gameState.MakeMove())
		b.gameState.NextTurn()
	}

	return nil
}

// Layout implements ebiten.Game's Layout.
func (b *Board) Layout(outsideWidth, outsideHeight int) (boardWidth, boardHeight int) {
	return 300, 300
}

// NewBoard generates a new Board.
func NewBoard(player, difficulty int) (*Board, error) {
	b := &Board{
		gameState: &GameState{
			currentPlayer: XPlayer,
			board:         [gridSize]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			player:        player,
			difficulty:    difficulty,
		},
		winner: 0,
	}
	return b, nil
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {

	boardImage.Clear()
	for i := 0; i < gridSize; i++ {
		if b.gameState.board[i] == XPlayer {
			posX, posY := indexToPos(i)
			b.drawX(boardImage, posX, posY, XSize, XColor, 4)
		} else if b.gameState.board[i] == OPlayer {
			posX, posY := indexToPos(i)
			b.drawO(boardImage, posX, posY, ORadius, OColor, 5)
		}
	}
	b.drawBoard(boardImage, boardFrameColor)

	if b.winner == XPlayer {
		b.winnerXText(boardImage)
	} else if b.winner == OPlayer {
		b.winnerOText(boardImage)
	} else if b.winner == Draw {
		b.drawText(boardImage)
	}
}

func (b *Board) drawText(boardImage *ebiten.Image) {
	const x, y = 80, 160
	text.Draw(boardImage, "Draw!!!", mplusBigFont, x, y, gameOverColor)
}

func (b *Board) winnerXText(boardImage *ebiten.Image) {
	const x, y = 20, 160
	text.Draw(boardImage, "Player X wins!!!", mplusBigFont, x, y, gameOverColor)
}

func (b *Board) winnerOText(boardImage *ebiten.Image) {
	const x, y = 20, 160
	text.Draw(boardImage, "Player O wins!!!", mplusBigFont, x, y, gameOverColor)
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
