package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var maxX, maxy int
	var points [][2]int
	var folds []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "fold") {
			folds = append(folds, txt)
			continue
		}

		var x, y int
		fmt.Sscanf(txt, "%d,%d", &x, &y)
		maxX = max(x, maxX)
		maxy = max(y, maxy)
		points = append(points, [2]int{x, y})
	}

	size := max(maxX, maxy) + 1
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	for _, p := range points {
		x, y := p[0], p[1]
		grid[x][y] = 1
	}

	grid = foldGrid(grid, folds...)

	for i := 0; i < 10; i++ {
		for j := 0; j < 50; j++ {
			if grid[j][i] == 1 {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}

	return nil
}

func foldGrid(grid [][]int, folds ...string) [][]int {
	if len(folds) == 0 {
		return grid
	}
	fold := folds[0]

	var direction rune
	var position int

	fmt.Sscanf(fold, "fold along %c=%d", &direction, &position)
	switch direction {
	case 'y':
		dest := position - 1
		for y := position + 1; y < len(grid); y++ {
			for x := 0; x < len(grid); x++ {
				if dest >= 0 {
					grid[x][dest] |= grid[x][y]
				}
				grid[x][y] = 0
			}
			dest--
		}
	case 'x':
		dest := position - 1
		for x := position + 1; x < len(grid); x++ {
			for y := 0; y < len(grid); y++ {
				if dest >= 0 {
					grid[dest][y] |= grid[x][y]
				}
				grid[x][y] = 0
			}
			dest--
		}
	}
	return foldGrid(grid, folds[1:]...)
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
