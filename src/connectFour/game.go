package connectfour

import "fmt"

type Game struct {
	p1       Player
	p2       Player
	activeP  *Player
	board    Board
	end      bool
	winningP *Player
}

func NewGame() Game {
	p1 := NewPlayer(1, "X")
	p2 := NewPlayer(2, "O")
	return Game{
		p1:       p1,
		p2:       p2,
		activeP:  &p1,
		board:    NewBoard(6, 7),
		end:      false,
		winningP: nil,
	}
}

func (g *Game) switchTurn() {
	if g.activeP.id == g.p1.id {
		g.activeP = &g.p2
	} else {
		g.activeP = &g.p1
	}
}

type ErrColOutOfRange struct {
	col int
}

func (e ErrColOutOfRange) Error() string {
	return fmt.Sprintf("%v is not a valid column", e.col)
}

func (g *Game) TakeTurn(col int) (ok bool, err error) {
	if (col >= g.board.cols) || (col < 0) {
		return false, ErrColOutOfRange{col}
	}

	ok, err, win := g.board.MakeMove(col, g.activeP.piece)

	if !ok {
		return ok, err
	}

	if win {
		g.winningP = g.activeP
	}
	g.end = g.board.IsFull() || win

	if g.end == false {
		g.switchTurn()
	}
	return true, nil
}

func (g Game) PrintBoard() string {
	return g.board.Print()
}

func (g Game) GetActivePlayer() Player {
	return *g.activeP
}

func (g Game) IsGameOver() bool {
	return g.end
}

func (g Game) GetWinningPlayer() Player {
	if g.winningP != nil {
		return *g.winningP
	} else {
		return Player{}
	}
}
