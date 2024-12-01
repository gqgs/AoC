package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

type Scanner struct {
	Name   string
	Points [][3]int
}

func transform(x, y, z, r int) [3]int {
	return [24][3]int{
		{x, y, z},
		{x, -y, -z},
		{x, z, -y},
		{x, -z, y},

		{y, x, -z},
		{y, -x, z},
		{y, z, x},
		{y, -z, -x},

		{z, x, y},
		{z, -x, -y},
		{z, y, -x},
		{z, -y, x},

		{-x, y, -z},
		{-x, -y, z},
		{-x, z, y},
		{-x, -z, -y},

		{-y, x, z},
		{-y, -x, -z},
		{-y, z, -x},
		{-y, -z, x},

		{-z, x, -y},
		{-z, -x, y},
		{-z, y, x},
		{-z, -y, -x},
	}[r]
}

func transforms(x, y, z int) [24][3]int {
	var rotations [24][3]int
	for r := 0; r < 24; r++ {
		rotations[r] = transform(x, y, z, r)
	}
	return rotations
}

func hash(x, y, z int) uint {
	// FNV-1 hash
	var prime uint = 1099511628211
	var hash uint = 14695981039346656037
	for _, n := range []int{x, y, z} {
		hash *= prime
		hash ^= uint(n)
	}
	return hash
}

func overlap(s0, s1 [][3]int) ([][3]int, [3]int) {
	overlaps := make(map[uint]int)
	for _, p1 := range s1 {
		for index, t1 := range transforms(p1[0], p1[1], p1[2]) {
			for _, p0 := range s0 {
				x1, y1, z1 := t1[0], t1[1], t1[2]
				x2, y2, z2 := p0[0], p0[1], p0[2]
				x, y, z := x2-x1, y2-y1, z2-z1
				key := hash(x, y, z)
				overlaps[key]++
				if overlaps[key] == 12 {
					var normalizedPoints [][3]int
					for _, p := range s1 {
						t := transform(p[0], p[1], p[2], index)
						normalizedPoints = append(normalizedPoints, [3]int{t[0] + x, t[1] + y, t[2] + z})
					}
					return normalizedPoints, [3]int{x, y, z}
				}
			}
		}
	}
	return nil, [3]int{}
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var scanner []Scanner
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		scannerName := reader.Text()
		var points [][3]int
		for {
			reader.Scan()
			coord := reader.Text()
			if len(coord) == 0 {
				scanner = append(scanner, Scanner{
					Name:   scannerName,
					Points: points,
				})
				break
			}
			c := strings.Split(coord, ",")
			x, _ := strconv.Atoi(c[0])
			y, _ := strconv.Atoi(c[1])
			z, _ := strconv.Atoi(c[2])
			points = append(points, [3]int{x, y, z})
		}
	}

	transformed := make([][][3]int, 0)
	transformed = append(transformed, scanner[0].Points)
	queue := make(generic.Queue[[][3]int], 0)
	for _, s := range scanner[1:] {
		queue.Push(s.Points)
	}

	var scanners [][3]int
	scanners = append(scanners, [3]int{0, 0, 0})
Next:
	for len(queue) > 0 {
		next := queue.Pop()
		for _, s := range transformed {
			normalized, scanner := overlap(s, next)
			if len(normalized) > 0 {
				transformed = append(transformed, normalized)
				scanners = append(scanners, scanner)
				continue Next
			}
		}
		queue.Push(next)
	}

	points := make(map[string]struct{})
	for _, t := range transformed {
		for _, p := range t {
			points[fmt.Sprint(p[0], p[1], p[2])] = struct{}{}
		}
	}

	println("silver:", len(points))

	var maxDistance int
	for i := 0; i < len(scanners); i++ {
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}
			maxDistance = max(maxDistance, manhattanDistance(scanners[i], scanners[j]))
		}
	}

	println("gold:", maxDistance)

	return nil
}

func manhattanDistance(p0, p1 [3]int) int {
	return generic.Abs(p0[0]-p1[0]) + generic.Abs(p0[1]-p1[1]) + generic.Abs(p0[2]-p1[2])

}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
