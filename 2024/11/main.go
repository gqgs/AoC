package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gqgs/AoC2021/ints"
)

func eval(number int, iter int, cache map[string]int) int {
	key := fmt.Sprintf("%d,%d", number, iter)
	if cached, ok := cache[key]; ok {
		return cached
	}
	var result int
	switch {
	case iter == 0:
		result = 1
	case number == 0:
		result = eval(1, iter-1, cache)
	case len(strconv.Itoa(number))%2 == 0:
		mid := len(strconv.Itoa(number)) / 2
		n1, _ := strconv.Atoi(strconv.Itoa(number)[:mid])
		n2, _ := strconv.Atoi(strconv.Itoa(number)[mid:])
		result = eval(n1, iter-1, cache) + eval(n2, iter-1, cache)
	default:
		result = eval(number*2024, iter-1, cache)
	}
	cache[key] = result
	return result
}

func shared(line string, iter int) int {
	var total int
	cache := make(map[string]int)
	for _, number := range ints.FromString(line, ",") {
		total += eval(number, iter, cache)
	}
	return total
}

func silver(line string) int {
	return shared(line, 25)
}

func gold(line string) int {
	return shared(line, 75)
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

	if err := scanner.Err(); err != nil {
		return err
	}

	println(silver(lines[0]))
	println(gold(lines[0]))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
