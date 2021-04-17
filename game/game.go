package game

import (
	"sort"
	"strings"
)

const (
	ALIVE, DEAD = 1, 0
)

// Game contains the h,w and seeds specificied by input
// and the grid itself
type Game struct {
	grid                [][]int
	h, w                int
	seeds, prevGenSeeds [][2]int
}

// ValidInputs checks the h, w and seeds provided
// must be greater than 0, and seeds within h,w boarder
func ValidInputs(h, w int, seeds [][2]int) bool {
	if h < 0 || w < 0 {
		return false
	}
	for _, s := range seeds {
		if s[0] < 0 || s[0] >= h ||
			s[1] < 0 || s[1] >= w {
			return false
		}
	}
	return true
}

// NewGame returns a new instance of Game with
// the first gen grid created
func NewGame(h, w int, seeds [][2]int) *Game {
	g := &Game{
		h:     h,
		w:     w,
		seeds: seeds,
	}

	g.createGrid()
	return g
}

// Run is the function that runs all game logic
// and creates the grid for the next generation
func (g *Game) Run() string {
	g.prevGenSeeds = g.seeds
	g.seeds = nil
	for y, row := range g.grid {
		for x, cell := range row {
			alive := isAlive(cell)
			numNeighbors := g.getNumNeighors(y, x)
			if survives(alive, numNeighbors) {
				g.seeds = append(g.seeds, [2]int{y, x})
			}
		}
	}
	return g.createGrid()
}

//creates the grid as a 2D array from h,w and seeds coordinates
//and returns the string formatted grid
func (g *Game) createGrid() string {
	grid := make([][]int, g.h)
	for i, row := range grid {
		grid[i] = make([]int, g.w)
		for j := range row {
			grid[i][j] = DEAD
		}
	}

	for _, seed := range g.seeds {
		grid[seed[0]][seed[1]] = ALIVE
	}

	g.grid = grid
	return g.printGrid()
}

// returns whether cell at y,x coordinate in grid is alive
func isAlive(cell int) bool {
	return cell == ALIVE
}

// returns the number of alive neighbors for a given cell in a grid
func (g *Game) getNumNeighors(y, x int) (numNeighbors int) {
	for i := -1; i <= 1; i++ {
		// beyond top & bottom edge
		if g.h <= y+i || y+i < 0 {
			continue
		}
		for j := -1; j <= 1; j++ {
			// beyond left & right edge
			if g.w <= x+j || x+j < 0 || (i == 0 && j == 0) {
				continue
			}
			numNeighbors += g.grid[y+i][x+j]
		}
	}
	return numNeighbors
}

// applies the Game of Life rules, to determine if cell should survive
func survives(alive bool, numNeighbors int) bool {
	return (alive && (numNeighbors == 2 || numNeighbors == 3)) || (!alive && numNeighbors == 3)
}

// Returns a string from the elements of the grid
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

// NoNewLife returns true if no cells survive from the previous generation
func (g *Game) NoNewLife() bool {
	return len(g.seeds) == 0
}

// HasEvolved returns true if seeds and previous gen seeds are different
func (g *Game) HasEvolved() bool {
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
