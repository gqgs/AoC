package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func enhance(input []string, enhancer, infinityValue string, times int) []string {
	if times == 0 {
		return input
	}
	gridSize := len(input) + 2
	padding := strings.Repeat(infinityValue, gridSize)
	grid := make([]string, 0)
	grid = append(grid, padding)
	for i := range input {
		grid = append(grid, infinityValue+input[i]+infinityValue)
	}
	grid = append(grid, padding)

	var updatedGrid []string
	for y := 0; y < len(grid); y++ {
		var line string
		for x := 0; x < len(grid[y]); x++ {
			var enchanceIndex string
			for dy := y - 1; dy <= y+1; dy++ {
				for dx := x - 1; dx <= x+1; dx++ {
					if dy < 0 || dx < 0 || dy == len(grid) || dx == len(grid[y]) {
						enchanceIndex += infinityValue
						continue
					}
					enchanceIndex += string(grid[dy][dx])
				}
			}
			index, _ := strconv.ParseInt(decode(enchanceIndex), 2, 0)
			line += string(enhancer[index])
		}
		updatedGrid = append(updatedGrid, line)
	}

	infinityValueindex, _ := strconv.ParseInt(decode(strings.Repeat(infinityValue, 9)), 2, 0)
	return enhance(updatedGrid, enhancer, string(enhancer[infinityValueindex]), times-1)
}

func decode(str string) string {
	var decoded string
	for _, s := range str {
		if s == '#' {
			decoded += "1"
		} else {
			decoded += "0"
		}
	}
	return decoded
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	enhancer := scanner.Text()

	input := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		input = append(input, text)

	}

	println("silver:", countLit(enhance(input, enhancer, ".", 2)))
	println("gold:", countLit(enhance(input, enhancer, ".", 50)))

	return nil
}

func printGrid(grid []string) {
	for _, l := range grid {
		fmt.Println(l)
	}
}

func countLit(grid []string) int {
	var lit int
	for _, y := range grid {
		for _, x := range y {
			if x == '#' {
				lit++
			}
		}
	}
	return lit
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
