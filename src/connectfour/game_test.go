package connectfour

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

func Test_Game_TakeTrun(t *testing.T) {
	g := NewGame()
	ok, err := g.TakeTurn(-1)
	assert.Equal(t, false, ok)
	assert.Error(t, err)
	ok, err = g.TakeTurn(1)
	assert.Equal(t, true, ok)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	ok, err = g.TakeTurn(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, g.winningP, g.activeP)
}

func Test_Game_GetActivePlayer(t *testing.T) {
	g := NewGame()
	p := g.GetActivePlayer()
	assert.Equal(t, g.p1, p)
}

func Test_Game_IsGameOver(t *testing.T) {
	g := NewGame()
	o := g.IsGameOver()
	assert.Equal(t, false, o)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	o = g.IsGameOver()
	assert.Equal(t, true, o)
}

func Test_Game_GetWinningPlayer(t *testing.T) {
	g := NewGame()
	p := g.GetWinningPlayer()
	assert.Equal(t, Player{}, p)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	g.TakeTurn(1)
	g.TakeTurn(2)
	p = g.GetWinningPlayer()
	assert.Equal(t, g.p1, p)
}
