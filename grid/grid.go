package grid

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

type Grid [][]rune

func ParseLines(lines []string) Grid {
	state := make([][]rune, len(lines))
	for i, s := range lines {
		for _, r := range s {
			state[i] = append(state[i], r)
		}
	}
	return state
}

func NewSquared(size int) Grid {
	state := make(Grid, size+2)
	for i := range size + 2 {
		for range size + 2 {
			state[i] = append(state[i], '.')
		}
	}
	return state
}

func (g Grid) FillPerimeter() {
	size := len(g) - 2
	for i := range size + 2 {
		for j := range size + 2 {
			if i == 0 || i == size+1 || j == 0 || j == size+1 {
				g[i][j] = '#'
			}
		}
	}
}

func (g Grid) FindPosition(char rune) Point {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == char {
				return Point{
					X: i,
					Y: j,
				}
			}
		}
	}

	panic("char not found")
}

func (g Grid) String() string {
	var builder strings.Builder
	for i := range g {
		for j := range g[i] {
			builder.WriteRune(g[j][i])
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (g Grid) StraightPathExists(p1, p2 Point) bool {
	c := g[p1.X][p1.Y]

	minx := min(p1.X, p2.X)
	maxx := max(p1.X, p2.X)
	if p1.Y == p2.Y {
		var x int
		for x = minx; x < maxx && g[x][p1.Y] == c; x++ {
		}
		if x == maxx {
			return true
		}
	}

	miny := min(p1.Y, p2.Y)
	maxy := max(p1.Y, p2.Y)
	if p1.X == p2.X {
		var y int
		for y = miny; y < maxy && g[p1.X][y] == c; y++ {

		}
		if y == maxy {
			return true
		}
	}

	return false
}

type Point struct {
	X, Y  int
	Score int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) Equal(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
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

func (p Point) Around() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for _, ap := range p.around() {
			if !yield(Point{
				X: ap[0],
				Y: ap[1],
			}) {
				return
			}
		}
	}
}

func (p Point) around() [][2]int {
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

func (p Point) UpRightDownLeft() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for _, ap := range p.upRightDownLeft() {
			if !yield(Point{
				X: ap[0],
				Y: ap[1],
			}) {
				return
			}
		}
	}
}

func (p Point) upRightDownLeft() [][2]int {
	return [][2]int{
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X + 1, p.Y},
		{p.X, p.Y - 1},
	}
}

func Points(lines []string, ignoreChar byte) []Point {
	var points []Point
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == ignoreChar {
				continue
			}
			p := Point{
				X: i,
				Y: j,
			}
			points = append(points, p)
		}
	}
	return points
}

func Fill(lines []string, fillChar string, appendSize int) []string {
	result := make([]string, 0, len(lines)+appendSize)
	fillLine := strings.Repeat(fillChar, len(lines[0])+appendSize*2)

	result = append(result, fillLine)
	for _, line := range lines {
		str := strings.Repeat(fillChar, appendSize)
		result = append(result, str+line+str)
	}
	result = append(result, fillLine)
	return result
}

func Replaced(lines []string, c string, x, y int) []string {
	clone := slices.Clone(lines)
	clone[x] = clone[x][:y] + c + clone[x][y+1:]
	return clone
}
