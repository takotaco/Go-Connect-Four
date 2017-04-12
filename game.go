package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	p1      Player
	p2      Player
	activeP *Player
	board   Board
	r       *bufio.Reader
}

func newGame() (g Game) {
	g.p1 = Player{1, "X"}
	g.p2 = Player{2, "O"}
	g.activeP = &g.p1
	g.board = NewBoard(6, 7)
	g.r = bufio.NewReader(os.Stdin)
	return
}

func (g Game) switchTurn() {
	fmt.Println("player before switch", g.activeP.id)
	if g.activeP.id == g.p1.id {
		g.activeP = &g.p2
	} else {
		g.activeP = &g.p1
	}
	fmt.Println("player after switch", g.activeP.id)
}

func (g Game) printTurn() {
	fmt.Printf("It is player %v's turn \n", g.activeP.id)
}

func (g Game) takeTurn() {
	g.printTurn()
	fmt.Println("Which column do you want to play in?")
	sCol, _ := g.r.ReadString('\n')
	fmt.Println("column input: ", sCol)
	col, _ := strconv.Atoi(strings.TrimSpace(sCol))
	fmt.Printf("column chosen: %v, %T", col, col)
	if (col >= g.board.cols) || (col < 0) {
		fmt.Println("That's not a valid column.")
		g.takeTurn()
		return
	}

	ok, err := g.board.makeMove(col, g.activeP.piece)

	if !ok {
		fmt.Println(err)
		g.takeTurn()
		return
	}

	g.board.Print()
	g.switchTurn()
	return
}

func main() {
	g := newGame()
	g.board.Print()
	g.takeTurn()
	g.takeTurn()
}
