package tic_tac_toe

import "math/rand"

const (
	X              = 1
	O              = 2
	Draw           = 3
	DifficultyEasy = 101
	DifficultyHard = 102
	BoardSize      = 9
)

type GameState struct {
	board         [9]int
	currentPlayer int
}

func (gs *GameState) Play(move int) bool {
	if gs.board[move] == 0 {
		gs.board[move] = gs.currentPlayer
		return true
	}

	return false
}

func (gs *GameState) NextTurn() {
	if gs.currentPlayer == X {
		gs.currentPlayer = O
	} else {
		gs.currentPlayer = X
	}
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
		if item != X && item != O {
			return false
		}
	}

	return true
}

func (gs *GameState) MakeMove(difficulty int) int {
	switch difficulty {
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
	if gs.currentPlayer == O {
		return gs.bestMoveO()
	} else {
		return gs.bestMoveX()
	}
}

func (gs *GameState) bestMoveX() int {
	bestScore := -100
	bestMove := 0
	for i := 0; i < BoardSize; i++ {
		if gs.board[i] == X {
			temp := gs.board[i]
			gs.board[i] = X
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
		if gs.board[i] < X {
			temp := gs.board[i]
			gs.board[i] = O
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
			if gs.board[i*3] == X {
				return 10
			} else if gs.board[i*3] == O {
				return -10
			}
		}

		if gs.board[i] != 0 && gs.board[i] == gs.board[i+3] && gs.board[i] == gs.board[i+6] {
			if gs.board[i] == X {
				return 10
			} else if gs.board[i] == O {
				return -10
			}
		}

		if (gs.board[4] != 0 && gs.board[0] == gs.board[4] && gs.board[4] == gs.board[8]) ||
			(gs.board[2] != 0 && gs.board[2] == gs.board[4] && gs.board[4] == gs.board[6]) {
			if gs.board[4] == X {
				return 10
			} else if gs.board[4] == O {
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
		if v != X && v != O {
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
			if gs.board[i] < X {
				temp := gs.board[i]
				// Make the move
				gs.board[i] = X

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
			if gs.board[i] < X {
				temp := gs.board[i]
				// Make the move
				gs.board[i] = O

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
