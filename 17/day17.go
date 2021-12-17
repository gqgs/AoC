package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gqgs/AoC2021/generic"
)

func isInt(v float64) bool {
	return v == float64(int(v))
}

func gold(minX, maxX, minY, maxY, maxV0 int) int {
	solutions := make(map[string]struct{})
	for startVelocityX := 0; startVelocityX <= maxX; startVelocityX++ {
		velocityX := startVelocityX
		positionX := 0
		t := 0
		for {
			t++
			positionX += velocityX
			if positionX >= minX && positionX <= maxX {
				update := func(y, t int) {
					v0 := float64(2*y+t*t-t) / float64(2*t)
					if !isInt(v0) {
						return
					}
					solutions[fmt.Sprint(startVelocityX, int(v0))] = struct{}{}
				}

				if velocityX == 0 {
					for i := t; i <= 2*maxV0; i++ {
						for y := minY; y <= maxY; y++ {
							update(y, i)
						}
					}
					break
				}

				for y := minY; y <= maxY; y++ {
					update(y, t)
				}
			}

			if velocityX == 0 {
				break
			}
			velocityX--
		}
	}

	return len(solutions)
}

func silver(minX, maxX, minY, maxY int) (int, int) {
	checked := make(map[int]struct{})
	var globalYMax float64
	var globalMaxYV0 int
	for startVelocityX := 0; startVelocityX < maxX; startVelocityX++ {
		velocityX := startVelocityX
		positionX := 0
		t := 0
		for {
			t++
			positionX += velocityX
			if positionX >= minX && positionX <= maxX {
				if _, alreadyChecked := checked[t]; alreadyChecked {
					continue
				}
				checked[t] = struct{}{}
				for y := minY; y < maxY; y++ {
					v0 := float64(2*y+t*t-t) / float64(2*t)
					if !isInt(v0) {
						continue
					}
					localYMax := (v0 + 0.5) * (v0 + 0.5) * 0.5
					globalYMax = generic.Max(globalYMax, localYMax)
					globalMaxYV0 = generic.Max(globalMaxYV0, generic.Abs(int(v0)))
				}

			}

			if velocityX == 0 {
				break
			}
			velocityX--
		}
	}

	return int(globalYMax), int(globalMaxYV0)
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var x0, x1, y0, y1 int
	fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &x0, &x1, &y0, &y1)

	maxY, maxV0 := silver(x0, x1, y0, y1)
	println("silver:", maxY)
	println("gold:", gold(x0, x1, y0, y1, maxV0))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
