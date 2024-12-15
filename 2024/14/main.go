package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gqgs/AoC2021/grid"
)

type Robot struct {
	Position grid.Point
	Velocity grid.Point
}

func (r *Robot) Update(wide, tall int) {
	r.Position.X = (r.Position.X + r.Velocity.X) % wide
	if r.Position.X < 0 {
		r.Position.X += wide
	}

	r.Position.Y = (r.Position.Y + r.Velocity.Y) % tall
	if r.Position.Y < 0 {
		r.Position.Y += tall
	}
}

func draw(robots []*Robot, wide, tall int) bool {
	count := make(map[string]int)
	for _, robot := range robots {
		count[robot.Position.String()]++
	}

	var maxConnectedHorizontal int
	var maxConnectedVertical int

	for _, robot := range robots {
		var connectedHorizontal int
		horizontalIter := grid.Point{X: robot.Position.X, Y: robot.Position.Y}
		for {
			if _, ok := count[horizontalIter.String()]; !ok {
				break
			}
			connectedHorizontal++
			horizontalIter.X++
		}

		var connectedVertial int
		vertitalIter := grid.Point{X: robot.Position.X, Y: robot.Position.Y}
		for {
			if _, ok := count[vertitalIter.String()]; !ok {
				break
			}
			connectedVertial++
			vertitalIter.Y++
		}

		if connectedHorizontal > 0 && connectedHorizontal == connectedVertial {
			maxConnectedHorizontal = max(maxConnectedHorizontal, connectedHorizontal)
			maxConnectedVertical = max(maxConnectedVertical, connectedVertial)
		}
	}

	if maxConnectedHorizontal < 5 {
		return false
	}
	if maxConnectedHorizontal != maxConnectedVertical {
		return false
	}

	for y := range tall {
		for x := range wide {
			if val, ok := count[grid.Point{X: x, Y: y}.String()]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return true
}

func silver(lines []string) int {
	wide := 101
	tall := 103
	iterations := 100

	var robots []*Robot
	for _, line := range lines {
		robot := new(Robot)
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.Position.X, &robot.Position.Y, &robot.Velocity.X, &robot.Velocity.Y)
		robots = append(robots, robot)
	}

	for range iterations {
		for _, robot := range robots {
			robot.Update(wide, tall)
		}
	}

	var q1, q2, q3, q4 int
	for _, robot := range robots {
		if robot.Position.X == (wide-1)/2 {
			continue
		}
		if robot.Position.Y == (tall-1)/2 {
			continue
		}
		switch {
		case robot.Position.X < (wide-1)/2 && robot.Position.Y < (tall-1)/2:
			q1++
		case robot.Position.X > (wide-1)/2 && robot.Position.Y < (tall-1)/2:
			q2++
		case robot.Position.X < (wide-1)/2 && robot.Position.Y > (tall-1)/2:
			q3++
		case robot.Position.X > (wide-1)/2 && robot.Position.Y > (tall-1)/2:
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func gold(lines []string) int {
	wide := 101
	tall := 103

	var robots []*Robot
	for _, line := range lines {
		robot := new(Robot)
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.Position.X, &robot.Position.Y, &robot.Velocity.X, &robot.Velocity.Y)
		robots = append(robots, robot)
	}

	for i := 0; ; i++ {
		for _, robot := range robots {
			robot.Update(wide, tall)
		}

		if draw(robots, wide, tall) {
			fmt.Println("iteration", i)
			time.Sleep(time.Second)
			break
		}
	}

	return 0
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

	if err := scanner.Err(); err != nil {
		return err
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
