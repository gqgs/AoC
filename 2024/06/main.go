package main

import (
	"bufio"
	"fmt"
	"os"

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

func goldAux(lines []string) int {
	visited := make(map[string]struct{})
	initPosition := findGuardPosition(lines)

	direction := 0
	s := initPosition

	key := func() string {
		return fmt.Sprintf("%s-%d", s.String(), direction)
	}

	for {
		if _, ok := visited[key()]; ok {
			return 1
		}
		v := peak(s, lines, direction)
		switch v {
		case '^':
			visited[key()] = struct{}{}
			next(&s, direction)
		case '.':
			visited[key()] = struct{}{}
			next(&s, direction)
		case '#', 'O':
			visited[key()] = struct{}{}
			direction = (direction + 1) % 4
		case 'X':
			visited[key()] = struct{}{}
			return 0
		}
	}
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
