package tictactoe

func TicTacToeDemo() {
	player1 := NewPlayer(1, "Hari", X)
	player2 := NewPlayer(2, "Suve", O)
	game := NewGame(*player1, *player2)
	game.Start()
}
