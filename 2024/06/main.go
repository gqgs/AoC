package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func findGuardPosition(lines []string) grid.Point {
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '^' {
				return grid.Point{
					X: i,
					Y: j,
				}
			}
		}
	}
	panic("not found")
}

func peak(p grid.Point, lines []string, direction int) byte {
	var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	x := p.X + directions[direction][0]
	y := p.Y + directions[direction][1]
	return lines[x][y]
}

func next(p *grid.Point, direction int) {
	var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	p.X += directions[direction][0]
	p.Y += directions[direction][1]
}

func silver(lines []string) int {
	visited := make(map[string]struct{})
	initPosition := findGuardPosition(lines)

	direction := 0
	s := initPosition
	for {
		v := peak(s, lines, direction)
		switch v {
		case '^':
			visited[s.String()] = struct{}{}
			next(&s, direction)
		case '.':
			visited[s.String()] = struct{}{}
			next(&s, direction)
		case '#':
			visited[s.String()] = struct{}{}
			direction = (direction + 1) % 4
		case 'X':
			visited[s.String()] = struct{}{}
			return len(visited)
		}
	}
}

func replace(lines []string, c string, x, y int) []string {
	var newLines []string
	for i := range lines {
		if i == x {
			line := lines[i][:y] + c + lines[i][y+1:]
			newLines = append(newLines, line)
			continue
		}
		newLines = append(newLines, lines[i])
	}
	return newLines
}

func gold(lines []string) int {
	var total int
	for i := range lines {
		for j := range lines[i] {
			switch lines[i][j] {
			case '#', '^':
				continue
			}
			total += goldAux(replace(lines, "O", i, j))
		}
	}

	return total
}

func nextPoint(p grid.Point, direction int) grid.Point {
	var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	return grid.Point{
		X: p.X + directions[direction][0],
		Y: p.Y + directions[direction][1],
	}
}

func goldAux(lines []string) int {
	visited := make(map[string]struct{})
	initPosition := findGuardPosition(lines)

	direction := 0
	key := func(p grid.Point) string {
		return fmt.Sprintf("%s-%d", p.String(), direction)
	}

	stack := new(generic.Stack[grid.Point])
	stack.Push(initPosition)

	for !stack.Empty() {
		next := stack.Pop()
		nextKey := key(next)
		if _, isCycle := visited[nextKey]; isCycle {
			return 1
		}
		visited[nextKey] = struct{}{}

		v := peak(next, lines, direction)
		switch v {
		case '^':
			stack.Push(nextPoint(next, direction))
		case '.':
			stack.Push(nextPoint(next, direction))
		case '#', 'O':
			stack.Push(next)
			direction = (direction + 1) % 4
		case 'X':
			return 0
		}
	}
	return 0
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

	lines = grid.Fill(lines, "X", 1)
	// println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
