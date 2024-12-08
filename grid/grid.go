package grid

import (
	"fmt"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) IsOutBounds(limit int) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}

	if p.X >= limit || p.Y >= limit {
		return true
	}

	return false
}

func (p Point) VerticalDist(p2 Point) int {
	return generic.Abs(p.X - p2.X)
}

func (p Point) HorizontalDist(p2 Point) int {
	return generic.Abs(p.Y - p2.Y)
}

func (p Point) Distance(p2 Point) int {
	return p.VerticalDist(p2) + p.HorizontalDist(p2)
}

func (p Point) AdjacentFunc(fn func(Point)) {
	for _, ap := range p.adjacent() {
		fn(Point{
			X: ap[0],
			Y: ap[1],
		})
	}
}

func (p Point) adjacent() [][2]int {
	return [][2]int{
		{p.X - 1, p.Y - 1},
		{p.X - 1, p.Y},
		{p.X - 1, p.Y + 1},
		{p.X, p.Y - 1},
		{p.X, p.Y},
		{p.X, p.Y + 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y + 1},
	}
}

func Fill(lines []string, fillChar string, appendSize int) []string {
	result := make([]string, 0, len(lines)+appendSize)
	fillLine := strings.Repeat(fillChar, len(lines[0])+appendSize)

	result = append(result, fillLine)
	for _, line := range lines {
		result = append(result, fillChar+line+fillChar)
	}
	result = append(result, fillLine)
	return result
}
