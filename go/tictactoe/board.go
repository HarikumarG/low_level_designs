package tictactoe

import (
	"errors"
	"fmt"
)

const GRID_SIZE int = 3
const XWIN int = 3
const OWIN int = -3

type Symbol int

const (
	Empty Symbol = 0
	X     Symbol = 1
	O     Symbol = -1
)

type Board struct {
	matrix [GRID_SIZE][GRID_SIZE]Symbol
}

func NewBoard() *Board {
	board := &Board{
		matrix: [GRID_SIZE][GRID_SIZE]Symbol{},
	}

	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			board.matrix[i][j] = Empty
		}
	}

	return board
}

func (b *Board) MakeMove(i int, j int, symbol Symbol) (int, error) {
	isValid := b.isValidMove(i, j)

	if len(isValid) > 0 {
		return -1, errors.New(isValid)
	}

	b.matrix[i][j] = symbol

	if b.isWon(i, j) {
		return 1, nil
	}

	if b.isFull() {
		return 2, nil
	}

	return 0, nil
}

func (b *Board) Display() {
	fmt.Println("--- Current Board State ---")
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			switch b.matrix[i][j] {
			case X:
				fmt.Print("X ")
			case O:
				fmt.Print("O ")
			default:
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) isValidMove(i int, j int) string {
	if i >= GRID_SIZE || i < 0 || j >= GRID_SIZE || j < 0 {
		return "Invalid position"
	}
	if b.matrix[i][j] != Empty {
		return "Already position is filled"
	}
	return ""
}

func (b *Board) isFull() bool {
	isFull := true
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			if b.matrix[i][j] == Empty {
				isFull = false
			}
		}
	}
	return isFull
}

func (b *Board) isWon(i, j int) bool {
	sum := 0

	// check column
	for row := 0; row < GRID_SIZE; row++ {
		sum += int(b.matrix[row][j])
	}
	if sum == XWIN || sum == OWIN {
		fmt.Printf("Col match %v\n", sum)
		return true
	}

	// check row
	sum = 0
	for col := 0; col < GRID_SIZE; col++ {
		sum += int(b.matrix[i][col])
	}
	if sum == XWIN || sum == OWIN {
		fmt.Printf("Row match %v\n", sum)
		return true
	}

	// check main diagonal (top-left to bottom-right)
	if i == j {
		sum = 0
		for d := 0; d < GRID_SIZE; d++ {
			sum += int(b.matrix[d][d])
		}
		if sum == XWIN || sum == OWIN {
			fmt.Printf("Normal diagonal match %v\n", sum)
			return true
		}
	}

	// check anti-diagonal (top-right to bottom-left)
	if i+j == GRID_SIZE-1 {
		sum = 0
		for d := 0; d < GRID_SIZE; d++ {
			sum += int(b.matrix[d][GRID_SIZE-1-d])
		}
		if sum == XWIN || sum == OWIN {
			fmt.Printf("Inverted Normal diagonal match %v\n", sum)
			return true
		}
	}

	return false
}
