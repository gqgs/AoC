package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func silver(lines []string) int {
	var total int
	for i := range lines {
		s := new(generic.Stack[grid.Point])
		visited := make(map[string]struct{})
		number := new(strings.Builder)
		for j := range lines[i] {
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				newPoint := grid.Point{
					X: i,
					Y: j,
				}
				s.Push(newPoint)
				visited[newPoint.String()] = struct{}{}
				number.WriteByte(lines[i][j])
				continue
			}

			var valid bool
			for !s.Empty() {
				p := s.Pop()
				for ap := range p.Around() {
					if _, ok := visited[ap.String()]; ok {
						continue
					}
					if lines[ap.X][ap.Y] != '.' {
						valid = true
					}
				}
			}

			if valid {
				digits, _ := strconv.Atoi(number.String())
				total += digits
			}
			number.Reset()
		}
	}
	return total
}

func gold(lines []string) int {
	var total int
	valid := make(map[string][]int)
	for i := range lines {
		s := new(generic.Stack[grid.Point])
		visited := make(map[string]struct{})
		number := new(strings.Builder)
		for j := range lines[i] {
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				newPoint := grid.Point{
					X: i,
					Y: j,
				}
				s.Push(newPoint)
				visited[newPoint.String()] = struct{}{}
				number.WriteByte(lines[i][j])
				continue
			}

			for !s.Empty() {
				p := s.Pop()
				for ap := range p.Around() {
					if _, ok := visited[ap.String()]; ok {
						continue
					}
					if lines[ap.X][ap.Y] == '*' {
						digits, _ := strconv.Atoi(number.String())
						valid[ap.String()] = append(valid[ap.String()], digits)
					}
				}
			}

			number.Reset()
		}
	}

	for _, value := range valid {
		sort.Ints(value)
		value = slices.Compact(value)
		if len(value) == 2 {
			total += value[0] * value[1]
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

	lines = grid.Fill(lines, ".", 2)

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
