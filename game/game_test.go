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
	g := NewGame(h, w, seeds)
	assert.Equal(t, h, g.h)
	assert.Equal(t, w, g.w)
	assert.Equal(t, seeds, g.seeds)
}

// var tests = []struct {
// 	g   *Game
// 	exp string
// }{
// 	{
// 		NewGame(4, 8, [][2]int{{2, 5}, {3, 4}, {3, 5}}),
// 		"........\n........\n....**..\n....**..\n",
// 	},
// 	{
// 		NewGame(10, 10, [][2]int{{7, 5}, {6, 5}, {5, 5}}),
// 		"..........\n..........\n..........\n..........\n..........\n..........\n....***...\n..........\n..........\n..........\n",
// 	},
// 	{
// 		NewGame(12, 11, [][2]int{{0, 9}, {1, 8}, {2, 7}}),
// 		"...........\n........*..\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n...........\n",
// 	},
// }

// func TestFind(t *testing.T) {
// 	for _, e := range tests {
// 		res := e.g.Run()
// 		assert.Equal(t, e.exp, res)
// 	}
// }
