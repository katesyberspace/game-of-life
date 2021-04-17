package main

import (
	"bufio"
	"fmt"
	"mobbing/game"
	"os"
)

const boarder = `''''''''''''''''''''''`

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		h, w  int
		seeds [][2]int
	)
	//TODO: h,w,seeds as input
	h, w = 40, 80
	seeds = [][2]int{
		{10, 11}, {10, 12}, {10, 13}, {11, 13},
		{14, 13}, {15, 14}, {16, 14}, {15, 13},
		{16, 12}, {2, 30}, {2, 31}, {1, 30},
		{0, 30},
		{3, 27}, {3, 28}, {3, 29}, {3, 28}, {4, 28}, {5, 28},
		{3, 40}, {3, 41}, {4, 42}, {5, 42}, {6, 43}, {7, 43},
		{8, 43}, {8, 70}, {9, 71}, {10, 72}, {10, 73}, {11, 74},
		{35, 43}, {35, 70}, {35, 71}, {35, 72}, {35, 73}, {35, 74},
		{20, 43}, {20, 42}, {20, 41}, {20, 44}, {19, 43}, {18, 43},
		{17, 43}, {16, 70}, {16, 71}, {17, 72}, {17, 73}, {15, 74},
		{17, 60}, {16, 60}, {16, 59}, {17, 59}, {17, 61}, {15, 60},
	}

	g, _ := game.NewGame(h, w, seeds)
	fmt.Println(boarder)
	fmt.Print("Game of Life starting\nWarning! This is not the Milton Bradley board game.\nPress any enter/return key to spawn the next generation\n")
	fmt.Println(boarder)

	for {
		if g.NoNewLife() || !g.HasEvolved() {
			break
		}
		reader.ReadByte()
		fmt.Print(g.Run())
	}
	fmt.Println("Game Over")
}
