package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countFishes(state map[int]int, days int) int {
	for ; days > 0; days-- {
		updatedState := make(map[int]int)
		for number, fishes := range state {
			if number == 0 {
				updatedState[8] += fishes
				updatedState[6] += fishes
				continue
			}
			updatedState[number-1] += fishes
		}
		state = updatedState
	}

	var sum int
	for _, v := range state {
		sum += v
	}
	return sum
}

func solve() error {
	file, err := os.Open("day6")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	initialState := strings.Split(scanner.Text(), ",")

	state := make(map[int]int)
	for _, s := range initialState {
		n, _ := strconv.Atoi(s)
		state[n]++
	}

	println("silver:", countFishes(state, 80))
	println("gold:", countFishes(state, 256))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
