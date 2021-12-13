package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	graph := make(map[string]map[string]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		node1, node2 := line[0], line[1]

		for _, node := range []string{node1, node2} {
			if graph[node] == nil {
				graph[node] = make(map[string]struct{})
			}
		}

		graph[node1][node2] = struct{}{}
		graph[node2][node1] = struct{}{}
	}

	cache := make(map[string]int)

	println("silver:", silver("start", []string{"start"}, graph))
	println("gold:", gold("start", []string{"start"}, false, graph, cache))

	return nil
}

func visitedTimes(list []string, n string) int {
	var times int
	for _, e := range list {
		if e == n {
			times++
		}
	}
	return times
}

func silver(node string, path []string, graph map[string]map[string]struct{}) int {
	if node == "end" {
		return 1
	}

	var paths int
	for adjacent := range graph[node] {
		if adjacent == "start" {
			continue
		}
		if unicode.IsLower(rune(adjacent[0])) {
			vTimes := visitedTimes(path, adjacent)
			switch vTimes {
			case 0:
			default:
				continue
			}
		}
		paths += silver(adjacent, append(path, adjacent), graph)
	}
	return paths
}

func key(node string, list []string) string {
	filteres := make([]string, 0, len(list))
	for _, e := range list {
		if unicode.IsUpper(rune(e[0])) {
			continue
		}
		filteres = append(filteres, e)
	}
	filteres = append(filteres, node)
	return strings.Join(filteres, ",")
}

func gold(node string, path []string, alreadyDoubleVisited bool, graph map[string]map[string]struct{}, cache map[string]int) int {
	cacheKey := key(node, path)
	if _, cacheHit := cache[cacheKey]; cacheHit {
		return cache[cacheKey]
	}

	if node == "end" {
		return 1
	}

	var paths int
	for adjacent := range graph[node] {
		if adjacent == "start" {
			continue
		}
		if unicode.IsLower(rune(adjacent[0])) {
			vTimes := visitedTimes(path, adjacent)
			switch vTimes {
			case 0:
			case 1:
				if alreadyDoubleVisited {
					continue
				}
				paths += gold(adjacent, append(path, adjacent), true, graph, cache)
				continue
			default:
				continue
			}
		}
		paths += gold(adjacent, append(path, adjacent), alreadyDoubleVisited, graph, cache)
	}
	cache[cacheKey] = paths
	return paths
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
