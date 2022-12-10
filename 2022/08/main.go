package main

import (
	"bufio"
	"os"

	"github.com/gqgs/AoC2021/generic"
)

func isVisible(input []string, i, j int) int {
	var x, y int
	for x = 0; x < j && input[i][x] < input[i][j]; x++ {
	}
	if x == j {
		return 1
	}

	for y = 0; y < i && input[y][j] < input[i][j]; y++ {
	}
	if y == i {
		return 1
	}

	for x = len(input[0]) - 1; x > j && input[i][x] < input[i][j]; x-- {
	}
	if x == j {
		return 1
	}

	for y = len(input) - 1; y > i && input[y][j] < input[i][j]; y-- {
	}
	if y == i {
		return 1
	}

	return 0
}

func scenicScore(input []string, i, j int) int {
	var x, y int
	score := 1
	for x = j - 1; x > 0 && input[i][x] < input[i][j]; x-- {
	}
	score *= j - x

	for y = i - 1; y > 0 && input[y][j] < input[i][j]; y-- {
	}
	score *= i - y

	for x = j + 1; x < len(input[0])-1 && input[i][x] < input[i][j]; x++ {
	}
	score *= x - j

	for y = i + 1; y < len(input[0])-1 && input[y][j] < input[i][j]; y++ {
	}
	score *= y - i

	return score
}

func silver(input []string) int {
	var total int
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			total += isVisible(input, i, j)
		}
	}
	return total + 4*len(input) - 4
}

func gold(input []string) int {
	var scores []int
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			scores = append(scores, scenicScore(input, i, j))
		}
	}
	return generic.Max(scores...)
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		next := scanner.Text()
		input = append(input, next)
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
