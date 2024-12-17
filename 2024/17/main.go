package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gqgs/AoC2021/ints"
)

func silver(lines []string) string {
	var registerA int
	fmt.Sscanf(lines[0], "Register A: %d", &registerA)
	target := strings.TrimPrefix(lines[3], "Program: ")
	program := ints.FromString(target, ",")
	return shared(registerA, program)
}

func gold(lines []string) int {
	target := strings.TrimPrefix(lines[3], "Program: ")
	program := ints.FromString(target, ",")
	for i := 0; ; {
		result := shared(i, program)
		if result == target {
			return i
		}
		var j int
		for j = 0; j < len(result) && result[len(result)-1-j] == target[len(target)-1-j]; j++ {
		}
		r := ints.FromString(result, ",")
		i += 1 << (3 * (len(r) - j/2 - 1))
	}
}

func shared(registerA int, program []int) string {
	registers := make(map[rune]int)
	registers['A'] = registerA
	combo := func(c int) int {
		switch c {
		case 0, 1, 2, 3:
			return c
		case 4:
			return registers['A']
		case 5:
			return registers['B']
		case 6:
			return registers['C']
		default:
			panic("invalid:" + strconv.Itoa(c))
		}
	}
	var outs []int
	for pc := 0; pc < len(program); {
		switch program[pc] {
		case 0: // adv
			registers['A'] /= 1 << combo(program[pc+1])
			pc += 2
		case 1: // bxl
			registers['B'] ^= program[pc+1]
			pc += 2
		case 2: // bst
			registers['B'] = combo(program[pc+1]) % 8
			pc += 2
		case 3: // jnz
			if registers['A'] == 0 {
				pc += 2
				continue
			}
			pc = program[pc+1]
		case 4: // bxn
			registers['B'] ^= registers['C']
			pc += 2
		case 5: // out
			rr := combo(program[pc+1]) % 8
			outs = append(outs, rr)
			pc += 2
		case 6: // bdv
			registers['B'] = registers['A'] / (1 << combo(program[pc+1]))
			pc += 2
		case 7: // cdv
			registers['C'] = registers['A'] / (1 << combo(program[pc+1]))
			pc += 2
		default:
			panic("invalid op:" + fmt.Sprint(program[pc]))
		}

	}

	var result []string
	for _, o := range outs {
		result = append(result, strconv.Itoa(o))
	}
	return strings.Join(result, ",")
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
		panic(err)
	}
}
