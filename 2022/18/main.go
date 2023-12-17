package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func touchingSides(cube string) [][3]float64 {
	split := strings.Split(cube, ",")
	x, _ := strconv.ParseFloat(split[0], 64)
	y, _ := strconv.ParseFloat(split[1], 64)
	z, _ := strconv.ParseFloat(split[2], 64)

	return [][3]float64{
		{x, y, z - 0.5},
		{x, y, z + 0.5},
		{x - 0.5, y, z},
		{x + 0.5, y, z},
		{x, y - 0.5, z},
		{x, y + 0.5, z},
	}
}

func silver(cubes []string) int {
	var total int
	surface := make(map[string]struct{})
	for _, c := range cubes {
		sides := touchingSides(c)
		var conflict int
		for _, s := range sides {
			key := fmt.Sprintf("%0.2f/%0.2f/%0.2f", s[0], s[1], s[2])
			if _, ok := surface[key]; ok {
				conflict++
			}
			if conflict == 6 {
				println("key", key)
			}
			surface[key] = struct{}{}
		}
		total = total + 6
		total = total - (2 * conflict)
	}
	return total
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cubes []string
	for scanner.Scan() {
		next := scanner.Text()
		cubes = append(cubes, next)
	}

	println("silver:", silver(cubes))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
