package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/gqgs/AoC2021/generic"
)

type Point struct {
	X, Y int
}

func (p Point) IsTouching(p2 *Point) bool {
	return generic.Abs(p.X-p2.X) <= 1 && generic.Abs(p.Y-p2.Y) <= 1
}

func (p Point) AdjacentPoints(head *Point) []Point {
	if p.X == head.X || p.Y == head.Y {
		return []Point{
			{X: p.X - 1, Y: p.Y},
			{X: p.X + 1, Y: p.Y},
			{X: p.X, Y: p.Y - 1},
			{X: p.X, Y: p.Y + 1},
		}
	}
	return []Point{
		{X: p.X - 1, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y - 1},
	}
}

func (p Point) Distance(p2 *Point) float64 {
	width := math.Abs(float64(p.X) - float64(p2.X))
	height := math.Abs(float64(p.Y) - float64(p2.Y))
	return math.Sqrt(width*width + height*height)
}

func (p Point) PositionKey() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func silver(moves []string) int {
	return countVisited(moves, 2)
}

func gold(moves []string) int {
	return countVisited(moves, 10)
}

func countVisited(moves []string, knots int) int {
	visited := make(map[string]struct{})
	rope := make([]*Point, 0, knots)
	for i := 0; i < cap(rope); i++ {
		rope = append(rope, new(Point))
	}

	head := rope[0]
	tail := rope[knots-1]
	visited[tail.PositionKey()] = struct{}{}
	for _, move := range moves {
		var direction rune
		var distance int
		var moveFunc func(p *Point)
		fmt.Sscanf(move, "%c %d", &direction, &distance)
		switch direction {
		case 'R':
			moveFunc = func(p *Point) { p.X++ }
		case 'L':
			moveFunc = func(p *Point) { p.X-- }
		case 'U':
			moveFunc = func(p *Point) { p.Y++ }
		case 'D':
			moveFunc = func(p *Point) { p.Y-- }
		}
		for ; distance > 0; distance-- {
			moveFunc(head)
			for i := 1; i < knots; i++ {
				if rope[i].IsTouching(rope[i-1]) {
					continue
				}
				dstAdjPoint := Point{}
				minAdjDistance := math.Inf(0)
				for _, adj := range rope[i].AdjacentPoints(rope[i-1]) {
					dist := adj.Distance(rope[i-1])
					if dist < minAdjDistance {
						minAdjDistance = dist
						dstAdjPoint = adj
					}
				}
				rope[i].X = dstAdjPoint.X
				rope[i].Y = dstAdjPoint.Y
			}
			visited[tail.PositionKey()] = struct{}{}
		}
	}
	return len(visited)
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var moves []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		moves = append(moves, next)
	}

	println("silver:", silver(moves))
	println("gold:", gold(moves))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
