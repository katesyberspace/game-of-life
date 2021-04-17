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
		h, w, numSeeds int
	)
	//TODO: h,w,seeds as input
	h, w, numSeeds = 30, 60, 1000

	g, _ := game.NewGame(h, w, numSeeds)
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
