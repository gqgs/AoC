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

func (o Operand) Value(state State) uint16 {
	result, err := strconv.ParseUint(string(o), 10, 16)
	if err == nil {
		return uint16(result)
	}
	return state[o]
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

	for key, value := range run(instuctions, make(State)) {
		fmt.Println(key, value)
	}

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
