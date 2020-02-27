package game

import (
	"fmt"
	"math/rand"
	"time"
)

type SlotValue int

const (
	Empty SlotValue = iota
	X
	O

	BOARDWIDTH  = 3
	BOARDHEIGHT = 3
)

func SlotValueToString(v SlotValue) string {
	switch {
	case v == Empty:
		return "0"
	case v == X:
		return "X"
	case v == O:
		return "O"
	}
	return ""
}

type TicTacToeGame struct {
	*TicTacToeBoard
	Player1            Player
	Player2            Player
	CurrentPlayersTurn Player
}

func NewTicTacToeGame(player1, player2 Player) *TicTacToeGame {
	return &TicTacToeGame{
		TicTacToeBoard: NewTicTacToeBoard(),
		Player1:        player1,
		Player2:        player2,
	}
}

func (t *TicTacToeGame) Run() {
	// Set the current players turn
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Float64() <= 0.5 {
		t.CurrentPlayersTurn = t.Player1
	} else {
		t.CurrentPlayersTurn = t.Player2
	}
	for !t.HasWinner() && !t.IsBoardFull() {
		fmt.Print(t.String())
		fmt.Println()
		t.CurrentPlayersTurn.MakeMove(t)
		t.SwapPlayers()
		fmt.Println()
	}
	fmt.Print(t.String())
	fmt.Println()
	if t.HasWinner() {
		fmt.Println("The game is over we have a winner!")
	}
	if t.IsBoardFull() {
		fmt.Println("The game board is full!")
	}
}

func (t *TicTacToeGame) HasWinner() bool {
	winningPositions := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	// Check for player 1 win
	for _, positions := range winningPositions {
		if t.Board[positions[0]] == t.Player1.Value() && t.Board[positions[1]] == t.Player1.Value() && t.Board[positions[2]] == t.Player1.Value() {
			return true
		}
	}
	// Check for player 2 win
	for _, positions := range winningPositions {
		if t.Board[positions[0]] == t.Player2.Value() && t.Board[positions[1]] == t.Player2.Value() && t.Board[positions[2]] == t.Player2.Value() {
			return true
		}
	}
	return false
}

func (t *TicTacToeGame) IsBoardFull() bool {
	for _, slot := range t.Board {
		if slot == Empty git{
			return false
		}
	}
	return true
}

func (t *TicTacToeGame) SwapPlayers() {
	if t.CurrentPlayersTurn == t.Player1 {
		t.CurrentPlayersTurn = t.Player2
	} else {
		t.CurrentPlayersTurn = t.Player1
	}
}

type TicTacToeBoard struct {
	Board []SlotValue
}

func NewTicTacToeBoard() *TicTacToeBoard {
	board := &TicTacToeBoard{
		Board: make([]SlotValue, BOARDHEIGHT*BOARDWIDTH),
	}
	for i := 0; i < BOARDHEIGHT*BOARDWIDTH; i++ {
		board.Board[i] = Empty
	}
	return board
}

func (t *TicTacToeBoard) MakeMove(slot int, slotValue SlotValue) bool {
	if slot >= 0 && slot < BOARDWIDTH*BOARDHEIGHT {
		if t.Board[slot] == Empty {
			t.Board[slot] = slotValue
			return true
		}
		return false
	}
	return false
}

func (t TicTacToeBoard) String() string {
	str := ""
	for i := 0; i < BOARDHEIGHT*BOARDWIDTH; i++ {
		str += SlotValueToString(t.Board[i])
		if (i+1)%3 == 0 {
			str += "\n"
		} else {
			str += " |"
		}
	}
	return str
}
