package main

import (
	"bufio"
	"log"
	"os"
)

func silver(lines []string) int {
	for _, line := range lines {
		println(line)
	}
	return 0
}

// func gold(lines []string) int {
// 	for _, line := range lines {
// 		println(line)
// 	}
// 	return 0
// }

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
	// println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
