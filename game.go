package main

import "fmt"

var activeP int
var pieces map[int]string
var board []Stack
var cols, rows int

func setup() {
	activeP = 1

	pieces = map[int]string{
		1: "X",
		2: "O",
	}

	cols = 7
	rows = 6

	board = makeBoard(rows, cols)
}

func makeBoard(rows int, cols int) (b []Stack) {
	for i := 0; i < cols; i++ {
		b = append(b, Stack{max: rows})
	}
	return
}

func printBoard(b []Stack) {
	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			items := b[j].count
			if items >= (i + 1) {
				fmt.Printf(b[j].items[i])
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func switchTurn() {
	if activeP == 1 {
		activeP = 2
	} else {
		activeP = 1
	}
	printTurn()
}

func printTurn() {
	fmt.Printf("It is player %v's turn \n", activeP)
}

func makeMove(b []Stack, c int, p int) {
	piece := pieces[p]
	ok, err := b[c].Push(piece)

	if !ok {
		fmt.Println(err)
	}

	printBoard(b)
}

func main() {
	setup()
	printBoard(board)
	printTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
	makeMove(board, 1, activeP)
	switchTurn()
}
