package main

import (
	"bufio"
	"container/heap"
	"os"
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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
	basins := IntHeap{0, 0, 0}
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
			if size > basins[0] {
				heap.Pop(&basins)
				heap.Push(&basins, size)
			}
			mu.Unlock()
		}(point)
	}
	wg.Wait()

	return basins[0] * basins[1] * basins[2]
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
