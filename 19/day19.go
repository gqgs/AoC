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

func transforms(x, y, z int) [24][3]int {
	var rotations [24][3]int
	rotations[0] = [3]int{x, y, z}
	rotations[1] = [3]int{x, -y, -z}
	rotations[2] = [3]int{x, z, -y}
	rotations[3] = [3]int{x, -z, y}

	rotations[4] = [3]int{y, x, -z}
	rotations[5] = [3]int{y, -x, z}
	rotations[6] = [3]int{y, z, x}
	rotations[7] = [3]int{y, -z, -x}

	rotations[8] = [3]int{z, x, y}
	rotations[9] = [3]int{z, -x, -y}
	rotations[10] = [3]int{z, y, -x}
	rotations[11] = [3]int{z, -y, x}

	rotations[12] = [3]int{-x, y, -z}
	rotations[13] = [3]int{-x, -y, z}
	rotations[14] = [3]int{-x, z, y}
	rotations[15] = [3]int{-x, -z, -y}

	rotations[16] = [3]int{-y, x, z}
	rotations[17] = [3]int{-y, -x, -z}
	rotations[18] = [3]int{-y, z, -x}
	rotations[19] = [3]int{-y, -z, x}

	rotations[20] = [3]int{-z, x, -y}
	rotations[21] = [3]int{-z, -x, y}
	rotations[22] = [3]int{-z, y, x}
	rotations[23] = [3]int{-z, -y, -x}

	return rotations
}

func overlap(s0, s1 [][3]int) ([][3]int, [3]int) {
	overlaps := make(map[string]int)
	for _, p1 := range s1 {
		for index, t1 := range transforms(p1[0], p1[1], p1[2]) {
			for _, p0 := range s0 {
				x1, y1, z1 := t1[0], t1[1], t1[2]
				x2, y2, z2 := p0[0], p0[1], p0[2]
				x, y, z := x2-x1, y2-y1, z2-z1
				key := fmt.Sprint(x, y, z)
				overlaps[key]++
				if overlaps[key] == 12 {
					var normalizedPoints [][3]int
					for _, p := range s1 {
						t := transforms(p[0], p[1], p[2])[index]
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
			result, scanner := overlap(s, next)
			if len(result) > 0 {
				transformed = append(transformed, result)
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

	println("points:", len(points))

	var maxDistance int
	for i := 0; i < len(scanners); i++ {
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}
			maxDistance = generic.Max(maxDistance, distance(scanners[i], scanners[j]))
		}
	}

	println("maxDistance:", maxDistance, len(scanners))

	return nil
}

func distance(p0, p1 [3]int) int {
	return generic.Abs(p0[0]-p1[0]) + generic.Abs(p0[1]-p1[1]) + generic.Abs(p0[2]-p1[2])

}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
