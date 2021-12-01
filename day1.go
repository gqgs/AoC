package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func compare(values []uint64) uint {
	var increased uint
	var prev uint64 = math.MaxUint

	for _, next := range values {
		if prev < next {
			increased++
		}
		prev = next
	}
	return increased
}

func solve() error {
	file, err := os.Open("day1")
	if err != nil {
		return err
	}
	defer file.Close()

	var values []uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, err := strconv.ParseUint(scanner.Text(), 0, 64)
		if err != nil {
			return err
		}
		values = append(values, next)
	}

	println("silver:", compare(values))

	var windowed []uint64
	for i := 0; i < len(values); i++ {
		if i+2 < len(values) {
			windowed = append(windowed, values[i]+values[i+1]+values[i+2])
		}
	}

	println("gold:", compare(windowed))

	return nil
}

func main() {
	err := solve()
	if err != nil {
		panic(err)
	}
}
