package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func hasSolution(input string, programs map[string]struct{}) bool {
	if len(input) == 0 {
		return true
	}

	for prefix := range programs {
		if strings.HasPrefix(input, prefix) {
			if hasSolution(strings.TrimPrefix(input, prefix), programs) {
				return true
			}
		}
	}

	return false
}

func allPossibleSolutions(input string, programs map[string]struct{}, cache map[string]int) int {
	if len(input) == 0 {
		return 1
	}

	value, cached := cache[input]
	if cached {
		return value
	}

	var total int
	for prefix := range programs {
		if strings.HasPrefix(input, prefix) {
			total += allPossibleSolutions(strings.TrimPrefix(input, prefix), programs, cache)
		}
	}

	cache[input] = total

	return total
}

func silver(lines []string) int {
	programs := make(map[string]struct{})
	for _, r := range strings.Split(lines[0], ", ") {
		programs[r] = struct{}{}
	}

	var total int
	for _, input := range lines[1:] {
		if hasSolution(input, programs) {
			total++
		}
	}
	return total
}

func gold(lines []string) int {
	programs := make(map[string]struct{})
	for _, r := range strings.Split(lines[0], ", ") {
		programs[r] = struct{}{}
	}

	cache := make(map[string]int)

	var total int
	for _, input := range lines[1:] {
		total += allPossibleSolutions(input, programs, cache)
	}
	return total
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
		log.Fatal(err)
	}
}
