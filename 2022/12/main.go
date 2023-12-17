package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

var (
	posInfinity = math.Inf(0)
	negInfinity = math.Inf(-1)
)

type Point struct {
	X, Y     int
	Neighbor []*Point
	Dist     float64
	Value    byte
}

func (p Point) String() string {
	return fmt.Sprintf("%d_%d", p.X, p.Y)
}

func (p Point) ElevationDifference(p2 Point) int {
	// return generic.Abs(p.Elevation() - p2.Elevation())
	diff := p.Elevation() - p2.Elevation()
	if diff >= -1 {
		return 1
	}
	return 999999
}

func (p Point) Elevation() int {
	switch p.Value {
	case 'S':
		return int('a')
	case 'E':
		return int('z')
	default:
		return int(p.Value)
	}
}

func NewPoint(x, y int, value byte) *Point {
	return &Point{
		X:     x,
		Y:     y,
		Dist:  posInfinity,
		Value: value,
	}
}

type Graph map[string]*Point

func (g *Graph) Point(x, y int, value byte) *Point {
	key := fmt.Sprintf("%d_%d", x, y)
	if (*g)[key] == nil {
		(*g)[key] = NewPoint(x, y, value)
	}
	return (*g)[key]
}

func silver(input []string) int {
	graph := make(Graph)
	heap := generic.NewMinHeap(func(e1, e2 *Point) bool {
		return e1.Dist < e2.Dist
	})
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			p := graph.Point(i, j, input[i][j])
			if input[i][j] == 'S' {
				p.Dist = 0
				// graph.Update(p)
				// fmt.Println(p, p.X, p.Y, p.Dist)
				heap.Push(p)
			}
			for _, ns := range [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
				p.Neighbor = append(p.Neighbor, graph.Point(ns[0], ns[1], input[ns[0]][ns[1]]))
			}
		}
	}

	inQueue := make(map[string]struct{})
	for heap.Len() > 0 {
		next := heap.Pop()
		// delete(inQueue, next.String())
		if input[next.X][next.Y] == 'E' {
			return int(next.Dist)
		}

		// println("next.Neighbor", next.X, next.Y, len(next.Neighbor))
		for _, n := range next.Neighbor {
			// println("diff", next.X, next.Y, n.X, n.Y, next.ElevationDifference(*n))
			if next.ElevationDifference(*n) > 1 {
				continue
			}
			if _, ok := inQueue[n.String()]; ok {
				continue
			}
			// println("n.Dist", next.Dist, next.X, next.Y, n.X, n.Y)
			n.Dist = next.Dist + 1
			// time.Sleep(time.Second)
			heap.Push(n)
			inQueue[n.String()] = struct{}{}
		}
	}

	panic("not found")
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		input = append(input, "~"+next+"~")
	}
	padding := strings.Repeat("~", len(input[0]))
	input = append(input, padding)
	input = append([]string{padding}, input...)

	println("silver:", silver(input))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
