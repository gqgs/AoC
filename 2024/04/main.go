package main

import (
	"bufio"
	"os"

	"github.com/gqgs/AoC2021/grid"
)

func checkSilverCases(lines []string, i, j int) int {
	var total int
	if lines[i][j] == 'X' &&
		lines[i][j+1] == 'M' &&
		lines[i][j+2] == 'A' &&
		lines[i][j+3] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i][j-1] == 'M' &&
		lines[i][j-2] == 'A' &&
		lines[i][j-3] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i+1][j] == 'M' &&
		lines[i+2][j] == 'A' &&
		lines[i+3][j] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i-1][j] == 'M' &&
		lines[i-2][j] == 'A' &&
		lines[i-3][j] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i+1][j+1] == 'M' &&
		lines[i+2][j+2] == 'A' &&
		lines[i+3][j+3] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i-1][j-1] == 'M' &&
		lines[i-2][j-2] == 'A' &&
		lines[i-3][j-3] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i-1][j+1] == 'M' &&
		lines[i-2][j+2] == 'A' &&
		lines[i-3][j+3] == 'S' {
		total += 1
	}

	if lines[i][j] == 'X' &&
		lines[i+1][j-1] == 'M' &&
		lines[i+2][j-2] == 'A' &&
		lines[i+3][j-3] == 'S' {
		total += 1
	}

	return total
}

func silver(lines []string) int {
	var total int
	for i := range lines {
		for j := range lines[i] {
			total += checkSilverCases(lines, i, j)
		}
	}
	return total
}

func checkGoldCases(lines []string, i, j int) int {
	var total int
	if lines[i][j] == 'A' &&
		lines[i-1][j-1] == 'M' &&
		lines[i-1][j+1] == 'M' &&
		lines[i+1][j-1] == 'S' &&
		lines[i+1][j+1] == 'S' {
		total += 1
	}

	if lines[i][j] == 'A' &&
		lines[i-1][j-1] == 'S' &&
		lines[i-1][j+1] == 'M' &&
		lines[i+1][j-1] == 'S' &&
		lines[i+1][j+1] == 'M' {
		total += 1
	}

	if lines[i][j] == 'A' &&
		lines[i-1][j-1] == 'S' &&
		lines[i-1][j+1] == 'S' &&
		lines[i+1][j-1] == 'M' &&
		lines[i+1][j+1] == 'M' {
		total += 1
	}

	if lines[i][j] == 'A' &&
		lines[i-1][j-1] == 'M' &&
		lines[i-1][j+1] == 'S' &&
		lines[i+1][j-1] == 'M' &&
		lines[i+1][j+1] == 'S' {
		total += 1
	}

	return total
}

func gold(lines []string) int {
	var total int
	for i := range lines {
		for j := range lines[i] {
			total += checkGoldCases(lines, i, j)
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

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			continue
		}
		lines = append(lines, next)
	}

	lines = grid.Fill(lines, ".", 4)

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
