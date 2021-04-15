package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter grid size: ")
	gridSize, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("enter live cells: ")
	liveCells, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(liveCells)
	}
	fmt.Println("grid size: " + gridSize)
	fmt.Println("liveCells: " + liveCells)
}
