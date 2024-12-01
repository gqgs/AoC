package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Step struct {
	On     bool
	X0, X1 int
	Y0, Y1 int
	Z0, Z1 int
	Sign   int
}

func gold(input []string) int {
	var steps []Step
	for _, step := range input {
		var minx, maxx, miny, maxy, minz, maxz int
		if strings.HasPrefix(step, "on") {
			fmt.Sscanf(step, "on x=%d..%d,y=%d..%d,z=%d..%d)", &minx, &maxx, &miny, &maxy, &minz, &maxz)
			steps = append(steps, Step{
				On: true,
				X0: minx,
				X1: maxx,
				Y0: miny,
				Y1: maxy,
				Z0: minz,
				Z1: maxz,
			})
		} else {
			fmt.Sscanf(step, "off x=%d..%d,y=%d..%d,z=%d..%d)", &minx, &maxx, &miny, &maxy, &minz, &maxz)
			steps = append(steps, Step{
				On: false,
				X0: minx,
				X1: maxx,
				Y0: miny,
				Y1: maxy,
				Z0: minz,
				Z1: maxz,
			})
		}
	}

	var totalVolume int
	var cuboids []Step
	for _, step := range steps {
		var intersections []Step
		for _, cuboid := range cuboids {
			intersection, exists := intersect(step, cuboid)
			if !exists {
				continue
			}

			intersection.Sign = -cuboid.Sign
			intersections = append(intersections, intersection)
			totalVolume += cuboid.Sign * volume(intersection)
		}

		cuboids = append(cuboids, intersections...)
		if step.On {
			step.Sign = -1
			cuboids = append(cuboids, step)
			totalVolume += volume(step)
		}
	}

	return totalVolume
}

func silver(steps []string) int {
	volume := make(map[int]map[int]map[int]struct{})
	for _, step := range steps {
		var minx, maxx, miny, maxy, minz, maxz int
		if strings.HasPrefix(step, "on") {
			fmt.Sscanf(step, "on x=%d..%d,y=%d..%d,z=%d..%d)", &minx, &maxx, &miny, &maxy, &minz, &maxz)
			for x := minx; x <= maxx; x++ {
				for y := miny; y <= maxy; y++ {
					for z := minz; z <= maxz; z++ {
						if _, exists := volume[x]; !exists {
							volume[x] = make(map[int]map[int]struct{})
						}
						if _, exists := volume[x][y]; !exists {
							volume[x][y] = make(map[int]struct{})
						}
						if _, exists := volume[x][y][z]; !exists {
							volume[x][y][z] = struct{}{}
						}
					}
				}
			}
		} else {
			fmt.Sscanf(step, "off x=%d..%d,y=%d..%d,z=%d..%d)", &minx, &maxx, &miny, &maxy, &minz, &maxz)
			for x := minx; x <= maxx; x++ {
				for y := miny; y <= maxy; y++ {
					for z := minz; z <= maxz; z++ {
						delete(volume[x][y], z)
					}
				}
			}
		}
	}

	var count int
	for x := range volume {
		for y := range volume[x] {
			count += len(volume[x][y])
		}
	}

	return count
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
		input = append(input, scanner.Text())
	}

	println("silver:", silver(input))
	println("gold:", gold(input))

	return nil
}

func volume(s Step) int {
	return (s.X1 - s.X0 + 1) * (s.Y1 - s.Y0 + 1) * (s.Z1 - s.Z0 + 1)
}

func intersect(s0, s1 Step) (Step, bool) {
	var s Step
	s.X0 = max(s0.X0, s1.X0)
	s.X1 = min(s0.X1, s1.X1)

	s.Y0 = max(s0.Y0, s1.Y0)
	s.Y1 = min(s0.Y1, s1.Y1)

	s.Z0 = max(s0.Z0, s1.Z0)
	s.Z1 = min(s0.Z1, s1.Z1)

	return s, s.X0 <= s.X1 && s.Y0 <= s.Y1 && s.Z0 <= s.Z1
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
