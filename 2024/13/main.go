package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	AX, AY         int
	BX, BY         int
	PrizeX, PrizeY int
}

func shared(games []Game) int {
	var total int
	for _, g := range games {
		y := (g.PrizeX*g.AY - g.PrizeY*g.AX) / (g.BX*g.AY - g.BY*g.AX)
		x := (g.PrizeX*g.BY - g.PrizeY*g.BX) / (g.AX*g.BY - g.AY*g.BX)

		valid := (x*g.AX + y*g.BX) == g.PrizeX
		valid = valid && (x*g.AY+y*g.BY) == g.PrizeY
		if valid {
			total += 3*x + y
		}
	}

	return total
}

func silver(lines []string) int {
	var games []Game
	for i := 0; i < len(lines); i += 3 {
		var game Game
		var button rune
		fmt.Sscanf(lines[i], "Button %c: X+%d, Y+%d", &button, &game.AX, &game.AY)
		fmt.Sscanf(lines[i+1], "Button %c: X+%d, Y+%d", &button, &game.BX, &game.BY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &game.PrizeX, &game.PrizeY)
		games = append(games, game)
	}

	return shared(games)
}

func gold(lines []string) int {
	var games []Game
	for i := 0; i < len(lines); i += 3 {
		var game Game
		var button rune
		fmt.Sscanf(lines[i], "Button %c: X+%d, Y+%d", &button, &game.AX, &game.AY)
		fmt.Sscanf(lines[i+1], "Button %c: X+%d, Y+%d", &button, &game.BX, &game.BY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &game.PrizeX, &game.PrizeY)

		game.PrizeX += 10000000000000
		game.PrizeY += 10000000000000

		games = append(games, game)
	}

	return shared(games)
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
		next = strings.TrimSpace(next)
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
		panic(err)
	}
}
