package main

import (
	"bufio"
	"os"

	"github.com/gqgs/AoC2021/generic"
	"github.com/gqgs/AoC2021/grid"
)

func possibleAntinodes(p1, p2 grid.Point) (int, int, []grid.Point) {
	vdist := p1.VerticalDist(p2)
	hdist := p1.HorizontalDist(p2)
	return vdist, hdist, []grid.Point{
		{
			X: p1.X - vdist,
			Y: p1.Y - hdist,
		},
		{
			X: p1.X - vdist,
			Y: p1.Y + hdist,
		},
		{
			X: p1.X + vdist,
			Y: p1.Y - hdist,
		},
		{
			X: p1.X + vdist,
			Y: p1.Y + hdist,
		},
	}
}

func isAntinode(c, p1, p2 grid.Point) bool {
	targetDist := p1.Distance(p2)
	return targetDist == c.Distance(p1) && (2*targetDist) == c.Distance(p2)
}

func silver(lines []string) int {
	points := generic.NewSet[grid.Point]()
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '.' {
				continue
			}
			p := grid.Point{
				X: i,
				Y: j,
			}
			points.Add(p)
		}
	}

	limit := len(lines[0])
	antinode := generic.NewSet[grid.Point]()
	for key := range points {
	NextPoint:
		for point := range points {
			if key == point {
				continue
			}

			// don't interfere with each other
			if lines[key.X][key.Y] != lines[point.X][point.Y] {
				continue
			}
			_, _, possible := possibleAntinodes(key, point)
			for _, candidate := range possible {
				if candidate.IsOutBounds(limit) {
					continue
				}

				if isAntinode(candidate, key, point) {
					antinode.Add(candidate)
					// since a co-point can have at most one antinode
					// we can skip to the next point here
					continue NextPoint
				}
			}

		}
	}
	return len(antinode)
}

func gold(lines []string) int {
	antinode := generic.NewSet[grid.Point]()
	points := generic.NewSet[grid.Point]()
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '.' {
				continue
			}
			p := grid.Point{
				X: i,
				Y: j,
			}
			points.Add(p)
			antinode.Add(p)
		}
	}

	limit := len(lines[0])
	for key := range points {
		for point := range points {
			if key == point {
				continue
			}

			// don't interfere with each other
			if lines[key.X][key.Y] != lines[point.X][point.Y] {
				continue
			}
			vdist, hdist, possible := possibleAntinodes(key, point)
			for i, candidate := range possible {
				if isAntinode(candidate, key, point) {
					for !candidate.IsOutBounds(limit) {
						antinode.Add(candidate)
						switch i {
						case 0:
							candidate.X -= vdist
							candidate.Y -= hdist
						case 1:
							candidate.X -= vdist
							candidate.Y += hdist
						case 2:
							candidate.X += vdist
							candidate.Y -= hdist
						case 3:
							candidate.X += vdist
							candidate.Y += hdist
						}
					}
				}
			}

		}
	}

	return len(antinode)
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

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
