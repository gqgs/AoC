package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

type Point struct {
	grid.Point
	Dist      int
	Direction int
	Prev      *Point
}

func (p Point) String() string {
	return fmt.Sprintf("%s,%d", p.Point.String(), p.Direction)
}

func silver(lines []string) int {
	state := grid.ParseLines(lines)
	start := Point{
		Point:     state.FindPosition('S'),
		Direction: 3,
	}
	end := Point{
		Point: state.FindPosition('E'),
	}

	minheap := generic.NewMinHeap(func(e1 Point, e2 Point) bool {
		return e1.Dist < e2.Dist
	})

	minheap.Push(start)

	visited := make(map[string]int)
	for minheap.Len() > 0 {
		next := minheap.Pop()
		key := next.String()
		score, ok := visited[key]
		if ok && score < next.Dist {
			continue
		}
		visited[key] = next.Dist
		direction := next.Direction
		moves := slices.Collect(next.UpRightDownLeft())
		for i := direction; i < direction+4; i++ {
			move := Point{Point: moves[i%4]}
			if state[move.X][move.Y] == '#' {
				continue
			}

			cost := 1
			switch i - direction {
			case 1, 3:
				cost = 1001
			case 2:
				// robot can't turn 180
				continue
			}

			move.Direction = i % 4
			move.Dist = next.Dist + cost
			move.Prev = &next
			minheap.Push(move)
		}
	}

	var results []int
	for i := range 4 {
		key := fmt.Sprintf("%s,%d", end.Point, i)
		if result, ok := visited[key]; ok {
			results = append(results, result)
		}
	}
	return slices.Min(results)
}

func gold(lines []string, lowestScore int) int {
	state := grid.ParseLines(lines)
	start := Point{
		Point:     state.FindPosition('S'),
		Direction: 3,
	}
	end := Point{
		Point: state.FindPosition('E'),
	}

	minheap := generic.NewMinHeap(func(e1 Point, e2 Point) bool {
		return e1.Dist < e2.Dist
	})

	minheap.Push(start)

	visited := make(map[string]int)
	bestPath := make(map[grid.Point]struct{})

	for minheap.Len() > 0 {
		next := minheap.Pop()
		key := next.String()
		score, ok := visited[key]
		if ok && score < next.Dist {
			continue
		}
		visited[key] = next.Dist

		if end.Equal(next.Point) && next.Dist == lowestScore {
			for iter := &next; iter != nil; iter = iter.Prev {
				bestPath[iter.Point] = struct{}{}
			}
		}

		direction := next.Direction
		moves := slices.Collect(next.UpRightDownLeft())
		for i := direction; i < direction+4; i++ {
			move := Point{Point: moves[i%4]}
			if state[move.X][move.Y] == '#' {
				continue
			}

			cost := 1
			switch i - direction {
			case 1, 3:
				cost = 1001
			case 2:
				// robot can't turn 180
				continue
			}

			move.Direction = i % 4
			move.Dist = next.Dist + cost
			move.Prev = &next
			minheap.Push(move)
		}
	}

	return len(bestPath)
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

	lowest := silver(lines)
	println(lowest)
	println(gold(lines, lowest))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
