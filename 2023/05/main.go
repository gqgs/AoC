package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Map struct {
	Destination int
	Source      int
	Range       int
}

func (m Map) transform(n int) (int, bool) {
	if n >= m.Source && n < (m.Source+m.Range) {
		return m.Destination + (n - m.Source), true
	}
	return n, false
}

func silver(seeds []int, maps []Map, mapSizes []int) []int {
	var i int
	for _, r := range mapSizes {
		var newSeeds []int
	Next:
		for _, s := range seeds {
			for _, m := range maps[i : i+r] {
				t, inRange := m.transform(s)
				if inRange {
					newSeeds = append(newSeeds, t)
					continue Next
				}
			}
			newSeeds = append(newSeeds, s)
		}
		i += r
		seeds = newSeeds
	}
	return seeds
}

func gold(seeds []int, maps []Map, mapSizes []int) []int {
	return silver(goldSeeds(seeds), maps, mapSizes)
}

func goldSeeds(seeds []int) []int {
	var newSeeds []int
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			newSeeds = append(newSeeds, j)
		}
	}
	return newSeeds
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var seeds []int
	var maps []Map
	var mapSize int
	var mapSizes []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			continue
		}
		if strings.HasPrefix(next, "seeds: ") {
			seeds = parseSeeds(strings.TrimPrefix(next, "seeds: "))
			continue
		}
		if strings.Contains(next, "map") {
			if mapSize > 0 {
				mapSizes = append(mapSizes, mapSize)
				mapSize = 0
			}
			continue
		}

		maps = append(maps, parseParseMap(next))
		mapSize++
	}
	mapSizes = append(mapSizes, mapSize)

	fmt.Println("silver:", slices.Min(silver(seeds, maps, mapSizes)))
	fmt.Println("gold:", slices.Min(gold(seeds, maps, mapSizes)))

	return nil
}

func parseParseMap(line string) Map {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		panic(fmt.Errorf("unexpected parts length: %d", len(parts)))
	}

	dest, _ := strconv.Atoi(parts[0])
	src, _ := strconv.Atoi(parts[1])
	rnge, _ := strconv.Atoi(parts[2])

	return Map{
		Destination: dest,
		Source:      src,
		Range:       rnge,
	}
}

func parseSeeds(input string) []int {
	var seeds []int
	for _, s := range strings.Split(input, " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, n)
	}
	return seeds
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
