package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func silver(lines []string) int {
	var total int
	index := slices.Index(lines, "")
	rules := lines[:index]
	updates := lines[index+1:]

	var digitRules [][2]int
	for _, rule := range rules {
		var left, right int
		fmt.Sscanf(rule, "%d|%d", &left, &right)
		digitRules = append(digitRules, [2]int{left, right})
	}

	var valid []string
Next:
	for _, line := range updates {
		for _, digitRule := range digitRules {
			leftIndex := strings.Index(line, strconv.Itoa(digitRule[0]))
			rightIndex := strings.Index(line, strconv.Itoa(digitRule[1]))
			if leftIndex == -1 || rightIndex == -1 {
				continue
			}

			if leftIndex > rightIndex {
				continue Next
			}
		}
		valid = append(valid, line)
	}

	for _, line := range valid {
		digits := strings.Split(line, ",")
		middleString := digits[len(digits)/2]
		middle, _ := strconv.Atoi(middleString)
		total += middle
	}

	return total
}

func gold(lines []string) int {
	var total int
	index := slices.Index(lines, "")
	rules := lines[:index]
	updates := lines[index+1:]

	var digitRules [][2]int
	for _, rule := range rules {
		var left, right int
		fmt.Sscanf(rule, "%d|%d", &left, &right)
		digitRules = append(digitRules, [2]int{left, right})
	}

	var splitList [][]string
	for _, line := range updates {
		split := strings.Split(line, ",")
		splitList = append(splitList, split)
	}

	var invalid [][]string
Next:
	for _, split := range splitList {
		for _, digitRule := range digitRules {
			leftIndex := slices.Index(split, strconv.Itoa(digitRule[0]))
			rightIndex := slices.Index(split, strconv.Itoa(digitRule[1]))
			if leftIndex == -1 || rightIndex == -1 {
				continue
			}

			if leftIndex > rightIndex {
				invalid = append(invalid, split)
				continue Next
			}
		}
	}

	for _, line := range invalid {
		for _, digit := range line {
			var left, right int
			for _, digitRule := range digitRules {
				leftIndex := slices.Index(line, strconv.Itoa(digitRule[0]))
				rightIndex := slices.Index(line, strconv.Itoa(digitRule[1]))
				if leftIndex == -1 || rightIndex == -1 {
					continue
				}
				if digit == fmt.Sprint(digitRule[0]) {
					left++
				}
				if digit == fmt.Sprint(digitRule[1]) {
					right++
				}
			}

			if left == right {
				middle, _ := strconv.Atoi(digit)
				total += middle
			}
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
		lines = append(lines, next)
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
