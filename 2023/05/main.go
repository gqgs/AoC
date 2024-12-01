package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

type Map struct {
	Destination int
	Source      int
	Range       int
}

func (m Map) Left() int {
	return m.Source
}

func (m Map) Right() int {
	return m.Source + m.Range
}

func (m Map) transform(n int) (int, bool) {
	if n >= m.Source && n < (m.Source+m.Range) {
		return m.Destination + (n - m.Source), true
	}
	return n, false
}

func seedValue(seed int, maps []Map, mapSizes []int) int {
	var i int
	for _, r := range mapSizes {
		for _, m := range maps[i : i+r] {
			t, inRange := m.transform(seed)
			if inRange {
				seed = t
				break
			}
		}
		i += r
	}
	return seed
}

func silver(seeds []int, maps []Map, mapSizes []int) []int {
	for i, seed := range seeds {
		seeds[i] = seedValue(seed, maps, mapSizes)
	}
	return seeds
}

func transformRange(initMap Map, transformMaps []Map) (resultMaps []Map) {
	remaining := make(generic.Queue[Map], 0)
	remaining.Push(initMap)
Next:
	for len(remaining) > 0 {
		rangeMap := remaining.Pop()
		for _, transformMap := range transformMaps {
			// 0
			if transformMap.Left() <= rangeMap.Left() && transformMap.Right() >= rangeMap.Right() {
				resultMaps = append(resultMaps, Map{
					Source: transformMap.Destination - transformMap.Source + rangeMap.Left(),
					Range:  rangeMap.Range,
				})
				continue Next
			}

			// 1
			if transformMap.Left() < rangeMap.Left() && transformMap.Right() > rangeMap.Left() && transformMap.Right() < rangeMap.Right() {
				resultMaps = append(resultMaps, Map{
					Source: transformMap.Destination - transformMap.Source + rangeMap.Left(),
					Range:  transformMap.Right() - rangeMap.Left(),
				})
				remaining.Push(Map{
					Source: transformMap.Right(),
					Range:  rangeMap.Right() - transformMap.Right(),
				})
				continue Next
			}

			// 2
			if transformMap.Right() > rangeMap.Right() && transformMap.Left() < rangeMap.Right() && transformMap.Left() > rangeMap.Left() {
				resultMaps = append(resultMaps, Map{
					Source: transformMap.Destination - transformMap.Source + transformMap.Left(),
					Range:  rangeMap.Right() - transformMap.Left(),
				})
				remaining.Push(Map{
					Source: rangeMap.Left(),
					Range:  transformMap.Left() - rangeMap.Left(),
				})
				continue Next
			}

			// 5
			if transformMap.Left() > rangeMap.Left() && transformMap.Right() < rangeMap.Right() {
				resultMaps = append(resultMaps, Map{
					Source: transformMap.Destination - transformMap.Source + transformMap.Left(),
					Range:  transformMap.Right() - transformMap.Left(),
				})
				remaining.Push(Map{
					Source: rangeMap.Left(),
					Range:  transformMap.Left() - rangeMap.Left(),
				})
				remaining.Push(Map{
					Source: transformMap.Right(),
					Range:  rangeMap.Right() - transformMap.Right(),
				})
				continue Next
			}
		}
		// 3 && 4
		resultMaps = append(resultMaps, rangeMap)
	}

	return
}

func gold(seeds []int, maps []Map, mapSizes []int) []int {
	rangeMaps := rangeMap(seeds)
	var i int
	for _, ms := range mapSizes {
		var newRangeMaps []Map
		for _, rm := range rangeMaps {
			newRangeMaps = append(newRangeMaps, transformRange(rm, maps[i:i+ms])...)
		}
		rangeMaps = newRangeMaps
		i += ms
	}

	results := make([]int, len(rangeMaps))
	for j, rm := range rangeMaps {
		results[j] = rm.Source
	}

	return results
}

func rangeMap(seeds []int) []Map {
	var rangeMap []Map
	for i := 0; i < len(seeds); i += 2 {
		rangeMap = append(rangeMap, Map{
			Source: seeds[i],
			Range:  seeds[i+1],
		})
	}
	return rangeMap
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
