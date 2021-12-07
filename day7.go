package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func silver(hPositions map[int]int) int {
	minCost := 1<<63 - 1
	for hPosition := range hPositions {
		var cost int
		for position, nCrabs := range hPositions {
			cost += abs(hPosition-position) * nCrabs
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func gold(hPositions map[int]int) int {
	minX := 1<<63 - 1
	maxX := 0

	for position := range hPositions {
		if position < minX {
			minX = position
		}
		if position > maxX {
			maxX = position
		}
	}

	moveCost := func(d int) int {
		return d * (d + 1) / 2
	}

	minCost := 1<<63 - 1
	for x := minX; x <= maxX; x++ {
		var cost int
		for position, nCrabs := range hPositions {
			cost += moveCost(abs(x-position)) * nCrabs
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func solve() error {
	file, err := os.Open("day7")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	crabs := strings.Split(scanner.Text(), ",")
	hPositions := make(map[int]int)
	for _, c := range crabs {
		h, _ := strconv.Atoi(c)
		hPositions[h]++
	}

	println("silver:", silver(hPositions))
	println("gold:", gold(hPositions))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
