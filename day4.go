package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solve() error {
	file, err := os.Open("day4")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	draw := scanner.Text()
	rows := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		rows = append(rows, text)
	}

	var boards []Board
	for len(rows) > 0 {
		numbers := make([]string, 0, 25)
		numberIndex := make(map[string]int)
		var index int
		for i := 0; i < 5; i++ {
			inputNumbers := strings.Split(rows[i], " ")
			for _, inputNumber := range inputNumbers {
				if len(inputNumber) == 0 {
					continue
				}
				numbers = append(numbers, inputNumber)
				numberIndex[inputNumber] = index
				index++
			}
		}
		boards = append(boards, Board{
			numbers:     numbers,
			numberIndex: numberIndex,
			numberState: make(map[string]bool),
		})

		rows = rows[5:]
	}

	var first, last int
	skip := make(map[int]bool)
	for _, n := range strings.Split(draw, ",") {
		for i, b := range boards {
			if !skip[i] && b.Bingo(n) {
				number, _ := strconv.Atoi(n)
				last = number * b.UnmarkedSum()
				skip[i] = true
				if first == 0 {
					first = last
				}
			}
		}
	}

	println("silver:", first)
	println("gold:", last)

	return nil
}

type Board struct {
	numbers     []string
	numberIndex map[string]int
	numberState map[string]bool
}

func (b *Board) Bingo(newNumber string) bool {
	index, exists := b.numberIndex[newNumber]
	if !exists {
		return false
	}

	b.numberState[newNumber] = true
	columnBingo := true
	rowBingo := true

	// Column check
	for minColumn := index % 5; minColumn < 25; minColumn += 5 {
		columnBingo = columnBingo && b.numberState[b.numbers[minColumn]]
	}
	// Row check
	for minRow := (index / 5) * 5; minRow < (index/5)*5+5; minRow++ {
		rowBingo = rowBingo && b.numberState[b.numbers[minRow]]
	}

	return columnBingo || rowBingo
}

func (b Board) UnmarkedSum() int {
	var sum int
	for number := range b.numberIndex {
		if !b.numberState[number] {
			n, _ := strconv.Atoi(number)
			sum += n
		}
	}
	return sum
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
