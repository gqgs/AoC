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

	println("silver:", silver("start", []string{"start"}, graph))
	println("gold:", gold("start", []string{"start"}, false, graph))

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

func gold(node string, path []string, alreadyDoubleVisited bool, graph map[string]map[string]struct{}) int {
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
				paths += gold(adjacent, append(path, adjacent), true, graph)
				continue
			default:
				continue
			}
		}
		paths += gold(adjacent, append(path, adjacent), alreadyDoubleVisited, graph)
	}
	return paths
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
