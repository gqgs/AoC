package main

import (
	"bufio"
	"os"
)

func priority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 96)
	}
	return int(c - 38)
}

func overlap(left, right string) int {
	set := make(map[rune]struct{})
	for _, c := range left {
		set[c] = struct{}{}
	}
	for _, c := range right {
		if _, s := set[c]; s {
			return priority(c)
		}
	}
	panic("not found")
}

func silver(list [][2]string) int {
	var total int
	for _, l := range list {
		left, right := l[0], l[1]
		total += overlap(left, right)
	}
	return total
}

func overlap3(l1, l2, l3 string) int {
	set1 := make(map[rune]struct{})
	set2 := make(map[rune]struct{})
	for _, c := range l1 {
		set1[c] = struct{}{}
	}
	for _, c := range l2 {
		set2[c] = struct{}{}
	}
	for _, c := range l3 {
		if _, s1 := set1[c]; s1 {
			if _, s2 := set2[c]; s2 {
				return priority(c)
			}
		}
	}
	panic("not found")
}

func gold(list []string) int {
	var total int
	for i := 0; i < len(list); i += 3 {
		total += overlap3(list[i], list[i+1], list[i+2])
	}
	return total
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var list [][2]string
	var fullList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		fullList = append(fullList, next)
		left, right := next[:len(next)/2], next[len(next)/2:]
		list = append(list, [2]string{left, right})
	}

	println("silver:", silver(list))
	println("gold:", gold(fullList))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
