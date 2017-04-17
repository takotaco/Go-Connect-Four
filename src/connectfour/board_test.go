package connectfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Board_MakeMove(t *testing.T) {
	b := NewBoard(1, 1)
	ok, err, win := b.MakeMove(0, "Liz")
	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, win)

	ok, err, win = b.MakeMove(0, "Liz")
	assert.Equal(t, false, ok)
}

func Test_Board_CheckMoveHelper(t *testing.T) {
	b := NewBoard(4, 4)
	s := b.checkMoveHelper(2, 2, "LK")
	assert.Equal(t, 0, s)
	b.MakeMove(1, "LK")
	b.MakeMove(1, "LK")
	b.MakeMove(1, "LK")
	b.MakeMove(1, "LK")
	s = b.checkMoveHelper(1, 0, "LK")
	assert.Equal(t, 3, s)
	b.MakeMove(0, "LK")
	b.MakeMove(2, "Z")
	b.MakeMove(2, "LK")
	b.MakeMove(2, "LK")
	b.MakeMove(3, "LK")
	b.MakeMove(3, "Z")
	b.MakeMove(3, "Z")
	b.MakeMove(3, "LK")
	s = b.checkMoveHelper(0, 0, "LK")
	assert.Equal(t, 3, s)
	b.MakeMove(0, "Z")
	b.MakeMove(0, "Z")
	b.MakeMove(0, "LK")
	s = b.checkMoveHelper(3, 0, "LK")
	assert.Equal(t, 3, s)
	b.MakeMove(2, "LK")
	s = b.checkMoveHelper(0, 3, "LK")
	assert.Equal(t, 3, s)
}

func Test_Board_checkUp(t *testing.T) {
	b := NewBoard(2, 1)
	s := b.checkUp(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "L")
	s = b.checkUp(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "L")
	s = b.checkUp(0, 0, "L")
	assert.Equal(t, 1, s)
}

func Test_Board_checkDown(t *testing.T) {
	b := NewBoard(2, 1)
	s := b.checkDown(0, 1, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "L")
	s = b.checkDown(0, 1, "L")
	assert.Equal(t, 1, s)
	b.MakeMove(0, "L")
	s = b.checkDown(0, 1, "L")
	assert.Equal(t, 1, s)

}

func Test_Board_checkLeft(t *testing.T) {
	b := NewBoard(1, 4)
	s := b.checkLeft(3, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(3, "L")
	s = b.checkLeft(3, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(2, "L")
	s = b.checkLeft(3, 0, "L")
	assert.Equal(t, 1, s)
	b.MakeMove(1, "L")
	s = b.checkLeft(3, 0, "L")
	assert.Equal(t, 2, s)
	b.MakeMove(0, "L")
	s = b.checkLeft(3, 0, "L")
	assert.Equal(t, 3, s)
}

func Test_Board_checkRight(t *testing.T) {
	b := NewBoard(1, 4)
	s := b.checkRight(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "L")
	s = b.checkRight(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(1, "L")
	s = b.checkRight(0, 0, "L")
	assert.Equal(t, 1, s)
	b.MakeMove(2, "L")
	s = b.checkRight(0, 0, "L")
	assert.Equal(t, 2, s)
	b.MakeMove(3, "L")
	s = b.checkRight(0, 0, "L")
	assert.Equal(t, 3, s)
}

func Test_Board_checkUpLeft(t *testing.T) {
	b := NewBoard(2, 2)
	s := b.checkUpLeft(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "NotL")
	b.MakeMove(0, "L")
	s = b.checkUpLeft(0, 1, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(1, "L")
	s = b.checkUpLeft(1, 0, "L")
	assert.Equal(t, 1, s)
}

func Test_Board_checkDownLeft(t *testing.T) {
	b := NewBoard(2, 2)
	s := b.checkDownLeft(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "L")
	s = b.checkDownLeft(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(1, "NotL")
	b.MakeMove(1, "L")
	s = b.checkDownLeft(1, 1, "L")
	assert.Equal(t, 1, s)
}

func Test_Board_checkUpRight(t *testing.T) {
	b := NewBoard(2, 2)
	s := b.checkUpRight(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(1, "NotL")
	b.MakeMove(1, "L")
	s = b.checkUpRight(1, 1, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "L")
	s = b.checkUpRight(0, 0, "L")
	assert.Equal(t, 1, s)
}

func Test_Board_checkDownRight(t *testing.T) {
	b := NewBoard(2, 2)
	s := b.checkDownRight(0, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(1, "L")
	s = b.checkDownRight(1, 0, "L")
	assert.Equal(t, 0, s)
	b.MakeMove(0, "NotL")
	b.MakeMove(0, "L")
	s = b.checkDownRight(0, 1, "L")
	assert.Equal(t, 1, s)
}

func Test_Board_IsFull(t *testing.T) {
	b := NewBoard(1, 1)
	f := b.IsFull()
	assert.Equal(t, false, f)
	b.MakeMove(0, "XD")
	f = b.IsFull()
	assert.Equal(t, true, f)
}
