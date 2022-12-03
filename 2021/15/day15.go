package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/gqgs/AoC2021/generic"
)

type Node struct {
	x, y int
	cost int
}

func gold(cost, risk, inqueue [][]int, maxX, maxY int) int {
	queue := generic.NewMinHeap(func(e1, e2 *Node) bool {
		return e1.cost < e2.cost
	})
	queue.Push(&Node{0, 0, 0})

	for queue.Len() > 0 {
		next := queue.Pop()
		x, y, c := next.x, next.y, next.cost
		if x == maxX-1 && y == maxY-1 {
			return c
		}

		if cost[x][y] == 0 {
			cost[x][y] = c
		}

		if c > cost[x][y] {
			continue
		}

		for _, p := range [][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}} {
			dx, dy := p[0], p[1]
			if dx < 0 || dx >= maxX || dy < 0 || dy >= maxY {
				continue
			}

			if inqueue[dx][dy] == 0 {
				queue.Push(&Node{dx, dy, c + risk[dx][dy]})
				inqueue[dx][dy] = 1
			}
		}
	}
	return -1
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

	maxX := len(input) * 5
	maxY := len(input[0]) * 5

	risk := make([][]int, maxX)
	cost := make([][]int, maxX)
	inqueue := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		risk[i] = make([]int, maxY)
		cost[i] = make([]int, maxY)
		inqueue[i] = make([]int, maxY)
	}

	for y, line := range input {
		for x, n := range line {
			risk[x][y], _ = strconv.Atoi(string(n))
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if i == 0 && j == 0 {
						continue
					}
					value := (risk[x][y] + i + j)
					if value > 9 {
						value = (value % 10) + 1
					}
					risk[x+i*len(line)][y+j*len(input)] = value
				}
			}
		}
	}

	println("gold:", gold(cost, risk, inqueue, maxX, maxY))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
