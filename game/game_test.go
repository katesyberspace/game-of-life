package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	g := NewGame(1, 1, [][2]int{})
	name := g.Run()
	assert.Equal(t, "Pree", name)
}
