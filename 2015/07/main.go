package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operand string
type Instruction string
type State map[Operand]uint16

// id -> dependencies
type Graph map[string][]string

func dependencyGraph(intructions []Instruction, g Graph) Graph {
	if len(intructions) == 0 {
		return g
	}

	next := intructions[0]
	g[string(next)] = next.Dependencies()
	return dependencyGraph(intructions[1:], g)
}

func (o Operand) Value(state State) uint16 {
	result, err := strconv.ParseUint(string(o), 10, 16)
	if err == nil {
		return uint16(result)
	}
	return state[o]
}

func (o Operand) IsIdentifier() bool {
	_, err := strconv.ParseUint(string(o), 10, 16)
	return err != nil
}

func (i Instruction) Dependencies() (dependencies []string) {
	switch {
	case strings.Contains(string(i), "AND"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s AND %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.Contains(string(i), "OR"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s OR %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.Contains(string(i), "LSHIFT"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s LSHIFT %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.Contains(string(i), "RSHIFT"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s RSHIFT %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.HasPrefix(string(i), "NOT"):
		var src, dst Operand
		fmt.Sscanf(string(i), "NOT %s -> %s", &src, &dst)
		if src.IsIdentifier() {
			dependencies = append(dependencies, string(src))
		}

	default:
		var src, dst Operand
		fmt.Sscanf(string(i), "%s -> %s", &src, &dst)
		if src.IsIdentifier() {
			dependencies = append(dependencies, string(src))
		}
	}
	return
}

func (i Instruction) Eval(state State) State {
	switch {
	case strings.Contains(string(i), "AND"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s AND %s -> %s", &src1, &src2, &dst)
		state[dst] = src1.Value(state) & src2.Value(state)

	case strings.Contains(string(i), "OR"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s OR %s -> %s", &src1, &src2, &dst)
		state[dst] = src1.Value(state) | src2.Value(state)

	case strings.Contains(string(i), "LSHIFT"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s LSHIFT %s -> %s", &src1, &src2, &dst)
		state[dst] = src1.Value(state) << src2.Value(state)

	case strings.Contains(string(i), "RSHIFT"):
		var src1, src2, dst Operand
		fmt.Sscanf(string(i), "%s RSHIFT %s -> %s", &src1, &src2, &dst)
		state[dst] = src1.Value(state) >> src2.Value(state)

	case strings.HasPrefix(string(i), "NOT"):
		var src, dst Operand
		fmt.Sscanf(string(i), "NOT %s -> %s", &src, &dst)
		state[dst] = src.Value(state) ^ 65_535

	default:
		var src, dst Operand
		fmt.Sscanf(string(i), "%s -> %s", &src, &dst)
		state[dst] = src.Value(state)
	}
	return state
}

func run(instuctions []Instruction, state State) State {
	if len(instuctions) == 0 {
		return state
	}

	return run(instuctions[1:], instuctions[0].Eval(state))
}

func solve() error {
	file, err := os.Open("input")
	if err != nil {
		return fmt.Errorf("file not file: %w", err)
	}
	defer file.Close()

	var instuctions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instuctions = append(instuctions, Instruction(scanner.Text()))
	}

	// for key, value := range run(instuctions, make(State)) {
	// 	fmt.Println(key, value)
	// }

	for key, value := range dependencyGraph(instuctions, make(Graph)) {
		fmt.Println(key, value)
	}

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
