package main

import (
	"bufio"
	"os"
)

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		println(next)
	}

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
