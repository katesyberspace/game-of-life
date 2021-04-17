package main

import (
	"bufio"
	"fmt"
	"mobbing/game"
	"os"
	"strconv"
	"strings"
)

const boarder = `''''''''''''''''''''''`

func getInputs(reader *bufio.Reader) (h, w, numSeeds int, err error) {
	fmt.Print("enter grid height: ")
	h, err = getInputAsInt(reader)
	if err != nil {
		return 0, 0, 0, err
	}
	fmt.Print("enter grid width: ")
	w, err = getInputAsInt(reader)
	if err != nil {
		return 0, 0, 0, err

	}
	fmt.Print("enter number of alive seeds: ")
	numSeeds, err = getInputAsInt(reader)
	if err != nil {
		return 0, 0, 0, err
	}
	return h, w, numSeeds, nil
}

func getInputAsInt(r *bufio.Reader) (int, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	s = strings.Replace(s, "\n", "", -1)
	return strconv.Atoi(s)
}

func printWelcome() {
	fmt.Println(boarder)
	fmt.Print("Game of Life starting\nWarning! This is not the Milton Bradley board game.\n")
	fmt.Println(boarder)
}

func printGameOver() {
	fmt.Println(boarder)
	fmt.Println("Game Over")
	fmt.Println(boarder)
}

func main() {
	var (
		h, w, numSeeds int
	)
	printWelcome()

	reader := bufio.NewReader(os.Stdin)

	h, w, numSeeds, err := getInputs(reader)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Press any enter/return key to spawn the next generation")
	g, _ := game.NewGame(h, w, numSeeds)
	for {
		if g.NoNewLife() || !g.HasEvolved() {
			break
		}
		reader.ReadByte()
		fmt.Print(g.Run())
	}
	printGameOver()
}
