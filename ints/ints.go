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
