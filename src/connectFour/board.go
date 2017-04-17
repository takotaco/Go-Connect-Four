package connectfour

import (
	"fmt"
	"stack"
)

type Board struct {
	board []stack.Stack
	rows  int
	cols  int
}

func NewBoard(rows int, cols int) (b Board) {
	b.rows = rows
	b.cols = cols

	for i := 0; i < cols; i++ {
		b.board = append(b.board, stack.New(rows))
	}

	return b
}

func (b Board) Print() (s string) {
	for i := 0; i < b.cols; i++ {
		s += fmt.Sprintf("|%v", i+1)
	}
	s += fmt.Sprintln("|")

	for i := b.rows - 1; i >= 0; i-- {
		for j := 0; j < b.cols; j++ {
			items := b.board[j].Count()
			if items >= (i + 1) {
				s += fmt.Sprintf("|%v", b.board[j].GetValueAt(i))
			} else {
				s += fmt.Sprint("|_")
			}
		}
		s += fmt.Sprintln("|")
	}
	return
}

func (b *Board) MakeMove(col int, piece string) (ok bool, err error, win bool) {
	ok, err = b.board[col].Push(piece)

	if !ok {
		return
	}

	row := b.board[col].Count()
	win = b.checkMove(col, row-1, piece)

	return true, nil, win
}

func (b Board) checkMove(col int, row int, piece string) (win bool) {
	if b.checkMoveHelper(col, row, piece) >= 3 {
		win = true
	}
	return
}

func (b Board) checkMoveHelper(col int, row int, piece string) (inARow int) {
	// check vertical
	inARow = b.checkUp(col, row, piece) + b.checkDown(col, row, piece)
	if inARow >= 3 {
		return
	}
	// check horizontal
	inARow = b.checkLeft(col, row, piece) + b.checkRight(col, row, piece)
	if inARow >= 3 {
		return
	}
	// check diagonal left up to right down
	inARow = b.checkUpLeft(col, row, piece) + b.checkDownRight(col, row, piece)
	if inARow >= 3 {
		return
	}
	// check diagonal left down to right up
	inARow = b.checkDownLeft(col, row, piece) + b.checkUpRight(col, row, piece)
	if inARow >= 3 {
		return
	}

	return
}

func (b Board) checkUp(col int, row int, piece string) (inARow int) {
	nextRow := row + 1
	nextPiece := b.board[col].GetValueAt(nextRow)
	if nextPiece == piece {
		return 1 + b.checkUp(col, nextRow, piece)
	} else {
		return 0
	}
}

func (b Board) checkDown(col int, row int, piece string) (inARow int) {
	nextRow := row - 1
	nextPiece := b.board[col].GetValueAt(nextRow)
	if nextPiece == piece {
		return 1 + b.checkDown(col, nextRow, piece)
	} else {
		return 0
	}
}

func (b Board) checkLeft(col int, row int, piece string) (inARow int) {
	nextCol := col - 1
	if nextCol < 0 || nextCol >= b.cols {
		return 0
	}
	nextPiece := b.board[nextCol].GetValueAt(row)
	if nextPiece == piece {
		return 1 + b.checkLeft(nextCol, row, piece)
	} else {
		return 0
	}
}

func (b Board) checkRight(col int, row int, piece string) (inARow int) {
	nextCol := col + 1
	if nextCol < 0 || nextCol >= b.cols {
		return 0
	}
	nextPiece := b.board[nextCol].GetValueAt(row)
	if nextPiece == piece {
		return 1 + b.checkRight(nextCol, row, piece)
	} else {
		return 0
	}
}

func (b Board) checkUpLeft(col int, row int, piece string) (inARow int) {
	nextCol := col - 1
	nextRow := row + 1
	if nextCol < 0 || nextCol >= b.cols {
		return 0
	}
	nextPiece := b.board[nextCol].GetValueAt(nextRow)
	if nextPiece == piece {
		return 1 + b.checkUpLeft(nextCol, nextRow, piece)
	} else {
		return 0
	}
}

func (b Board) checkDownLeft(col int, row int, piece string) (inARow int) {
	nextCol := col - 1
	nextRow := row - 1
	if nextCol < 0 || nextCol >= b.cols {
		return 0
	}
	nextPiece := b.board[nextCol].GetValueAt(nextRow)
	if nextPiece == piece {
		return 1 + b.checkDownLeft(nextCol, nextRow, piece)
	} else {
		return 0
	}
}

func (b Board) checkUpRight(col int, row int, piece string) (inARow int) {
	nextCol := col + 1
	nextRow := row + 1
	if nextCol < 0 || nextCol >= b.cols {
		return 0
	}
	nextPiece := b.board[nextCol].GetValueAt(nextRow)
	if nextPiece == piece {
		return 1 + b.checkUpRight(nextCol, nextRow, piece)
	} else {
		return 0
	}
}

func (b Board) checkDownRight(col int, row int, piece string) (inARow int) {
	nextCol := col + 1
	nextRow := row - 1
	if nextCol < 0 || nextCol >= b.cols {
		return 0
	}
	nextPiece := b.board[nextCol].GetValueAt(nextRow)
	if nextPiece == piece {
		return 1 + b.checkDownRight(nextCol, nextRow, piece)
	} else {
		return 0
	}
}

func (b Board) IsFull() bool {
	isFull := true
	for i := 0; i < b.cols; i++ {
		isFull = isFull && b.board[i].IsFull()
	}
	return isFull
}
