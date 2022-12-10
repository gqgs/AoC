package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gqgs/AoC2021/generic"
)

type stateMachine struct {
	X            int
	Cycle        int
	MonitorCycle []int
	Crt          string
}

func newStateMachine() stateMachine {
	return stateMachine{
		X:            1,
		Cycle:        1,
		MonitorCycle: []int{20, 60, 100, 140, 180, 220},
	}
}

func (sm stateMachine) signalStrength() int {
	for _, c := range sm.MonitorCycle {
		if sm.Cycle == c {
			return sm.X * c
		}
	}
	return 0
}

func (sm *stateMachine) updateCrt() {
	next := "."
	if generic.NewSet(sm.X-1, sm.X, sm.X+1).Contains(len(sm.Crt) % 40) {
		next = "#"
	}
	sm.Crt = sm.Crt + next
}

func (sm stateMachine) printCrt(w io.Writer) {
	fmt.Fprintln(w)
	for i := 0; i < 6; i++ {
		fmt.Fprintln(w, sm.Crt[40*i:40*i+40])
	}
}

func (sm *stateMachine) executeInst(inst string) int {
	var total int
	switch inst[0] {
	case 'n':
		sm.updateCrt()
		sm.Cycle++
		total += sm.signalStrength()
	case 'a':
		var value int
		fmt.Sscanf(inst, "addx %d", &value)
		sm.updateCrt()
		sm.Cycle++
		total += sm.signalStrength()
		sm.updateCrt()
		sm.X += value
		sm.Cycle++
		total += sm.signalStrength()
	}
	return total
}

func silver(insts []string) int {
	sm := newStateMachine()
	var total int
	for _, inst := range insts {
		total += sm.executeInst(inst)
	}
	return total
}

func gold(insts []string) string {
	sm := newStateMachine()
	for _, inst := range insts {
		sm.executeInst(inst)
	}

	builder := new(strings.Builder)
	sm.printCrt(builder)
	return builder.String()
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var insts []string
	for scanner.Scan() {
		next := scanner.Text()
		insts = append(insts, next)
	}

	println("silver:", silver(insts))
	println("gold:", gold(insts))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
