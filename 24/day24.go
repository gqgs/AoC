package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func generateZ3(w io.Writer, instructions []string) {
	ops := make(map[string]int)
	lastOP := func(op string) string {
		return op + fmt.Sprint(ops[op])
	}

	newOP := func(op string) string {
		ops[op]++
		fmt.Fprintf(w, "(declare-const %s Int)\n", op+fmt.Sprint(ops[op]))
		return op + fmt.Sprint(ops[op])
	}

	parseOperand := func(op string) string {
		if _, err := strconv.Atoi(op); err == nil {
			return op
		}
		return lastOP(op)
	}

	fmt.Fprintln(w, "(declare-const x0 Int)")
	fmt.Fprintln(w, "(declare-const y0 Int)")
	fmt.Fprintln(w, "(declare-const w0 Int)")
	fmt.Fprintln(w, "(declare-const z0 Int)")

	fmt.Fprintln(w, "(assert (= x0 0))")
	fmt.Fprintln(w, "(assert (= y0 0))")
	fmt.Fprintln(w, "(assert (= z0 0))")
	fmt.Fprintln(w, "(assert (= w0 0))")

	var variables []string
	for _, instruction := range instructions {
		switch {
		case strings.HasPrefix(instruction, "inp"):
			variables = append(variables, newOP("w"))
		case strings.HasPrefix(instruction, "add"):
			var op1, op2 string
			fmt.Sscanf(instruction, "add %s %s", &op1, &op2)
			last := lastOP(op1)
			new := newOP(op1)
			fmt.Fprintf(w, "(assert (= %s (+ %s %s)))\n", new, last, parseOperand(op2))
		case strings.HasPrefix(instruction, "mul"):
			var op1, op2 string
			fmt.Sscanf(instruction, "mul %s %s", &op1, &op2)
			last := lastOP(op1)
			new := newOP(op1)
			fmt.Fprintf(w, "(assert (= %s (* %s %s)))\n", new, last, parseOperand(op2))
		case strings.HasPrefix(instruction, "div"):
			var op1, op2 string
			fmt.Sscanf(instruction, "div %s %s", &op1, &op2)
			last := lastOP(op1)
			new := newOP(op1)
			fmt.Fprintf(w, "(assert (= %s (div %s %s)))\n", new, last, parseOperand(op2))
		case strings.HasPrefix(instruction, "mod"):
			var op1, op2 string
			fmt.Sscanf(instruction, "mod %s %s\n", &op1, &op2)
			last := lastOP(op1)
			new := newOP(op1)
			fmt.Fprintf(w, "(assert (= %s (mod %s %s)))\n", new, last, parseOperand(op2))
		case strings.HasPrefix(instruction, "eql"):
			var op1, op2 string
			fmt.Sscanf(instruction, "eql %s %s", &op1, &op2)
			last := lastOP(op1)
			new := newOP(op1)
			fmt.Fprintf(w, "(assert (= %s (ite (= %s %s) 1  0)))\n", new, last, parseOperand(op2))
		}
	}

	fmt.Fprintf(w, "(assert (= %s 0))\n", lastOP("z"))

	var model []string
	for i, v := range variables {
		fmt.Fprintf(w, "(assert (and (>= %[1]s 1) (<= %[1]s 9)))\n", v)
		model = append(model, fmt.Sprintf("* %s %.0f", v, math.Pow10(i)))
	}

	fmt.Fprintf(w, "(declare-const model Int)\n")
	fmt.Fprintf(w, "(assert (= model (+ (%s))))\n", strings.Join(model, ") ("))

	// TODO: without the condition bellow z3 will find a solution to the problem not the optimal solution.
	// fmt.Fprintln(w, "(maximize model)")
	fmt.Fprintln(w, "(check-sat)")
	fmt.Fprintln(w, "(get-value ((as model Int)))")
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	generateZ3(os.Stdout, instructions)

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
