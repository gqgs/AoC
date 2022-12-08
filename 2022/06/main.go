package main

import (
	"bufio"
	"os"

	"github.com/gqgs/AoC2021/generic"
)

func startMarker(input string, size int) int {
Loop:
	for i := 0; i < len(input)-size; i++ {
		seen := generic.NewSet[byte]()
		for j := i; j < i+size; j++ {
			if seen.Contains(input[j]) {
				continue Loop
			}
			seen.Add(input[j])
		}
		return i + size
	}
	return -1
}

func silver(input string) int {
	return startMarker(input, 4)
}

func gold(input string) int {
	return startMarker(input, 14)
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	println("silver:", silver(input))
	println("gold:", gold(input))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
