package tic_tac_toe

import "math/rand"

const (
	XPlayer        = 1
	OPlayer        = 2
	Draw           = 3
	DifficultyEasy = 101
	DifficultyHard = 102
	BoardSize      = 9
)

type GameState struct {
	board         [9]int
	currentPlayer int
	player        int
	difficulty    int
}

func (gs *GameState) Play(move int) bool {
	if gs.board[move] == 0 {
		gs.board[move] = gs.currentPlayer
		return true
	}

	return false
}

func (gs *GameState) NextTurn() {
	gs.currentPlayer = GetOponent(gs.currentPlayer)
}

func (gs *GameState) HasWinner() bool {
	for i := 0; i < 3; i++ {
		if (gs.board[i*3] != 0 && gs.board[i*3] == gs.board[i*3+1] && gs.board[i*3] == gs.board[i*3+2]) ||
			(gs.board[i] != 0 && gs.board[i] == gs.board[i+3] && gs.board[i] == gs.board[i+6]) {
			return true
		}
	}

	return (gs.board[0] != 0 && gs.board[0] == gs.board[4] && gs.board[0] == gs.board[8]) ||
		(gs.board[2] != 0 && gs.board[2] == gs.board[4] && gs.board[2] == gs.board[6])
}

func (gs *GameState) IsDraw() bool {
	for _, item := range gs.board {
		if item != XPlayer && item != OPlayer {
			return false
		}
	}

	return true
}

func GetOponent(player int) int {
	if player == XPlayer {
		return OPlayer
	}

	return XPlayer
}

func (gs *GameState) MakeMove() int {
	switch gs.difficulty {
	case DifficultyHard:
		return gs.findBestMove()
	default:
		return gs.findRandomMove()
	}
}

func (gs *GameState) findRandomMove() int {
	for i := 0; i < 100; i++ {
		n := rand.Intn(9)
		if gs.board[n] == 0 {
			return n
		}
	}

	return -1
}

func (gs *GameState) findBestMove() int {
	if gs.currentPlayer == OPlayer {
		return gs.bestMoveO()
	} else {
		return gs.bestMoveX()
	}
}

func (gs *GameState) bestMoveX() int {
	bestScore := -100
	bestMove := 0
	for i := 0; i < BoardSize; i++ {
		if gs.board[i] == XPlayer {
			temp := gs.board[i]
			gs.board[i] = XPlayer
			moveScore := gs.minimax(0, false)
			gs.board[i] = temp

			if moveScore > bestScore {
				bestScore = moveScore
				bestMove = i
			}
		}
	}

	return bestMove
}

func (gs *GameState) bestMoveO() int {
	bestScore := 100
	bestMove := 0

	for i := 0; i < BoardSize; i++ {
		if gs.board[i] < XPlayer {
			temp := gs.board[i]
			gs.board[i] = OPlayer
			moveScore := gs.minimax(0, true)
			gs.board[i] = temp

			if moveScore < bestScore {
				bestScore = moveScore
				bestMove = i
			}
		}
	}

	return bestMove
}

func (gs *GameState) evaluateBoard() int {
	for i := 0; i < 3; i++ {
		if gs.board[i*3] != 0 && gs.board[i*3] == gs.board[i*3+1] && gs.board[i*3] == gs.board[i*3+2] {
			if gs.board[i*3] == XPlayer {
				return 10
			} else if gs.board[i*3] == OPlayer {
				return -10
			}
		}

		if gs.board[i] != 0 && gs.board[i] == gs.board[i+3] && gs.board[i] == gs.board[i+6] {
			if gs.board[i] == XPlayer {
				return 10
			} else if gs.board[i] == OPlayer {
				return -10
			}
		}

		if (gs.board[4] != 0 && gs.board[0] == gs.board[4] && gs.board[4] == gs.board[8]) ||
			(gs.board[2] != 0 && gs.board[2] == gs.board[4] && gs.board[4] == gs.board[6]) {
			if gs.board[4] == XPlayer {
				return 10
			} else if gs.board[4] == OPlayer {
				return -10
			}
		}
	}

	return 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func (gs *GameState) isMoveLeft() bool {
	for _, v := range gs.board {
		if v != XPlayer && v != OPlayer {
			return true
		}
	}

	return false
}

func (gs *GameState) minimax(depth int, isMaximizing bool) int {
	score := gs.evaluateBoard()

	if score == 10 || score == -10 {
		return score
	}

	if !gs.isMoveLeft() {
		return 0
	}

	if isMaximizing {
		bestScore := -100

		for i := 0; i < BoardSize; i++ {
			// Check if cell is empty
			if gs.board[i] < XPlayer {
				temp := gs.board[i]
				// Make the move
				gs.board[i] = XPlayer

				// Call minimax recursively and choose
				// the maximum value
				bestScore = max(bestScore, gs.minimax(depth+1, false))

				// Undo the move
				gs.board[i] = temp
			}
		}

		return bestScore
	} else {
		bestScore := 100

		for i := 0; i < BoardSize; i++ {
			// Check if cell is empty
			if gs.board[i] < XPlayer {
				temp := gs.board[i]
				// Make the move
				gs.board[i] = OPlayer

				// Call minimax recursively and choose
				// the minimum value
				bestScore = min(bestScore, gs.minimax(depth+1, true))

				// Undo the move
				gs.board[i] = temp
			}
		}

		return bestScore
	}
}
