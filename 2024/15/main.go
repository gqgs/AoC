package main

import (
	"bufio"
	"os"
	"slices"
	"strings"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func gold(state []string, moves string) int {
	replacer := strings.NewReplacer(
		"#", "##",
		"O", "[]",
		".", "..",
		"@", "@.",
	)

	for i := range state {
		state[i] = replacer.Replace(state[i])
	}

	robot := robotStart(state)

	for _, c := range moves {
		switch c {
		case '<':
			if state[robot.X][robot.Y-1] == '#' {
				continue
			}
			if state[robot.X][robot.Y-1] == '.' {
				state = replace(state, "@", robot.X, robot.Y-1)
				state = replace(state, ".", robot.X, robot.Y)
				robot.Y--
				continue
			}
			if state[robot.X][robot.Y-1] == ']' {
				var found bool
				var y int
				for y = robot.Y - 2; state[robot.X][y] != '#'; y-- {
					if state[robot.X][y] == '.' {
						found = true
						break
					}
				}
				if found {
					for ; y < robot.Y; y += 2 {
						state = replace(state, "[", robot.X, y)
						state = replace(state, "]", robot.X, y+1)
					}
					state = replace(state, "@", robot.X, robot.Y-1)
					state = replace(state, ".", robot.X, robot.Y)
					robot.Y--
					continue
				}
			}
		case '^':
			if state[robot.X-1][robot.Y] == '#' {
				continue
			}
			if state[robot.X-1][robot.Y] == '.' {
				state = replace(state, "@", robot.X-1, robot.Y)
				state = replace(state, ".", robot.X, robot.Y)
				robot.X--
				continue
			}
			if state[robot.X-1][robot.Y] == '[' || state[robot.X-1][robot.Y] == ']' {
				toUpdate := make(generic.Set[grid.Point])
				toCheck := new(generic.Stack[grid.Point])
				if state[robot.X-1][robot.Y] == '[' {
					toCheck.Push(grid.Point{X: robot.X - 1, Y: robot.Y})
				} else {
					toCheck.Push(grid.Point{X: robot.X - 1, Y: robot.Y - 1})
				}
				var found bool
				var invalid bool
				for !toCheck.Empty() && !invalid {
					next := toCheck.Pop()
					toUpdate.Add(next)
					if state[next.X-1][next.Y] == '.' && state[next.X-1][next.Y+1] == '.' {
						found = true
						continue
					}
					if state[next.X-1][next.Y] == '#' || state[next.X-1][next.Y+1] == '#' {
						invalid = true
						break
					}
					if state[next.X-1][next.Y] == '[' {
						toCheck.Push(grid.Point{X: next.X - 1, Y: next.Y})
						continue
					}
					if state[next.X-1][next.Y] == ']' {
						toCheck.Push(grid.Point{X: next.X - 1, Y: next.Y - 1})
					}
					if state[next.X-1][next.Y+1] == '[' {
						toCheck.Push(grid.Point{X: next.X - 1, Y: next.Y + 1})
					}
				}

				if found && !invalid {
					for box := range toUpdate {
						state = replace(state, ".", box.X, box.Y)
						state = replace(state, ".", box.X, box.Y+1)
					}
					for box := range toUpdate {
						state = replace(state, "[", box.X-1, box.Y)
						state = replace(state, "]", box.X-1, box.Y+1)
					}

					state = replace(state, "@", robot.X-1, robot.Y)
					state = replace(state, ".", robot.X, robot.Y)
					robot.X--
				}
				continue
			}
		case 'v':
			if state[robot.X+1][robot.Y] == '#' {
				continue
			}
			if state[robot.X+1][robot.Y] == '.' {
				state = replace(state, "@", robot.X+1, robot.Y)
				state = replace(state, ".", robot.X, robot.Y)
				robot.X++
				continue
			}
			if state[robot.X+1][robot.Y] == '[' || state[robot.X+1][robot.Y] == ']' {
				toUpdate := make(generic.Set[grid.Point])
				toCheck := new(generic.Stack[grid.Point])
				if state[robot.X+1][robot.Y] == '[' {
					toCheck.Push(grid.Point{X: robot.X + 1, Y: robot.Y})
				} else {
					toCheck.Push(grid.Point{X: robot.X + 1, Y: robot.Y - 1})
				}
				var found bool
				var invalid bool
				for !toCheck.Empty() && !invalid {
					next := toCheck.Pop()
					toUpdate.Add(next)
					if state[next.X+1][next.Y] == '.' && state[next.X+1][next.Y+1] == '.' {
						found = true
						continue
					}
					if state[next.X+1][next.Y] == '#' || state[next.X+1][next.Y+1] == '#' {
						invalid = true
						break
					}
					if state[next.X+1][next.Y] == '[' {
						toCheck.Push(grid.Point{X: next.X + 1, Y: next.Y})
						continue
					}
					if state[next.X+1][next.Y] == ']' {
						toCheck.Push(grid.Point{X: next.X + 1, Y: next.Y - 1})
					}
					if state[next.X+1][next.Y+1] == '[' {
						toCheck.Push(grid.Point{X: next.X + 1, Y: next.Y + 1})
					}
				}

				if found && !invalid {
					for box := range toUpdate {
						state = replace(state, ".", box.X, box.Y)
						state = replace(state, ".", box.X, box.Y+1)
					}

					for box := range toUpdate {
						state = replace(state, "[", box.X+1, box.Y)
						state = replace(state, "]", box.X+1, box.Y+1)
					}

					state = replace(state, "@", robot.X+1, robot.Y)
					state = replace(state, ".", robot.X, robot.Y)
					robot.X++
				}
				continue
			}
		case '>':
			if state[robot.X][robot.Y+1] == '#' {
				continue
			}
			if state[robot.X][robot.Y+1] == '.' {
				state = replace(state, "@", robot.X, robot.Y+1)
				state = replace(state, ".", robot.X, robot.Y)
				robot.Y++
				continue
			}
			if state[robot.X][robot.Y+1] == '[' {
				var found bool
				var y int
				for y = robot.Y + 2; state[robot.X][y] != '#'; y++ {
					if state[robot.X][y] == '.' {
						found = true
						break
					}
				}
				if found {
					for ; y > robot.Y; y -= 2 {
						state = replace(state, "]", robot.X, y)
						state = replace(state, "[", robot.X, y-1)
					}
					state = replace(state, "@", robot.X, robot.Y+1)
					state = replace(state, ".", robot.X, robot.Y)
					robot.Y++
					continue
				}
			}
		default:
			panic("invalid char:" + string(c))
		}
	}

	var total int
	for i := range state {
		for j := range state[i] {
			if state[i][j] == '[' {
				total += 100*i + j
			}
		}
	}

	return total
}

func silver(state []string, moves string) int {
	robot := robotStart(state)
	for _, c := range moves {
		switch c {
		case '<':
			if state[robot.X][robot.Y-1] == '#' {
				continue
			}
			if state[robot.X][robot.Y-1] == '.' {
				state = replace(state, "@", robot.X, robot.Y-1)
				state = replace(state, ".", robot.X, robot.Y)
				robot.Y--
				continue
			}
			if state[robot.X][robot.Y-1] == 'O' {
				var found bool
				var y int
				for y = robot.Y - 2; state[robot.X][y] != '#'; y-- {
					if state[robot.X][y] == '.' {
						found = true
						break
					}
				}
				if found {
					for ; y < robot.Y; y++ {
						state = replace(state, "O", robot.X, y)
					}
					state = replace(state, "@", robot.X, robot.Y-1)
					state = replace(state, ".", robot.X, robot.Y)
					robot.Y--
					continue
				}
			}
		case '^':
			if state[robot.X-1][robot.Y] == '#' {
				continue
			}
			if state[robot.X-1][robot.Y] == '.' {
				state = replace(state, "@", robot.X-1, robot.Y)
				state = replace(state, ".", robot.X, robot.Y)
				robot.X--
				continue
			}
			if state[robot.X-1][robot.Y] == 'O' {
				var found bool
				var x int
				for x = robot.X - 2; state[x][robot.Y] != '#'; x-- {
					if state[x][robot.Y] == '.' {
						found = true
						break
					}
				}
				if found {
					for ; x < robot.X; x++ {
						state = replace(state, "O", x, robot.Y)
					}
					state = replace(state, "@", robot.X-1, robot.Y)
					state = replace(state, ".", robot.X, robot.Y)
					robot.X--
					continue
				}
			}
		case 'v':
			if state[robot.X+1][robot.Y] == '#' {
				continue
			}
			if state[robot.X+1][robot.Y] == '.' {
				state = replace(state, "@", robot.X+1, robot.Y)
				state = replace(state, ".", robot.X, robot.Y)
				robot.X++
				continue
			}
			if state[robot.X+1][robot.Y] == 'O' {
				var found bool
				var x int
				for x = robot.X + 2; state[x][robot.Y] != '#'; x++ {
					if state[x][robot.Y] == '.' {
						found = true
						break
					}
				}
				if found {
					for ; x > robot.X; x-- {
						state = replace(state, "O", x, robot.Y)
					}
					state = replace(state, "@", robot.X+1, robot.Y)
					state = replace(state, ".", robot.X, robot.Y)
					robot.X++
					continue
				}
			}
		case '>':
			if state[robot.X][robot.Y+1] == '#' {
				continue
			}
			if state[robot.X][robot.Y+1] == '.' {
				state = replace(state, "@", robot.X, robot.Y+1)
				state = replace(state, ".", robot.X, robot.Y)
				robot.Y++
				continue
			}
			if state[robot.X][robot.Y+1] == 'O' {
				var found bool
				var y int
				for y = robot.Y + 2; state[robot.X][y] != '#'; y++ {
					if state[robot.X][y] == '.' {
						found = true
						break
					}
				}
				if found {
					for ; y > robot.Y; y-- {
						state = replace(state, "O", robot.X, y)
					}
					state = replace(state, "@", robot.X, robot.Y+1)
					state = replace(state, ".", robot.X, robot.Y)
					robot.Y++
					continue
				}
			}
		default:
			panic("invalid char:" + string(c))
		}
	}

	var total int
	for i := range state {
		for j := range state[i] {
			if state[i][j] == 'O' {
				total += 100*i + j
			}
		}
	}

	return total
}

func robotStart(state []string) grid.Point {
	for i := range state {
		for j := range state[i] {
			if state[i][j] == '@' {
				return grid.Point{
					X: i,
					Y: j,
				}
			}
		}
	}
	panic("did not find robot in initial state")
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
		lines = append(lines, next)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	index := slices.Index(lines, "")
	state := lines[:index]
	moves := strings.Join(lines[index:], "")

	println(silver(state, moves))
	println(gold(state, moves))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
