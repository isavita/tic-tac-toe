package tic_tac_toe

const (
	X    = 1
	O    = 2
	Draw = 3
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
