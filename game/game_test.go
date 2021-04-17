package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGameSuccess(t *testing.T) {
	h, w, ns := 5, 5, 15
	g, err := NewGame(h, w, ns)
	assert.Equal(t, h, g.h)
	assert.Equal(t, w, g.w)
	assert.Equal(t, ns, len(g.seeds))
	assert.Nil(t, err)
}

func TestNewGameFail(t *testing.T) {
	h, w, ns := 1, 1, 3

	_, err := NewGame(h, w, ns)
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

func TestRun(t *testing.T) {
	for _, e := range tests {
		g := &Game{
			h:     e.h,
			w:     e.w,
			seeds: e.seeds,
			grid:  createGrid(e.h, e.w, e.seeds),
		}
		res := g.Run()
		assert.Equal(t, e.exp, res)
	}
}
