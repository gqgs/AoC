package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process(secret int) int {
	result := 64 * secret
	secret = secret ^ result   // mix
	secret = secret % 16777216 // prune

	result2 := secret / 32
	secret = secret ^ result2  // mix
	secret = secret % 16777216 // prune

	result3 := secret * 2048
	secret = secret ^ result3  // mix
	secret = secret % 16777216 // prune

	return secret
}

func diff(list []int) []int {
	var updated []int
	for i := range len(list) - 1 {
		updated = append(updated, list[i+1]-list[i])
	}
	return updated
}

func encode(list []int) string {
	builder := new(strings.Builder)
	for _, l := range list {
		switch l {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			builder.WriteString(fmt.Sprint(l))
		case -1:
			builder.WriteRune('A')
		case -2:
			builder.WriteRune('B')
		case -3:
			builder.WriteRune('C')
		case -4:
			builder.WriteRune('D')
		case -5:
			builder.WriteRune('E')
		case -6:
			builder.WriteRune('F')
		case -7:
			builder.WriteRune('G')
		case -8:
			builder.WriteRune('H')
		case -9:
			builder.WriteRune('I')
		}
	}

	return builder.String()
}

func silver(lines []string) int {
	var list []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		list = append(list, n)
	}

	for range 2000 {
		for i := range list {
			list[i] = process(list[i])
		}
	}

	var result int
	for _, d := range list {
		result += d
	}

	return result
}

func gold(lines []string) int {
	var list []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		list = append(list, n)
	}

	prices := make([][]int, len(list))
	for range 2000 {
		for i := range list {
			prices[i] = append(prices[i], list[i]%10)
			list[i] = process(list[i])
		}
	}

	var encodes []string
	for _, p := range prices {
		encodes = append(encodes, encode(diff(p)))
	}

	var encodingChangeValues []map[string]int
	existingChanges := make(map[string]struct{})
	for i, e := range encodes {
		encodeChanges := make(map[string]int)
		for index := len(e) - 4; index >= 0; index-- {
			substr := e[index : index+4]
			existingChanges[substr] = struct{}{}
			encodeChanges[substr] = prices[i][index+4]
		}
		encodingChangeValues = append(encodingChangeValues, encodeChanges)
	}

	values := make(map[string]int)
	for i := range encodes {
		for change := range existingChanges {
			values[change] += encodingChangeValues[i][change]
		}
	}

	var maxValue int
	for _, v := range values {
		maxValue = max(maxValue, v)
	}

	return maxValue
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

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
