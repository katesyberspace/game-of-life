package game

import "fmt"

//creates a 2D array from h,w and seeds of alive coordinates
func createGrid(h, w int, seeds [][2]int) (grid [][]int) {
	grid = make([][]int, h)
	for i, row := range grid {
		grid[i] = make([]int, w)
		for j := range row {
			grid[i][j] = 0
		}
	}

	for _, seed := range seeds {
		grid[seed[0]][seed[1]] = 1
	}

	return grid
}

// returns whether cell at y,x coordinate in grid is alive
func isAlive(y, x int, grid [][]int) bool {
	return grid[y][x] == 1
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

/*
1st:
0 0 0 0 0 0 0 0
0 0 0 0 1 0 0 0
0 0 0 1 1 0 0 0
0 0 0 0 0 0 0 0

2nd:
0 0 0 0 0 0 0 0
0 0 0 1 1 0 0 0
0 0 0 1 1 0 0 0
0 0 0 0 0 0 0 0

{{1, 4}, {2, 3}, {2, 4}}

*/

func game() string {
	var (
		h, w  int
		seeds [][2]int
	)
	h, w = 4, 8
	seeds = [][2]int{{1, 4}, {2, 3}, {2, 4}}

	nextGenSeeds := [][2]int{}
	grid := createGrid(h, w, seeds)

	for y, row := range grid {
		for x := range row {
			alive := isAlive(y, x, grid)
			numNeighbors := getNumNeighors(h, w, y, x, grid)
			if (alive && (2 <= numNeighbors || numNeighbors <= 3)) || (!alive && numNeighbors == 3) {
				nextGenSeeds = append(nextGenSeeds, [2]int{y, x})
			}
		}
	}

	grid = createGrid(h, w, nextGenSeeds)

	return fmt.Sprintf("%v", grid)

}
