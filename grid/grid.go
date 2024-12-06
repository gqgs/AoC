package grid

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
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
