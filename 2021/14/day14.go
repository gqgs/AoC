package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func minMaxDiff(initialState string, transitionGraph map[string]string, n int) int {
	counter := make(map[string]int)
	for i := 0; i < len(initialState)-1; i++ {
		counter[initialState[i:i+2]]++
	}

	updateValues := func(key string) [2]string {
		return [2]string{string(key[0]) + transitionGraph[key], transitionGraph[key] + string(key[1])}
	}

	for i := 0; i < n; i++ {
		updatedCounter := make(map[string]int)
		for cKey, cValue := range counter {
			for _, v := range updateValues(cKey) {
				updatedCounter[v] += cValue
			}
		}
		counter = updatedCounter
	}

	runeCount := make(map[rune]int)
	for key, count := range counter {
		for _, r := range key {
			runeCount[r] += count
		}
	}

	maxResult := 0
	minResult := 1<<63 - 1
	for _, c := range runeCount {
		res := int(math.Ceil(float64(c) / 2))
		maxResult = max(maxResult, res)
		minResult = min(minResult, res)
	}

	return maxResult - minResult
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	initialState := scanner.Text()

	transitionGraph := make(map[string]string)
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			continue
		}
		var from, to string
		fmt.Sscanf(next, "%s -> %s", &from, &to)
		transitionGraph[from] = to
	}

	println("silver:", minMaxDiff(initialState, transitionGraph, 10))
	println("gold:", minMaxDiff(initialState, transitionGraph, 40))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
