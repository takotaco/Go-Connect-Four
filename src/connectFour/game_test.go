package connectFour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Game_switchTurn(t *testing.T) {
	g := NewGame()
	assert.Equal(t, 1, g.activeP.id)
	g.switchTurn()
	assert.Equal(t, 2, g.activeP.id)
	g.switchTurn()
	assert.Equal(t, 1, g.activeP.id)
}
