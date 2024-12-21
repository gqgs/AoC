package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gqgs/AoC2021/grid"
)

var keyPad = map[rune]grid.Point{
	'7': {X: 0, Y: 0},
	'8': {X: 0, Y: 1},
	'9': {X: 0, Y: 2},
	'4': {X: 1, Y: 0},
	'5': {X: 1, Y: 1},
	'6': {X: 1, Y: 2},
	'1': {X: 2, Y: 0},
	'2': {X: 2, Y: 1},
	'3': {X: 2, Y: 2},
	'0': {X: 3, Y: 1},
	'A': {X: 3, Y: 2},
}

func validTranformation(line string, prev, next, invalid grid.Point) bool {
	if len(line) == 0 {
		return false
	}
	for _, c := range line {
		switch c {
		case '^':
			prev.X--
		case 'v':
			prev.X++
		case '<':
			prev.Y--
		case '>':
			prev.Y++
		}
		if prev.Equal(invalid) {
			return false
		}
	}

	return prev.Equal(next)
}

func directions(prev, next, invalid grid.Point) []string {
	distance := prev.Distance(next)
	results := []string{""}
	for ; distance > 0; distance-- {
		var updated []string
		for _, r := range results {
			for _, d := range []string{"^", ">", "v", "<"} {
				updated = append(updated, r+d)
			}
		}
		results = updated
	}

	var valid []string
	for _, r := range results {
		if validTranformation(r, prev, next, invalid) {
			valid = append(valid, r)
		}
	}

	return valid
}

func solveKeyPad(line string) string {
	invalid := grid.Point{X: 3, Y: 0}
	prev := 'A'
	var result strings.Builder
	for _, c := range line {
		next := c
		prevPoint := keyPad[prev]
		nextPoint := keyPad[next]

		results := directions(prevPoint, nextPoint, invalid)
		r := results[0]
		if len(results) > 1 {
			r = "(" + strings.Join(results, ",") + ")"
		}
		result.WriteString(r)
		result.WriteRune('A')

		prev = next
	}
	return result.String()
}

var directional = map[rune]grid.Point{
	'^': {X: 0, Y: 1},
	'A': {X: 0, Y: 2},
	'<': {X: 1, Y: 0},
	'v': {X: 1, Y: 1},
	'>': {X: 1, Y: 2},
}

type Case struct {
	Directions []string
}

func solveDirectional(line string, depth int, cache map[string]int) int {
	key := fmt.Sprintf("%s,%d", line, depth)
	if cached, ok := cache[key]; ok {
		return cached
	}
	if depth == 0 {
		return len(line)
	}

	prev := 'A'
	invalid := grid.Point{X: 0, Y: 0}
	var cases []Case
	for _, c := range line {
		next := c
		prevPoint := directional[prev]
		nextPoint := directional[next]

		results := directions(prevPoint, nextPoint, invalid)
		cases = append(cases, Case{
			Directions: results,
		})

		prev = next
	}

	var length int
	for _, c := range cases {
		var mindirections []int
		for _, d := range c.Directions {
			mindirections = append(mindirections, solveDirectional(d+"A", depth-1, cache))
		}

		if len(mindirections) > 0 {
			length += slices.Min(mindirections)
		} else {
			length += 1
		}
	}

	cache[key] = length
	return length
}

func expandCollapsedString(s string) []string {
	startIndex := strings.Index(s, "(")
	endIndex := strings.Index(s, ")")
	if startIndex == -1 {
		return []string{s}
	}
	var results []string
	parts := strings.Split(s[startIndex+1:endIndex], ",")
	for _, p := range parts {
		results = append(results, s[:startIndex]+p+s[endIndex+1:])
	}

	var updated []string
	for _, r := range results {
		if strings.Contains(r, "(") {
			updated = append(updated, expandCollapsedString(r)...)
		} else {
			updated = append(updated, r)
		}
	}

	return updated
}

func silver(lines []string) int {
	return shared(lines, 2)
}

func gold(lines []string) int {
	return shared(lines, 25)
}

func shared(lines []string, depth int) int {
	var total int
	cache := make(map[string]int)
	for _, line := range lines {
		var results []int
		for _, c := range expandCollapsedString(solveKeyPad(line)) {
			results = append(results, solveDirectional(c, depth, cache))

		}
		n, _ := strconv.Atoi(line[:3])
		total += n * slices.Min(results)
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
