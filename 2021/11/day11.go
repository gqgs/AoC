package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/gqgs/AoC2021/generic"
)

const gridSize = 10

type Stack = generic.Stack[[2]int]

func solve() {
	var totalFlashes int
	grid := readGrid()
	stack := make(Stack, 0)
	for i := 0; ; i++ {
		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				grid[x][y]++
				if grid[x][y] > 9 {
					stack.Push([2]int{x, y})
				}
			}
		}

		var flashed [gridSize][gridSize]int
		for len(stack) > 0 {
			next := stack.Pop()
			x, y := next[0], next[1]
			if flashed[x][y] > 0 {
				continue
			}
			flashed[x][y]++
			for dx := max(0, x-1); dx < min(gridSize, x+2); dx++ {
				for dy := max(0, y-1); dy < min(gridSize, y+2); dy++ {
					grid[dx][dy]++
					if grid[dx][dy] > 9 && flashed[dx][dy] == 0 {
						stack.Push([2]int{dx, dy})
					}
				}
			}
		}

		var currentFlahes int
		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				if grid[x][y] > 9 {
					grid[x][y] = 0
				}
				if flashed[x][y] > 0 {
					if i < 100 {
						totalFlashes++
					}
					currentFlahes++
				}
			}
		}

		if currentFlahes == gridSize*gridSize {
			println("silver:", totalFlashes)
			println("gold:", i+1)
			return
		}
	}
}

func readGrid() [gridSize][gridSize]int {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [gridSize][gridSize]int
	var row int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var column int
		for _, c := range scanner.Text() {
			n, _ := strconv.Atoi(string(c))
			grid[row][column] = n
			column++
		}
		row++
	}
	return grid
}

func main() {
	solve()
}
