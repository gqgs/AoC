package main

import (
	"bufio"
	"os"
)

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	println(silver(input))

	return nil
}

func silver(input []string) int {
	grid := make([][]rune, len(input))
	for i := range grid {
		grid[i] = make([]rune, len(input[i]))
		for j := range input[i] {
			grid[i][j] = rune(input[i][j])
		}
	}

	executeSwaps := func(swaps [][4]int) {
		for _, swap := range swaps {
			x0, y0, x1, y1 := swap[0], swap[1], swap[2], swap[3]
			grid[y0][x0], grid[y1][x1] = grid[y1][x1], grid[y0][x0]
		}
	}

	for step := 0; ; step++ {
		var swaped int
		var swaps [][4]int

		// move right
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				nextx := (x + 1) % len(grid[y])
				if grid[y][x] == '>' && grid[y][nextx] == '.' {
					swaps = append(swaps, [4]int{x, y, nextx, y})
				}
			}
		}

		swaped += len(swaps)
		executeSwaps(swaps)
		swaps = swaps[:0]

		// move down
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				nexty := (y + 1) % len(grid)
				if grid[y][x] == 'v' && grid[nexty][x] == '.' {
					swaps = append(swaps, [4]int{x, y, x, nexty})
				}
			}
		}

		swaped += len(swaps)
		if swaped == 0 {
			return step + 1
		}

		executeSwaps(swaps)
	}
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
