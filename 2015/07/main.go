package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Table map[string]Instruction
type Operand string
type Instruction struct {
	Raw               string
	InstructionsTable *Table
}
type State map[Operand]uint16

// id -> dependencies
type Graph map[string][]string

func dependencyGraph(intructions []Instruction, g Graph) Graph {
	if len(intructions) == 0 {
		return g
	}

	next := intructions[0]
	operand, dependencies := next.Dependencies()
	g[string(operand)] = dependencies
	return dependencyGraph(intructions[1:], g)
}

func (i Instruction) Dependencies() (dst Operand, dependencies []string) {
	switch {
	case strings.Contains(i.Raw, "AND"):
		var src1, src2 Operand
		fmt.Sscanf(i.Raw, "%s AND %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.Contains(i.Raw, "OR"):
		var src1, src2 Operand
		fmt.Sscanf(i.Raw, "%s OR %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.Contains(i.Raw, "LSHIFT"):
		var src1, src2 Operand
		fmt.Sscanf(i.Raw, "%s LSHIFT %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.Contains(i.Raw, "RSHIFT"):
		var src1, src2 Operand
		fmt.Sscanf(i.Raw, "%s RSHIFT %s -> %s", &src1, &src2, &dst)
		if src1.IsIdentifier() {
			dependencies = append(dependencies, string(src1))
		}
		if src2.IsIdentifier() {
			dependencies = append(dependencies, string(src2))
		}

	case strings.HasPrefix(i.Raw, "NOT"):
		var src Operand
		fmt.Sscanf(i.Raw, "NOT %s -> %s", &src, &dst)
		if src.IsIdentifier() {
			dependencies = append(dependencies, string(src))
		}

	default:
		var src Operand
		fmt.Sscanf(i.Raw, "%s -> %s", &src, &dst)
		if src.IsIdentifier() {
			dependencies = append(dependencies, string(src))
		}
	}
	return
}

func (o Operand) IsIdentifier() bool {
	_, err := strconv.ParseUint(string(o), 10, 16)
	return err != nil
}

func (t Table) IdValue(op Operand) uint16 {
	result, err := strconv.ParseUint(string(op), 10, 16)
	if err == nil {
		return uint16(result)
	}
	return t[string(op)].Value()
}

var memoization = make(map[string]uint16)

func (i Instruction) Value() uint16 {
	// fmt.Printf("value: %#v\n", i)
	if i.Raw == "" {
		panic("empty value")
	}
	if res, ok := memoization[i.Raw]; ok {
		return res
	}

	var res uint16
	switch {
	case strings.Contains(i.Raw, "AND"):
		var src1, src2, dst Operand
		fmt.Sscanf(i.Raw, "%s AND %s -> %s", &src1, &src2, &dst)
		res = i.InstructionsTable.IdValue(src1) & i.InstructionsTable.IdValue(src2)

	case strings.Contains(i.Raw, "OR"):
		var src1, src2, dst Operand
		fmt.Sscanf(i.Raw, "%s OR %s -> %s", &src1, &src2, &dst)
		res = i.InstructionsTable.IdValue(src1) | i.InstructionsTable.IdValue(src2)

	case strings.Contains(i.Raw, "LSHIFT"):
		var src1, src2, dst Operand
		fmt.Sscanf(i.Raw, "%s LSHIFT %s -> %s", &src1, &src2, &dst)
		res = i.InstructionsTable.IdValue(src1) << i.InstructionsTable.IdValue(src2)

	case strings.Contains(i.Raw, "RSHIFT"):
		var src1, src2, dst Operand
		fmt.Sscanf(i.Raw, "%s RSHIFT %s -> %s", &src1, &src2, &dst)
		res = i.InstructionsTable.IdValue(src1) >> i.InstructionsTable.IdValue(src2)

	case strings.HasPrefix(i.Raw, "NOT"):
		var src, dst Operand
		fmt.Sscanf(i.Raw, "NOT %s -> %s", &src, &dst)
		res = i.InstructionsTable.IdValue(src) ^ 65_535

	default:
		var src, dst Operand
		fmt.Sscanf(i.Raw, "%s -> %s", &src, &dst)
		res = i.InstructionsTable.IdValue(src)
	}

	memoization[i.Raw] = res
	return res
}

func run(id string, instructionsById map[string]Instruction) int {
	return 0
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
		instuctions = append(instuctions, Instruction{Raw: scanner.Text()})
	}

	instructionsById := make(Table)
	for _, instruction := range instuctions {
		_, id, _ := strings.Cut(string(instruction.Raw), "-> ")
		instruction.InstructionsTable = &instructionsById
		instructionsById[id] = instruction
	}

	fmt.Println(instructionsById["a"].Value())

	return nil
}

func main() {
	if err := solve(); err != nil {
		log.Fatal(err)
	}
}
