# Conway's Game of Life - Challenge

The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970.

![Gospers Glider gun](/resources/gospers-glider-gun.gif)

The game is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves. It is Turing complete and can simulate a universal constructor or any other Turing machine.

Source : [Wikipedia](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

## 4 rules

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, alive or dead, (or populated and unpopulated, respectively). Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

* Any live cell with fewer than two live neighbours dies, as if by underpopulation.
* Any live cell with two or three live neighbours lives on to the next generation.
* Any live cell with more than three live neighbours dies, as if by overpopulation.
* Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

These rules, which compare the behavior of the automaton to real life, can be condensed into the following:

1. Any live cell with two or three neighbors survives.
2. Any dead cell with three live neighbors becomes a live cell.
3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.

The initial pattern constitutes the seed of the system. The first generation is created by applying the above rules simultaneously to every cell in the seed; births and deaths occur simultaneously, and the discrete moment at which this happens is sometimes called a tick. Each generation is a pure function of the preceding one. The rules continue to be applied repeatedly to create further generations.

## Scope of this Coding Dojo :

* Cells that goes beyond the grid dies.

## Example input:
```
Grid height, grid width, number of living cells to seed
```
## Expected output:
Compute and display the next generation on any key press, and exit the program when the subsequent generation doesn't change/evolve or all cells are dead.
```
Generation 1:
........
....*...
...**...
........

Generation 2:
........
...**...
...**...
........
```
And so on and so forth.

## Instructions for running
`go run main.go`

### Docker
- Install Docker from https://www.docker.com/products/docker-desktop
- Run `docker build .`

### Using your local setup:
Install go, once done run `go test -v ./..` in your game folder