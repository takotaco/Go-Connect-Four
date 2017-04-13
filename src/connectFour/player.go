package connectFour

type Player struct {
	id    int
	piece string
}

func NewPlayer(id int, piece string) Player {
	return Player{
		id,
		piece,
	}
}

func (p Player) Piece() string {
	return p.piece
}

func (p Player) Id() int {
	return p.id
}
