package main

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"

	"github.com/gqgs/AoC2021/generic"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func cost(d int) int {
	return d * (d + 1) / 2
}

func gold(list []int, roudFunc func(x float64) float64) int {
	m := int(roudFunc(mean(list)))
	var sum int
	for _, l := range list {
		sum += cost(abs(m - l))
	}
	return sum
}

func mean(list []int) float64 {
	var sum int
	for _, l := range list {
		sum += l
	}
	return float64(sum) / float64(len(list))
}

func silver(list []int) int {
	m := int(median(list))
	var sum int
	for _, l := range list {
		sum += abs(m - l)
	}
	return sum
}

func median(list []int) float64 {
	mid := len(list) / 2

	if len(list) == 0 {
		return 0
	}
	if len(list)%2 == 0 {
		return float64(generic.QuickSelect(list, mid)+generic.QuickSelect(list, mid-1)) / 2

	}
	return float64(generic.QuickSelect(list, mid))
}

func stringsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i := range strs {
		var err error
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func solve() error {
	file, err := os.Open("day7")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.Read()
	if err != nil {
		return err
	}
	crabs, err := stringsToInts(data)
	if err != nil {
		return err
	}

	println("silver:", silver(crabs))
	println("gold: [", gold(crabs, math.Floor), gold(crabs, math.Ceil), "]")

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
