package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func silver(lines []string, gridSize, numberBytes int) int {
	return shared(lines, gridSize, numberBytes)
}

func gold(lines []string, gridSize, numberBytes int) string {
	low := numberBytes
	high := len(lines)
	for low < high {
		// binary search
		mid := ((high - low) / 2) + 1
		s := shared(lines, gridSize, low+mid)
		if s == 0 {
			high -= mid
		} else {
			low += mid
		}
	}
	return lines[low]
}

func shared(lines []string, gridSize, numberBytes int) int {
	var points []grid.Point
	for _, line := range lines {
		var point grid.Point
		fmt.Sscanf(line, "%d,%d", &point.X, &point.Y)
		points = append(points, point)
	}

	state := grid.NewSquared(gridSize)
	state.FillPerimeter()

	for _, p := range points[:numberBytes] {
		state[p.X+1][p.Y+1] = '#'
	}

	visited := make(map[string]int)
	minheap := generic.NewMinHeap(func(e1, e2 grid.Point) bool {
		return e1.Score < e2.Score
	})

	minheap.Push(grid.Point{
		X: 1,
		Y: 1,
	})

	for minheap.Len() > 0 {
		next := minheap.Pop()
		score, ok := visited[next.String()]

		if ok && score <= next.Score {
			continue
		}
		visited[next.String()] = next.Score
		for move := range next.UpRightDownLeft() {
			if state[move.X][move.Y] == '#' {
				continue
			}

			move.Score = next.Score + 1
			minheap.Push(move)
		}
	}

	return visited[grid.Point{X: gridSize, Y: gridSize}.String()]
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			continue
		}
		lines = append(lines, next)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	println(silver(lines, 71, 1024))
	println(gold(lines, 71, 1024))

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
