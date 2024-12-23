package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"sort"
	"strings"
)

func parseGraph(lines []string) map[string]map[string]struct{} {
	graph := make(map[string]map[string]struct{})
	for _, line := range lines {
		parts := strings.Split(line, "-")
		left := parts[0]
		right := parts[1]
		if graph[left] == nil {
			graph[left] = make(map[string]struct{})
		}
		graph[left][right] = struct{}{}
		if graph[right] == nil {
			graph[right] = make(map[string]struct{})
		}
		graph[right][left] = struct{}{}
	}
	return graph
}

func silver(lines []string) int {
	graph := parseGraph(lines)
	isConnected := func(n string, g map[string]struct{}) bool {
		_, r := g[n]
		return r
	}

	results := make(map[string]struct{})
	for k1, v1 := range graph {
		for k2, v2 := range graph {
			for k3, v3 := range graph {
				if k1[0] != 't' && k2[0] != 't' && k3[0] != 't' {
					continue
				}
				connected := isConnected(k2, v1) &&
					isConnected(k3, v1) &&
					isConnected(k1, v2) &&
					isConnected(k3, v2) &&
					isConnected(k1, v3) &&
					isConnected(k2, v3)
				if connected {
					nodes := []string{k1, k2, k3}
					sort.Strings(nodes)
					key := fmt.Sprintf("%s-%s-%s", nodes[0], nodes[1], nodes[2])
					results[key] = struct{}{}
				}
			}
		}
	}

	return len(results)
}

func connectedSet(current string, graph map[string]map[string]struct{}, visited map[string]struct{}) {
	// invariant: visited = connected nodes
	if _, ok := visited[current]; ok {
		return
	}
	visited[current] = struct{}{}

Next:
	for next := range graph[current] {
		for connected := range visited {
			if _, isConnected := graph[next][connected]; !isConnected {
				// this node is not reacheable from all nodes
				continue Next
			}
		}
		// node is part of connected set
		connectedSet(next, graph, visited)
	}
}

func gold(lines []string) string {
	graph := parseGraph(lines)
	var maxConnected int
	var password []string
	for node := range graph {
		connected := make(map[string]struct{})
		connectedSet(node, graph, connected)
		if len(connected) > maxConnected {
			maxConnected = len(connected)
			password = slices.Collect(maps.Keys(connected))
		}
	}

	sort.Strings(password)
	return strings.Join(password, ",")
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

	println(silver(lines))
	println(gold(lines))

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
