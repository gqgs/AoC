package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

func silver(lines []string) int {
	var total int
	for _, line := range lines {
		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning := parts[0]
		numbers := parts[1]
		winSet := make(generic.Set[string])
		numbersSet := make(generic.Set[string])

		for _, s := range strings.Split(winning, " ") {
			if len(s) == 0 {
				continue
			}
			winSet.Add(s)
		}
		for _, s := range strings.Split(numbers, " ") {
			if len(s) == 0 {
				continue
			}
			numbersSet.Add(s)
		}

		intersections := winSet.Intersect(numbersSet)
		if len(intersections) == 0 {
			continue
		}
		i := 1
		for j := 1; j < len(intersections); j++ {
			i *= 2
		}
		total += i
	}
	return total
}

func gold(lines []string) int {
	var total int
	totalSlice := make([]int, 1000)
	for iline, line := range lines {
		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning := parts[0]
		numbers := parts[1]
		winSet := make(generic.Set[string])
		numbersSet := make(generic.Set[string])

		for _, s := range strings.Split(winning, " ") {
			if len(s) == 0 {
				continue
			}
			winSet.Add(s)
		}
		for _, s := range strings.Split(numbers, " ") {
			if len(s) == 0 {
				continue
			}
			numbersSet.Add(s)
		}

		intersections := winSet.Intersect(numbersSet)
		for j := iline + 1; j < iline+1+len(intersections); j++ {
			totalSlice[j] += 1 + totalSlice[iline]
		}

		total += 1 + totalSlice[iline]
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

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
