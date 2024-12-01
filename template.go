package main

import (
	"bufio"
	"os"
)

func silver(lines []string) int {
	for _, line := range lines {
		println(line)
	}
	return 0
}

func gold(lines []string) int {
	for _, line := range lines {
		println(line)
	}
	return 0
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
