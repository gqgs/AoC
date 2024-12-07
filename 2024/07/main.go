package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	Target  int
	Numbers []int
}

func compute(target int, current int, numbers []int, gold bool) bool {
	if target == current && len(numbers) == 0 {
		return true
	}

	if current > target || len(numbers) == 0 {
		return false
	}

	return compute(target, current+numbers[0], numbers[1:], gold) ||
		compute(target, current*numbers[0], numbers[1:], gold) ||
		(gold && compute(target, concat(current, numbers[0]), numbers[1:], gold))
}

func concat(x, y int) int {
	c, _ := strconv.Atoi(fmt.Sprintf("%d%d", x, y))
	return c
}

func shared(lines []string, gold bool) int {
	cases := make([]*Case, 0, len(lines))
	for _, line := range lines {
		var numbers []int
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		for _, number := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}

		cases = append(cases, &Case{
			Target:  target,
			Numbers: numbers,
		})
	}

	var total int
	for _, c := range cases {
		if compute(c.Target, c.Numbers[0], c.Numbers[1:], gold) {
			total += c.Target
		}
	}

	return total
}

func silver(lines []string) int {
	return shared(lines, false)
}

func gold(lines []string) int {
	return shared(lines, true)
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
