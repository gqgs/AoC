package main

import (
	"bufio"
	"os"
)

func solve() error {
	file, err := os.Open("dayX")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_ = scanner.Text()
	}

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
