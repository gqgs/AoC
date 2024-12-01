package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/gqgs/AoC2021/generic"
)

func silver(lines []string) int {
	var total int
	var lleft []int
	var lright []int

	for _, l := range lines {
		var left, right int
		fmt.Sscanf(l, "%d %d", &left, &right)
		lleft = append(lleft, left)
		lright = append(lright, right)
	}

	slices.Sort(lleft)
	slices.Sort(lright)

	for i := range lleft {
		left := lleft[i]
		right := lright[i]
		total += generic.Abs(left - right)
	}

	return total
}

func gold(lines []string) int {
	var total int
	var lleft []int
	rightMap := make(map[int]int)
	for _, l := range lines {
		var left, right int
		fmt.Sscanf(l, "%d %d", &left, &right)
		lleft = append(lleft, left)
		rightMap[right]++
	}

	for i := range lleft {
		total += lleft[i] * rightMap[lleft[i]]
	}

	return total
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

	fmt.Println(silver(input))
	fmt.Println(gold(input))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
