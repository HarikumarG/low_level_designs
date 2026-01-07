package tictactoe

type Player struct {
	id     int
	name   string
	symbol Symbol
}

func NewPlayer(id int, name string, symbol Symbol) *Player {
	return &Player{
		id:     id,
		name:   name,
		symbol: symbol,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetSymbol() Symbol {
	return p.symbol
}
