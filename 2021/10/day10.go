package main

import (
	"bufio"
	"os"

	"github.com/gqgs/AoC2021/generic"
)

var (
	matchingChars = map[rune]rune{
		'(': ')',
		'[': ']',
		'<': '>',
		'{': '}',
	}

	autoCompletePoint = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	errorPoint = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

type Stack = generic.Stack[rune]

func silver(input []string) ([]Stack, int) {
	var sum int
	invalidStack := make([]Stack, 0, len(input))
Next:
	for _, i := range input {
		stack := make(Stack, 0, len(i))
		for _, c := range i {
			switch c {
			case '(', '[', '<', '{':
				stack.Push(c)
			default:
				if c == matchingChars[stack.Pop()] {
					continue
				}
				sum += errorPoint[c]
				continue Next
			}
		}
		invalidStack = append(invalidStack, stack)
	}

	return invalidStack, sum
}

func gold(invalidStack []Stack) int {
	var sums []int
	for _, stack := range invalidStack {
		var sum int
		for len(stack) > 0 {
			next := stack.Pop()
			sum *= 5
			sum += autoCompletePoint[next]
		}
		sums = append(sums, sum)
	}

	return generic.QuickSelect(sums, len(sums)/2)
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	invalidStack, silverResult := silver(input)

	println("silver:", silverResult)
	println("gold:", gold(invalidStack))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
