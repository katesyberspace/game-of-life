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
	h, w = 20, 50
	seeds = [][2]int{
		{10, 11}, {10, 12}, {10, 13}, {11, 13},
		{14, 13}, {15, 14}, {16, 14}, {15, 13},
		{16, 12}, {2, 30}, {2, 31}, {1, 30},
		{0, 30}, {3, 29}, {3, 28},
	}

	g, _ := game.NewGame(h, w, seeds)
	fmt.Println(boarder)
	fmt.Print("Game of Life starting\nWarning! This is not the Milton Bradley board game.\nPress any key to spawn the next generation\n")
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
