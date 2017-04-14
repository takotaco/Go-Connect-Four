package main

import (
	"bufio"
	"connectFour"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameIO struct {
	r *bufio.Reader
	g *connectFour.Game
}

func setup() GameIO {
	g := connectFour.NewGame()

	return GameIO{
		bufio.NewReader(os.Stdin),
		&g,
	}
}

func (i GameIO) printBoard() {
	fmt.Print(i.g.PrintBoard())
}

func (i GameIO) printTurn() {
	fmt.Printf("It is player %v's turn \n", i.g.GetActivePlayer().Piece())
}

func (i GameIO) printEndOfGame() {
	fmt.Println("Game Over")
	winningPlayer := i.g.GetWinningPlayer().Piece()
	if winningPlayer != "" {
		fmt.Printf("Player %v wins!\n", winningPlayer)
	} else {
		fmt.Println("It's a tie!")
	}
}

func (i GameIO) ReadUserInput() int {
	fmt.Println("Which column do you want to play in?")
	sCol, _ := i.r.ReadString('\n')
	col, _ := strconv.Atoi(strings.TrimSpace(sCol))
	return col - 1
}

func (i GameIO) GameLoop() {
	i.printTurn()
	col := i.ReadUserInput()
	ok, err := i.g.TakeTurn(col)

	if !ok {
		fmt.Println(err.Error())
	}

	i.printBoard()

	if i.g.IsGameOver() {
		i.printEndOfGame()
	}
}

func main() {
	i := setup()
	i.printBoard()
	for i.g.IsGameOver() != true {
		i.GameLoop()
	}
}
