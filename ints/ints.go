package ints

import (
	"strconv"
	"strings"
)

func FromString(line string, splitChar string) []int {
	parts := strings.Split(line, splitChar)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		number, _ := strconv.Atoi(part)
		numbers[i] = number
	}
	return numbers
}

func FromList(lines []string) []int {
	results := make([]int, 0, len(lines))
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		results = append(results, n)
	}
	return results
}

func Sum(list []int) int {
	var result int
	for _, l := range list {
		result += l
	}
	return result
}
