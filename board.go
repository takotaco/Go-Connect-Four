package main

import "fmt"

type Board struct {
	board []Stack
	rows  int
	cols  int
}

func NewBoard(rows int, cols int) (b Board) {
	b.rows = rows
	b.cols = cols

	for i := 0; i < cols; i++ {
		b.board = append(b.board, Stack{max: rows})
	}

	return b
}

func (b Board) Print() {
	for i := b.rows - 1; i >= 0; i-- {
		for j := 0; j < b.cols; j++ {
			items := b.board[j].count
			if items >= (i + 1) {
				fmt.Printf(b.board[j].items[i])
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func (b Board) makeMove(col int, piece string) (ok bool, err error) {
	ok, err = b.board[col].Push(piece)

	if !ok {
		return
	}

	return true, nil
}
