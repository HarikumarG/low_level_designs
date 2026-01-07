package tictactoe

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	player1       Player
	player2       Player
	board         Board
	currentPlayer Player
}

func NewGame(p1 Player, p2 Player) *Game {
	return &Game{
		player1:       p1,
		player2:       p2,
		board:         *NewBoard(),
		currentPlayer: p1,
	}
}

func (g *Game) Start() {
	for {
		g.board.Display()
		i, j, err := g.getPosition()
		if err != nil {
			fmt.Println(err)
			continue
		}

		status, err := g.board.MakeMove(i, j, g.currentPlayer.GetSymbol())

		if err != nil {
			fmt.Println(err)
			fmt.Printf("Replay for %v\n", g.currentPlayer.name)
			continue
		}

		if status == 1 {
			fmt.Printf("%v Won\n", g.currentPlayer.GetName())
			g.board.Display()
			break
		}
		if status == 2 {
			fmt.Printf("Game Draw\n")
			g.board.Display()
			break
		}

		g.switchPlayer()
	}
	fmt.Println("Game Ends....")
}

func (g *Game) getPosition() (int, int, error) {
	fmt.Printf("Enter position to Play for the player %v\n", g.currentPlayer.name)
	var position string
	fmt.Scan(&position)
	splitStrings := strings.Split(position, ",")
	pos1, err1 := strconv.Atoi(splitStrings[0])
	if err1 != nil {
		return -1, -1, err1
	}
	pos2, err2 := strconv.Atoi(splitStrings[1])
	if err2 != nil {
		return -1, -1, err2
	}
	return pos1, pos2, nil
}

func (g *Game) switchPlayer() {
	if g.currentPlayer.id == g.player1.id {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}
