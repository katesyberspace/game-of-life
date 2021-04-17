package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// returns whether cell at y,x coordinate in grid is alive
func isAlive(cell int) bool {
	return cell == 1
}

// returns the number of alive neighbors for a given cell in a grid
func getNumNeighors(h, w, y, x int, grid [][]int) (numNeighbors int) {
	for i := -1; i <= 1; i++ {
		// beyond top & bottom edge
		if h <= y+i || y+i < 0 {
			continue
		}
		for j := -1; j <= 1; j++ {
			// beyond left & right edge
			if w <= x+j || x+j < 0 || (i == 0 && j == 0) {
				continue
			}
			numNeighbors += grid[y+i][x+j]
		}
	}
	return numNeighbors
}

// applies the Game of Life rules, to determine if cell should survive
func survives(alive bool, numNeighbors int) bool {
	return (alive && (numNeighbors == 2 || numNeighbors == 3)) || (!alive && numNeighbors == 3)
}

func (g *Game) printGrid() string {
	sb := strings.Builder{}
	for _, row := range g.grid {
		for _, cell := range row {
			if isAlive(cell) {
				sb.WriteString("*")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

type Game struct {
	grid                [][]int
	h, w                int
	seeds, prevGenSeeds [][2]int
}

func NewGame(h, w int, seeds [][2]int) *Game {
	g := &Game{
		h:     h,
		w:     w,
		seeds: seeds,
	}

	g.createGrid()
	return g
}

//creates a 2D array from h,w and seeds of alive coordinates
func (g *Game) createGrid() {
	grid := make([][]int, g.h)
	for i, row := range grid {
		grid[i] = make([]int, g.w)
		for j := range row {
			grid[i][j] = 0
		}
	}

	for _, seed := range g.seeds {
		grid[seed[0]][seed[1]] = 1
	}

	g.grid = grid
}

func (g *Game) run() string {
	g.prevGenSeeds = g.seeds
	g.seeds = nil
	for y, row := range g.grid {
		for x, cell := range row {
			alive := isAlive(cell)
			numNeighbors := getNumNeighors(g.h, g.w, y, x, g.grid)
			if survives(alive, numNeighbors) {
				g.seeds = append(g.seeds, [2]int{y, x})
			}
		}
	}
	g.createGrid()

	return g.printGrid()

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("enter grid size: ")
	// gridSize, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Print("enter live cells: ")
	// seeds, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(seeds)
	// }
	// fmt.Println("grid size: " + gridSize)
	// fmt.Println("seeds: " + seeds)

	var (
		h, w  int
		seeds [][2]int
	)
	h, w = 20, 50
	seeds = [][2]int{{10, 11}, {10, 12}, {10, 13}, {11, 13}, {14, 13}, {15, 14}, {16, 14}, {15, 13}, {16, 12}}

	g := NewGame(h, w, seeds)
	for {
		if g.noNewLife() || !g.hasEvolved() {
			break
		}
		reader.ReadByte()
		fmt.Print(g.run())
	}
	fmt.Println("Game Over")
}

func (g *Game) noNewLife() bool {
	return len(g.seeds) == 0
}

func (g *Game) hasEvolved() bool {
	if len(g.seeds) != len(g.prevGenSeeds) {
		return true
	}
	sort.Slice(g.seeds, func(i, j int) bool { return g.seeds[i][0] < g.seeds[j][0] })
	sort.Slice(g.prevGenSeeds, func(i, j int) bool { return g.prevGenSeeds[i][0] < g.prevGenSeeds[j][0] })

	for i := range g.seeds {
		if g.seeds[i] != g.prevGenSeeds[i] {
			return true
		}
	}
	return false
}
