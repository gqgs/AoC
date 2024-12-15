package main

import (
	"bufio"
	"os"
	"slices"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func silver(lines []string) int {
	stack := new(generic.Stack[grid.Point])
	for _, point := range grid.Points(lines, '.') {
		stack.Push(point)
	}

	perimiter := make(map[string]int)
	visited := make(generic.Set[grid.Point])
	for !stack.Empty() {
		var area int
		next := stack.Pop()
		if visited.Contains(next) {
			continue
		}
		visited.Add(next)

		area++
		sameStack := new(generic.Stack[grid.Point])
		sameStack.Push(next)

		nodePerimiter := 4
		for nb := range next.UpRightDownLeft() {
			if lines[next.X][next.Y] == lines[nb.X][nb.Y] {
				nodePerimiter--
				sameStack.Push(nb)
			}
		}

		for !sameStack.Empty() {
			same := sameStack.Pop()
			if visited.Contains(same) {
				continue
			}
			visited.Add(same)
			area++

			nodePerimiter += 4
			for nb := range same.UpRightDownLeft() {
				if lines[same.X][same.Y] == lines[nb.X][nb.Y] {
					nodePerimiter--
					sameStack.Push(nb)
				}
			}
		}

		perimiter[string(lines[next.X][next.Y])] += area * nodePerimiter

	}

	var total int
	for _, value := range perimiter {
		total += value
	}

	return total
}

func gold(lines []string) int {
	stack := new(generic.Stack[grid.Point])
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '.' {
				continue
			}
			stack.Push(grid.Point{
				X: i,
				Y: j,
			})
		}
	}

	perimiter := make(map[string]int)
	visited := make(generic.Set[grid.Point])
	for !stack.Empty() {
		var area int
		next := stack.Pop()
		if visited.Contains(next) {
			continue
		}
		visited.Add(next)

		area++
		sameStack := new(generic.Stack[grid.Point])
		sameStack.Push(next)

		sides := make(map[grid.Point][]grid.Point)
		for nb := range next.UpRightDownLeft() {
			if lines[next.X][next.Y] == lines[nb.X][nb.Y] {
				sameStack.Push(nb)
			} else {
				sides[nb] = append(sides[nb], next)
			}
		}

		for !sameStack.Empty() {
			same := sameStack.Pop()
			if visited.Contains(same) {
				continue
			}
			visited.Add(same)
			area++

			for nb := range same.UpRightDownLeft() {
				if lines[same.X][same.Y] == lines[nb.X][nb.Y] {
					sameStack.Push(nb)
				} else {
					sides[nb] = append(sides[nb], same)
				}
			}
		}

		contains := func(n, p grid.Point) bool {
			nextOriginPoints, ok := sides[n]
			if !ok {
				return false
			}

			originPoints, ok := sides[p]
			if !ok {
				return false
			}
			for _, nextOrigPoint := range nextOriginPoints {
				for _, origPoint := range originPoints {
					if grid.StraightPathExists(nextOrigPoint, origPoint, lines) {
						return true
					}
				}
			}
			return false
		}

		deletePoint := func(n, p grid.Point) {
			nextOriginPoints, ok := sides[n]
			if !ok {
				return
			}

			originPoints, ok := sides[p]
			if !ok {
				return
			}
			for _, nextOrigPoint := range nextOriginPoints {
				for index, origPoint := range originPoints {
					if grid.StraightPathExists(nextOrigPoint, origPoint, lines) {
						sides[p] = slices.Delete(sides[p], index, index+1)
						return
					}
				}
			}
		}

		var numberSides int
		for len(sides) > 0 {
			for next := range sides {
				left := grid.Point{X: next.X, Y: next.Y - 1}
				right := grid.Point{X: next.X, Y: next.Y + 1}
				if contains(next, left) || contains(next, right) {
					// horizontal deletion
					for contains(next, left) {
						deletePoint(next, left)
						if len(sides[left]) == 0 {
							delete(sides, left)
						}
						left.Y--
					}
					for contains(next, right) {
						deletePoint(next, right)
						if len(sides[right]) == 0 {
							delete(sides, right)
						}
						right.Y++
					}
				}
				up := grid.Point{X: next.X - 1, Y: next.Y}
				down := grid.Point{X: next.X + 1, Y: next.Y}
				if contains(next, up) || contains(next, down) {
					// vertical deletion
					for contains(next, up) {
						deletePoint(next, up)
						if len(sides[up]) == 0 {
							delete(sides, up)
						}
						up.X--
					}
					for contains(next, down) {
						deletePoint(next, down)
						if len(sides[down]) == 0 {
							delete(sides, down)
						}
						down.X++
					}
				}

				numberSides++
				deletePoint(next, next)
				if len(sides[next]) == 0 {
					delete(sides, next)
				}
			}
		}
		perimiter[string(lines[next.X][next.Y])] += area * numberSides

	}

	var total int
	for _, value := range perimiter {
		total += value
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

	lines = grid.Fill(lines, ".", 1)

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
