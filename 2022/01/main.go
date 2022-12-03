package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/gqgs/AoC2021/generic"
)

func gold(elfCalories map[int]int) int {
	max := generic.NewMinHeap(func(e1, e2 int) bool {
		return e1 > e2
	})
	max.Push(0, 0, 0)

	for _, calorie := range elfCalories {
		max.Push(calorie)
	}
	return max.Pop() + max.Pop() + max.Pop()
}

func silver(elfCalories map[int]int) int {
	var max int
	for _, calorie := range elfCalories {
		if calorie > max {
			max = calorie
		}
	}
	return max
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	elf := 1
	elfCalories := make(map[int]int)
	totalCalorie := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			elf++
			totalCalorie = 0
			continue
		}
		calorie, _ := strconv.Atoi(next)
		totalCalorie += calorie
		elfCalories[elf] = totalCalorie
	}

	println("silver:", silver(elfCalories))
	println("gold:", gold(elfCalories))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
