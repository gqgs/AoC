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

// func silver(lines []string) int {
// 	visited := make(map[string]struct{})
// 	initPosition := findGuardPosition(lines)

// 	direction := 0
// 	s := initPosition
// 	for {
// 		switch v := peak(s, lines, direction); v {
// 		case '^':
// 			visited[key(s)] = struct{}{}
// 			next(&s, direction)
// 		case '.':
// 			visited[s.String()] = struct{}{}
// 			next(&s, direction)
// 		case '#':
// 			visited[s.String()] = struct{}{}
// 			direction = (direction + 1) % 4
// 		case 'X':
// 			visited[s.String()] = struct{}{}
// 			return len(visited)
// 		}
// 		fmt.Println("visited", len(visited))
// 		time.Sleep(time.Second)
// 	}
// }

var cycles int

func silver(lines []string) int {
	visited := make(map[string]struct{})
	initPosition := findGuardPosition(lines)

	direction := 0

	// fmt.Println("silver", len(lines))
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	// fmt.Println()

	key := func(p grid.Point) string {
		return fmt.Sprintf("%s-%d", p.String(), direction)
	}
	s := initPosition
	visitedPoints := make(map[string]struct{})
	for {
		visitedPoints[s.String()] = struct{}{}

		if _, isCycle := visited[key(s)]; isCycle {
			for i := range lines {
				for j := range lines[i] {
					if i == initPosition.X && j == initPosition.Y {
						fmt.Printf("\033[0;32m%c\033[0m", '^')
						continue
					}

					key := fmt.Sprintf("(%d,%d)", i, j)
					if _, ok := visitedPoints[key]; ok {
						fmt.Printf("\033[0;31m%c\033[0m", '*')
					} else {
						if lines[i][j] == 'O' {
							fmt.Printf("\033[0;36m%c\033[0m", 'O')
						} else {
							fmt.Printf("%c", lines[i][j])
						}
					}
				}
				fmt.Println()
			}

			cycles++
			if cycles > 50 {
				os.Exit(1)
			}
			return 0
		}

		v := peak(s, lines, direction)
		switch v {
		case '^':
			visited[key(s)] = struct{}{}
			next(&s, direction)
		case '.':
			visited[key(s)] = struct{}{}
			next(&s, direction)
		case '#', 'O':
			visited[key(s)] = struct{}{}
			direction = (direction + 1) % 4
		case 'X':
			visited[key(s)] = struct{}{}
			return len(visited)
		}
		// fmt.Printf("visited %d %s %c\n", len(visited), key(s), v)
		// time.Sleep(time.Second)
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
	initPosition := findGuardPosition(lines)
	return goldAux(initPosition, direction, lines, visited, false)
}

func goldAux(start grid.Point, direction int, lines []string, visited map[string]struct{}, blocked bool) int {
	var total int
	key := func(p grid.Point) string {
		return fmt.Sprintf("%s-%d", p.String(), direction)
	}

	stack := new(generic.Stack[grid.Point])
	stack.Push(start)

	for !stack.Empty() {
		next := stack.Pop()
		nextKey := key(next)
		if _, isCycle := visited[nextKey]; isCycle {
			silver(lines)
			return 1
		}
		visited[nextKey] = struct{}{}

		switch v := peak(next, lines, direction); v {
		case '^':
			stack.Push(nextPoint(next, direction))
		case '.':
			np := nextPoint(next, direction)
			stack.Push(np)
			if !blocked {
				total += goldAux(next, (direction+1)%4, replace(lines, "O", np.X, np.Y), maps.Clone(visited), true)
			}
		case '#', 'O':
			stack.Push(next)
			direction = (direction + 1) % 4
		case 'X':
			return total
		}
	}
	return total
}

// func goldAux(lines []string) int {
// 	visited := make(map[string]struct{})
// 	initPosition := findGuardPosition(lines)

// 	direction := 0
// 	key := func(p grid.Point) string {
// 		return fmt.Sprintf("%s-%d", p.String(), direction)
// 	}

// 	stack := new(generic.Stack[grid.Point])
// 	stack.Push(initPosition)

// 	for !stack.Empty() {
// 		next := stack.Pop()
// 		nextKey := key(next)
// 		if _, isCycle := visited[nextKey]; isCycle {
// 			return 1
// 		}
// 		visited[nextKey] = struct{}{}

// 		v := peak(next, lines, direction)
// 		switch v {
// 		case '^':
// 			stack.Push(nextPoint(next, direction))
// 		case '.':
// 			stack.Push(nextPoint(next, direction))
// 		case '#', 'O':
// 			stack.Push(next)
// 			direction = (direction + 1) % 4
// 		case 'X':
// 			return 0
// 		}
// 	}
// 	return 0
// }

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
