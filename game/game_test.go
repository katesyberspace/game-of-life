package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T){
	name := game()
	assert.Equal(t, "Pree",name)
}