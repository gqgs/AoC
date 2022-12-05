package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

func silver(stacks []generic.Stack[byte], moves [][3]int) string {
	for _, stack := range stacks {
		var elems []byte
		for len(stack) > 0 {
			elems = append(elems, stack.Pop())
		}
		for _, elem := range elems {
			stack.Push(elem)
		}
	}

	for _, move := range moves {
		count, from, to := move[0], move[1], move[2]
		for ; count > 0; count-- {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	var result strings.Builder
	for _, stack := range stacks {
		result.WriteByte(stack.Pop())
	}
	return result.String()
}

func gold(stacks []generic.Stack[byte], moves [][3]int) string {
	for _, stack := range stacks {
		var elems []byte
		for len(stack) > 0 {
			elems = append(elems, stack.Pop())
		}
		for _, elem := range elems {
			stack.Push(elem)
		}
	}

	for _, move := range moves {
		count, from, to := move[0], move[1], move[2]
		stack := make(generic.Stack[byte], 0)
		for ; count > 0; count-- {
			stack.Push(stacks[from-1].Pop())
		}
		for len(stack) > 0 {
			stacks[to-1].Push(stack.Pop())
		}
	}

	var result strings.Builder
	for _, stack := range stacks {
		result.WriteByte(stack.Pop())
	}
	return result.String()
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var stacks []generic.Stack[byte]
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) == 0 {
			break
		}
		if len(stacks) == 0 {
			for i := 1; i < len(next); i += 4 {
				stacks = append(stacks, generic.Stack[byte]{})
			}
		}

		var j int
		for i := 1; i < len(next); i += 4 {
			if next[i] >= 'A' && next[i] <= 'Z' {
				stacks[j].Push(next[i])
			}
			j++
		}
	}

	var moves [][3]int
	for scanner.Scan() {
		next := scanner.Text()
		var count, from, to int
		fmt.Sscanf(next, "move %d from %d to %d\n", &count, &from, &to)
		moves = append(moves, [3]int{count, from, to})
	}

	stacksCopy := make([]generic.Stack[byte], len(stacks))
	for i, stack := range stacks {
		stacksCopy[i] = append(stacksCopy[i], stack...)
	}

	println("silver:", silver(stacks, moves))
	println("gold:", gold(stacksCopy, moves))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
