package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gqgs/AoC2021/grid"
)

func silver(lines []string) int {
	return shared(lines, 2)
}

func gold(lines []string) int {
	return shared(lines, 20)
}

func shared(lines []string, cheatLen int) int {
	state := grid.ParseLines(lines)
	start := state.FindPosition('S')
	end := state.FindPosition('E')

	path, shortest := state.ShortestPath(start, end)
	distances := make(map[grid.Point]int)
	for distance := range path {
		distances[path[distance]] = shortest - distance
	}

	var total int
	for p1, d1 := range distances {
		for p2, d2 := range distances {
			d := p1.Distance(p2)
			if d <= cheatLen && (d1-d2-d) >= 100 {
				total++
			}

		}
	}

	return total
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

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
