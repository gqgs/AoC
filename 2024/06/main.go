package main

import (
	"bufio"
	"fmt"
	"maps"
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
		switch v := peak(s, lines, direction); v {
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

func nextPoint(p grid.Point, direction int) grid.Point {
	var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	return grid.Point{
		X: p.X + directions[direction][0],
		Y: p.Y + directions[direction][1],
	}
}

func gold(lines []string) int {
	var direction int
	visited := make(map[string]struct{})
	vvisited := make(map[string]struct{})
	initPosition := findGuardPosition(lines)
	return len(goldAux(initPosition, direction, lines, visited, vvisited, nil))
}

func goldAux(start grid.Point, direction int, lines []string, visited, vvisited map[string]struct{}, obstruction *grid.Point) []*grid.Point {
	key := func(p grid.Point, direction int) string {
		return fmt.Sprintf("%s-%d", p.String(), direction)
	}

	var obstructions []*grid.Point
	stack := new(generic.Stack[grid.Point])
	stack.Push(start)

	for !stack.Empty() {
		next := stack.Pop()
		nextKey := key(next, direction)
		if _, isCycle := visited[nextKey]; isCycle {
			if obstruction == nil {
				return obstructions
			}
			obstructions = append(obstructions, obstruction)
			return obstructions
		}
		visited[nextKey] = struct{}{}
		vvisited[next.String()] = struct{}{}

		switch v := peak(next, lines, direction); v {
		case '^':
			stack.Push(nextPoint(next, direction))
		case '.':
			np := nextPoint(next, direction)
			stack.Push(np)
			_, alreadyVisited := vvisited[np.String()]
			if obstruction == nil && !alreadyVisited {
				result := goldAux(next, (direction+1)%4, grid.Replace(lines, "O", np.X, np.Y), maps.Clone(visited), maps.Clone(vvisited), &np)
				obstructions = append(obstructions, result...)
			}
		case '#', 'O':
			stack.Push(next)
			direction = (direction + 1) % 4
		case 'X':
			return obstructions
		}
	}
	return obstructions
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
	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
