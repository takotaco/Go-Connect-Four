package connectFour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_Id(t *testing.T) {
	p := NewPlayer(5, "T")
	assert.Equal(t, 5, p.Id())
}

func Test_Player_Piece(t *testing.T) {
	p := NewPlayer(1, "Q")
	assert.Equal(t, "Q", p.Piece())
}
