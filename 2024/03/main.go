package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func silver(lines []string) int {
	var total int
	line := strings.Join(lines, "")
	re := regexp.MustCompile(`mul\([\d+]{1,3},[\d+]{1,3}\)`)
	matches := re.FindAllString(line, -1)
	for _, match := range matches {
		var first, second int
		fmt.Sscanf(match, "mul(%d,%d)", &first, &second)
		total += first * second
	}
	return total
}

func gold(lines []string) int {
	var total int
	line := strings.Join(lines, "")
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllString(line, -1)
	enabled := true
	for _, match := range matches {
		if match == "don't()" {
			enabled = false
			continue
		}
		if match == "do()" {
			enabled = true
			continue
		}
		if !enabled {
			continue
		}
		var first, second int
		fmt.Sscanf(match, "mul(%d,%d)", &first, &second)
		total += first * second
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
		if len(next) == 0 {
			continue
		}
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
