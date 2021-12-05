package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Line struct {
	P1, P2 Point
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve() error {
	file, err := os.Open("day5")
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []Line
	var maxX int
	var maxY int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x1, y1, x2, y2 int
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		maxX = max(maxX, max(x1, x2))
		maxY = max(maxY, max(y1, y2))
		lines = append(lines, Line{
			P1: Point{
				X: x1,
				Y: y1,
			},
			P2: Point{
				X: x2,
				Y: y2,
			},
		})
	}

	maxX++
	maxY++

	size := max(maxX, maxY)

	grid := make([][]int, size)
	for y := 0; y < size; y++ {
		grid[y] = append(grid[y], make([]int, size)...)
	}

	for _, line := range lines {
		var minx, maxx, miny, maxy int
		if line.P1.X < line.P2.X {
			minx = line.P1.X
			maxx = line.P2.X
			miny = line.P1.Y
			maxy = line.P2.Y
		} else {
			minx = line.P2.X
			maxx = line.P1.X
			miny = line.P2.Y
			maxy = line.P1.Y
		}

		f := func(x int) int {
			a := (miny - maxy) / (minx - maxx)
			b := (miny*maxx - maxy*minx) / (maxx - minx)
			return a*x + b
		}

		if minx == maxx || miny == maxy || abs(maxy-miny) == abs(maxx-minx) {
			if minx == maxx {
				for y := min(miny, maxy); y <= max(miny, maxy); y++ {
					grid[minx][y]++
				}
			} else {
				for x := minx; x <= maxx; x++ {
					grid[x][f(x)]++
				}
			}
		}
	}

	var silver int
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if grid[x][y] >= 2 {
				silver++
			}
			//print(grid[x][y], " ")
		}
		// println()
	}
	println("silver:", silver)

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
