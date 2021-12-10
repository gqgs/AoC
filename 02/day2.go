package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solve() error {
	file, err := os.Open("day2")
	if err != nil {
		return err
	}
	defer file.Close()

	var posX, posY, aim int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		split := strings.Split(next, " ")
		cmd, unitStr := split[0], split[1]

		unit, err := strconv.Atoi(unitStr)
		if err != nil {
			return err
		}

		switch cmd {
		case "forward":
			posX += unit
			posY += aim * unit
		case "down":
			aim += unit
		case "up":
			aim -= unit
		}
	}

	println(posX * posY)

	return nil
}

func main() {
	err := solve()
	if err != nil {
		panic(err)
	}
}
