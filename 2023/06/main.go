package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func records(time, distance int) int {
	total := 1
	for i := 1; i < time; i++ {
		speed := i
		if speed*(time-i) > distance {
			total++
		}
	}
	return total
}

func silver(lines []string) int {
	var times []int
	var distances []int
	for _, char := range strings.Split(lines[0], " ")[1:] {
		if len(char) == 0 {
			continue
		}
		t, _ := strconv.Atoi(char)
		times = append(times, t)
	}

	for _, char := range strings.Split(lines[1], " ")[1:] {
		if len(char) == 0 {
			continue
		}
		d, _ := strconv.Atoi(char)
		distances = append(distances, d)
	}

	total := 1
	for i := range times {
		total *= (records(times[i], distances[i]) - 1)
	}

	return total
}

func gold(lines []string) int {
	time, _ := strconv.Atoi(strings.Join(strings.Split(lines[0], " ")[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Split(lines[1], " ")[1:], ""))
	return records(time, distance) - 1
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

	// println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
