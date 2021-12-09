package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"sync"
)

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func castRune(c byte) int {
	return int(c - '0')
}

func findMinPoints(input []string) [][]int {
	var points [][]int
	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			if input[y][x] < min(min(input[y][x-1], input[y][x+1]), min(input[y-1][x], input[y+1][x])) {
				points = append(points, []int{y, x})
			}
		}
	}
	return points
}

func silver(input []string, points [][]int) int {
	var riskLevel int
	for _, point := range points {
		y, x := point[0], point[1]
		riskLevel += 1 + castRune(input[y][x])
	}
	return riskLevel
}

func gold(input []string, points [][]int) int {
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(len(points))
	var basins []int
	for _, point := range points {
		go func(point []int) {
			defer wg.Done()
			y, x := point[0], point[1]
			size := basinSize(input, [][]int{
				{y, x - 1},
				{y, x + 1},
				{y - 1, x},
				{y + 1, x},
			})
			mu.Lock()
			basins = append(basins, size)
			mu.Unlock()
		}(point)
	}
	wg.Wait()

	sort.Ints(basins)

	l := len(basins)
	return basins[l-1] * basins[l-2] * basins[l-3]
}

func basinSize(input []string, stack [][]int) int {
	var size int
	visited := make(map[int]map[int]struct{})
	for len(stack) > 0 {
		next := stack[0]
		stack = stack[1:]

		y, x := next[0], next[1]
		if input[y][x] == '9' {
			continue
		}

		if visited[y] == nil {
			visited[y] = make(map[int]struct{})
		}
		if _, alreadyVisited := visited[y][x]; alreadyVisited {
			continue
		}
		visited[y][x] = struct{}{}

		size++

		stack = append(stack, [][]int{
			{y, x - 1},
			{y, x + 1},
			{y - 1, x},
			{y + 1, x},
		}...)

	}
	return size
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
		input = append(input, "9"+scanner.Text()+"9")
	}
	padding := strings.Repeat("9", len(input[0]))
	input = append(input, padding)
	input = append([]string{padding}, input...)

	minPoints := findMinPoints(input)
	println("silver:", silver(input, minPoints))
	println("gold:", gold(input, minPoints))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
