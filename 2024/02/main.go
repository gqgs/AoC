package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

func isSafe(numbers []int) bool {
	if !sort.IntsAreSorted(numbers) {
		return false
	}

	reversed := slices.Clone(numbers)
	slices.Reverse(reversed)
	if !sort.IntsAreSorted(reversed) {
		return false
	}

	for i := 1; i < len(numbers); i++ {
		diff := generic.Abs(numbers[i-1] - numbers[i])
		if diff == 0 || diff > 3 {
			return false
		}
	}

	return true
}

func gold(lines []string) int {
	var total int
Next:
	for _, line := range lines {
		var numbers []int
		for _, digit := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(digit)
			numbers = append(numbers, number)
		}

		if isSafe(numbers) {
			total++
			continue
		}

		for i := 0; i < len(numbers); i++ {
			newNumbers := slices.Clone(numbers)
			updated := slices.Delete(newNumbers, i, i+1)
			if isSafe(updated) {
				total++
				continue Next
			}
		}
	}
	return total
}

func silver(lines []string) int {
	var total int
	for _, line := range lines {
		var numbers []int
		for _, digit := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(digit)
			numbers = append(numbers, number)
		}

		if isSafe(numbers) {
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

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			continue
		}
		input = append(input, next)
	}

	fmt.Println(silver(input))
	fmt.Println(gold(input))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
