package main

import (
	"bufio"
	"fmt"
	"os"
)

func silver(assings [][4]int) int {
	var total int
	for _, assing := range assings {
		startA, endA, startB, endB := assing[0], assing[1], assing[2], assing[3]
		if (startA >= startB && endA <= endB) || (startB >= startA && endB <= endA) {
			total++
		}
	}
	return total
}

func gold(assings [][4]int) int {
	var total int
	for _, assing := range assings {
		startA, endA, startB, endB := assing[0], assing[1], assing[2], assing[3]
		if (startB >= startA && startB <= endA) || (endB >= startA && endB <= endA) ||
			(startA >= startB && startA <= endB) || (endA >= startB && endA <= endB) {
			total++
		}
	}
	return total
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var assings [][4]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		var startA, endA, startB, endB int
		fmt.Sscanf(next, "%d-%d,%d-%d", &startA, &endA, &startB, &endB)
		assings = append(assings, [4]int{startA, endA, startB, endB})
	}

	println("silver:", silver(assings))
	println("gold:", gold(assings))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
