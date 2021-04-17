package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGameSuccess(t *testing.T) {
	h, w := 5, 5
	seeds := [][2]int{
		{1, 2},
		{2, 3},
	}
	g, err := NewGame(h, w, seeds)
	assert.Equal(t, h, g.h)
	assert.Equal(t, w, g.w)
	assert.Equal(t, seeds, g.seeds)
	assert.Nil(t, err)
}

func TestNewGameFail(t *testing.T) {
	h, w := 1, 1
	seeds := [][2]int{
		{1, 2},
		{2, 3},
	}
	_, err := NewGame(h, w, seeds)
	assert.NotNil(t, err)
}

var tests = []struct {
	h, w  int
	seeds [][2]int
	exp   string
}{
	{
		h: 4, w: 8, seeds: [][2]int{{2, 5}, {3, 4}, {3, 5}},
		exp: "........\n........\n....**..\n....**..\n",
	},
	{
		h: 10, w: 10, seeds: [][2]int{{7, 5}, {6, 5}, {5, 5}},
		exp: "..........\n..........\n..........\n..........\n..........\n..........\n....***...\n..........\n..........\n..........\n",
	},
	{
		h: 12, w: 11, seeds: [][2]int{{0, 9}, {1, 8}, {2, 7}},
		exp: "...........\n........*..\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n",
	},
}

func TestFind(t *testing.T) {
	for _, e := range tests {
		g, _ := NewGame(e.h, e.w, e.seeds)
		res := g.Run()
		assert.Equal(t, e.exp, res)
	}
}
