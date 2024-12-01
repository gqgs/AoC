package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func silver(lines []string) int {
	var total int

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

Next:
	for _, l := range lines {
		var gameId int
		var game string
		fmt.Sscanf(l, "Game %d:", &gameId)
		game = strings.TrimPrefix(l, fmt.Sprintf("Game %d: ", gameId))
		sets := strings.Split(game, "; ")
		for _, set := range sets {
			for _, cubes := range strings.Split(set, ", ") {
				var count int
				var color string
				fmt.Sscanf(cubes, "%d %s", &count, &color)
				if count > limits[color] {
					continue Next
				}
			}
		}
		total += gameId
	}
	return total
}

func gold(lines []string) int {
	var total int
	for _, l := range lines {
		var gameId int
		var game string
		fmt.Sscanf(l, "Game %d:", &gameId)
		game = strings.TrimPrefix(l, fmt.Sprintf("Game %d: ", gameId))
		sets := strings.Split(game, "; ")
		power := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		for _, set := range sets {
			for _, cubes := range strings.Split(set, ", ") {
				var count int
				var color string
				fmt.Sscanf(cubes, "%d %s", &count, &color)
				power[color] = max(count, power[color])
			}
		}
		total += (power["red"] * power["blue"] * power["green"])
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
