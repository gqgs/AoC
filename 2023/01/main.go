package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stringReplacer = strings.NewReplacer(
	"oneight", "18",
	"twone", "21",
	"threeight", "38",
	"fiveight", "58",
	"sevenine", "79",
	"eightwo", "82",
	"eighthree", "83",
	"nineight", "98",
	"one", "1",
	"two", "2",
	"three", "3",
	"four", "4",
	"five", "5",
	"six", "6",
	"seven", "7",
	"eight", "8",
	"nine", "9",
)

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func findFirstAndLast(list []string) []int {
	var result []int
	for _, input := range list {
		var first, last int
		for first = 0; first < len(input); first++ {
			if isDigit(input[first]) {
				break
			}
		}
		for last = len(input) - 1; last >= 0; last-- {
			if isDigit(input[last]) {
				break
			}
		}
		// fmt.Println(input, first, last, string(input[first])+string(input[last]))
		number, err := strconv.Atoi(string(input[first]) + string(input[last]))
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	return result
}

func sum(list []int) int {
	var result int
	for _, e := range list {
		result += e
	}
	return result
}

func transformInput(list []string) []string {
	var newList []string
	for _, input := range list {
		newList = append(newList, stringReplacer.Replace(input))
	}
	return newList
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

	// fmt.Println("silver:", sum(findFirstAndLast(input)))
	fmt.Println("gold:", sum(findFirstAndLast(transformInput(input))))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
