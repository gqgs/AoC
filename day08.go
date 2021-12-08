package main

import (
	"bufio"
	"math/bits"
	"os"
	"sort"
	"strings"
)

func solve() error {
	file, err := os.Open("day08")
	if err != nil {
		return err
	}
	defer file.Close()

	runeMap := map[rune]uint{
		'a': 1,
		'b': 1 << 1,
		'c': 1 << 2,
		'd': 1 << 3,
		'e': 1 << 4,
		'f': 1 << 5,
		'g': 1 << 6,
	}

	sortIndexes := [][]int{
		{3, 6, 0},
		{6, 9, 3},
		{4, 6, 6},
		{7, 9, 0},
	}

	numbersIndex := [...]int{1, 7, 4, 3, 5, 2, 9, 0, 6, 8}

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "|")
		input := strings.Split(strings.TrimSpace(split[0]), " ")
		output := strings.Split(strings.TrimSpace(split[1]), " ")

		inputNumbers := make([]uint, 10)
		for i := range input {
			var inputNumber uint
			for _, c := range input[i] {
				inputNumber |= runeMap[c]
			}
			inputNumbers[i] = inputNumber
		}

		sort.Slice(inputNumbers, func(i, j int) bool {
			return bits.OnesCount(inputNumbers[i]) < bits.OnesCount(inputNumbers[j])
		})

		for _, indexes := range sortIndexes {
			startIndex, endIndex, maskIndex := indexes[0], indexes[1], indexes[2]
			sort.Slice(inputNumbers[startIndex:endIndex], func(i, j int) bool {
				mask := inputNumbers[maskIndex]
				return bits.OnesCount(inputNumbers[i+startIndex]&mask) > bits.OnesCount(inputNumbers[j+startIndex]&mask)
			})
		}

		var base int = 1e3
		for _, o := range output {
			var outputNumber uint
			for _, c := range o {
				outputNumber |= runeMap[c]
			}
			for i := range inputNumbers {
				if inputNumbers[i]^outputNumber == 0 {
					sum += base * numbersIndex[i]
					base /= 10
					break
				}
			}
		}
	}
	println(sum)

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
