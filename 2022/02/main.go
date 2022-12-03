package main

import (
	"bufio"
	"os"
)

// A: Rock
// B: Paper
// C: Scissors

// X: Rock
// Y: Paper
// Z: Scissors

// (1 for Rock, 2 for Paper, and 3 for Scissors)
// (0 if you lost, 3 if the round was a draw, and 6 if you won)

func format(p byte) byte {
	switch p {
	case 'A', 'X':
		return 'R'
	case 'B', 'Y':
		return 'P'
	case 'C', 'Z':
		return 'S'
	}
	panic("no")
}

func silver(p1, p2 byte) int {
	switch p1 {
	case 'R':
		switch p2 {
		case 'R':
			return 1 + 3
		case 'P':
			return 1 + 0
		case 'S':
			return 1 + 6
		}
	case 'P':
		switch p2 {
		case 'R':
			return 2 + 6
		case 'P':
			return 2 + 3
		case 'S':
			return 2 + 0
		}
	case 'S':
		switch p2 {
		case 'R':
			return 3 + 0
		case 'P':
			return 3 + 6
		case 'S':
			return 3 + 3
		}
	}
	panic("no")
}

// R means you need to lose
// P means you need to end the round in a draw
// S means you need to win

// (1 for Rock, 2 for Paper, and 3 for Scissors)
// (0 if you lost, 3 if the round was a draw, and 6 if you won)

func gold(p1, p2 byte) int {
	switch p1 {
	case 'R':
		switch p2 {
		case 'R':
			return 0 + 3
		case 'P':
			return 0 + 1
		case 'S':
			return 0 + 2
		}
	case 'P':
		switch p2 {
		case 'R':
			return 3 + 1
		case 'P':
			return 3 + 2
		case 'S':
			return 3 + 3
		}
	case 'S':
		switch p2 {
		case 'R':
			return 6 + 2
		case 'P':
			return 6 + 3
		case 'S':
			return 6 + 1
		}
	}
	panic("no")
}

// X means you need to lose
// Y means you need to end the round in a draw
// Z means you need to win

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var silverTotal int
	var goldTotal int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		p1 := format(next[0])
		p2 := format(next[2])
		silverTotal += silver(p2, p1)
		goldTotal += gold(p2, p1)
	}
	println("silver:", silverTotal)
	println("gold:", goldTotal)

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
