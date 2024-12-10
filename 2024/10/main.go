package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func findPath(lines []string, point grid.Point, path []grid.Point, valid map[string]struct{}) {
	if lines[point.X][point.Y] == '9' {
		var keys []string
		for _, p := range path {
			keys = append(keys, p.String())
		}

		key := strings.Join(keys, "|")
		valid[key] = struct{}{}
		return
	}

Next:
	for _, possible := range [][2]int{
		{point.X - 1, point.Y},
		{point.X, point.Y + 1},
		{point.X + 1, point.Y},
		{point.X, point.Y - 1},
	} {
		possibleX := possible[0]
		possibleY := possible[1]

		if lines[possibleX][possibleY]-lines[point.X][point.Y] == 1 {
			next := grid.Point{
				X: possibleX,
				Y: possibleY,
			}

			for _, visited := range path {
				if next.X == visited.X && next.Y == visited.Y {
					// already in path
					continue Next
				}
			}

			findPath(lines, next, append(path, next), valid)
		}
	}
}

func silver(lines []string) int {
	stack := new(generic.Stack[grid.Point])
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != '0' {
				continue
			}
			stack.Push(grid.Point{
				X: i,
				Y: j,
			})
		}
	}

	validPaths := make(map[string]struct{})
	for !stack.Empty() {
		next := stack.Pop()
		findPath(lines, next, []grid.Point{next}, validPaths)
	}

	score := make(map[string]struct{})
	for path := range validPaths {
		parts := strings.Split(path, "|")
		start := parts[0]
		end := parts[len(parts)-1]
		score[start+end] = struct{}{}
	}

	return len(score)
}

func gold(lines []string) int {
	stack := new(generic.Stack[grid.Point])
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != '0' {
				continue
			}
			stack.Push(grid.Point{
				X: i,
				Y: j,
			})
		}
	}

	validPaths := make(map[string]struct{})
	for !stack.Empty() {
		next := stack.Pop()
		findPath(lines, next, []grid.Point{next}, validPaths)
	}

	score := make(map[string]struct{})
	for path := range validPaths {
		score[path] = struct{}{}
	}

	return len(score)
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

	lines = grid.Fill(lines, "X", 1)

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
