package main

import (
	"bufio"
	"os"
	"sort"
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

	drawMap := make(map[string]int, 0)
	drawSlice := strings.Split(draw, ",")
	for i, n := range drawSlice {
		drawMap[n] = i
	}

	const size = 5

	var boardNumbers [][]string
	var boards [][]int
	for len(rows) > 0 {
		numbers := make([]string, 0, size*size)
		indexes := make([]int, 0, size*size)
		for i := 0; i < size; i++ {
			inputNumbers := strings.Split(rows[i], " ")
			for _, inputNumber := range inputNumbers {
				if len(inputNumber) == 0 {
					continue
				}
				numbers = append(numbers, inputNumber)
				indexes = append(indexes, drawMap[inputNumber])
			}
		}
		boardNumbers = append(boardNumbers, numbers)
		boards = append(boards, indexes)
		rows = rows[size:]
	}

	indexes := make([]BoardIndex, len(boards))
	for index, board := range boards {
		rows := make([]int, 0, size)
		for i := 0; i < size*size; i += size {
			row := board[i]
			for j := i; j < i+size; j++ {
				row = max(row, board[j])
			}
			rows = append(rows, row)
		}

		columns := make([]int, 0, size)
		for i := 0; i < size; i++ {
			column := board[i]
			for j := i; j < size*size; j += size {
				column = max(column, board[j])
			}
			columns = append(columns, column)
		}

		winIndex := min(rows[0], columns[0])
		for i := 0; i < len(rows); i++ {
			for j := 0; j < len(columns); j++ {
				newMin := min(rows[i], columns[j])
				winIndex = min(winIndex, newMin)
			}
		}
		indexes[index] = BoardIndex{
			Index:    index,
			WinIndex: winIndex,
		}
	}

	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i].WinIndex < indexes[j].WinIndex
	})

	println("silver:", bingo(indexes[0], drawSlice, boardNumbers))
	println("gold:", bingo(indexes[len(indexes)-1], drawSlice, boardNumbers))

	return nil
}

func bingo(boardIndex BoardIndex, drawSlice []string, boardNumbers [][]string) int {
	winDrawNumbers := make(map[string]struct{})
	for _, n := range drawSlice[:boardIndex.WinIndex+1] {
		winDrawNumbers[n] = struct{}{}
	}

	var sum int
	for _, n := range boardNumbers[boardIndex.Index] {
		if _, ok := winDrawNumbers[n]; ok {
			continue
		}
		value, _ := strconv.Atoi(n)
		sum += value
	}
	winDraw, _ := strconv.Atoi(drawSlice[boardIndex.WinIndex])
	return winDraw * sum
}

type BoardIndex struct {
	Index    int
	WinIndex int
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
