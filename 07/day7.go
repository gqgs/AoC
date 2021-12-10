package main

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"strconv"
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
		return float64(quickSelect(list, mid)+quickSelect(list, mid-1)) / 2

	}
	return float64(quickSelect(list, mid))
}

func quickSelect(list []int, n int) int {
	return selectKth(list, 0, len(list)-1, n)
}

func selectKth(list []int, left, right, k int) int {
	for {
		if left == right {
			return list[left]
		}
		pivotIndex := partition(list, left, right, rand.Intn(right-left+1)+left)
		if k == pivotIndex {
			return list[k]
		}
		if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func partition(list []int, left, right, pivotIndex int) int {
	pivot := list[pivotIndex]
	list[pivotIndex], list[right] = list[right], list[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if list[i] < pivot {
			list[storeIndex], list[i] = list[i], list[storeIndex]
			storeIndex++
		}
	}
	list[storeIndex], list[right] = list[right], list[storeIndex]
	return storeIndex
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
