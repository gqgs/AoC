package main

import (
	"fmt"
	"math/big"
)

// type Monkey struct {
// 	StartingItems []int
// 	Operation     func(value int) int
// 	Test          func(value int) bool
// 	IfTrue        func() int
// 	IfFalse       func() int
// }

type Monkey struct {
	StartingItems []*big.Int
	Operation     func(value *big.Int) *big.Int
	Test          func(value *big.Int) bool
	IfTrue        func() int
	IfFalse       func() int
}

func IsDivisible(x *big.Int, y int) bool {
	xx := new(big.Int).SetBytes(x.Bytes())
	return xx.Mod(xx, big.NewInt(int64(y))).Int64() == 0
}

func solve() error {
	// monkeys := []*Monkey{
	// 	{
	// 		StartingItems: []int{79, 98},
	// 		Operation:     func(old int) int { return old * 19 },
	// 		Test:          func(value int) bool { return value%23 == 0 },
	// 		IfTrue:        func() int { return 2 },
	// 		IfFalse:       func() int { return 3 },
	// 	},
	// 	{
	// 		StartingItems: []int{54, 65, 75, 74},
	// 		Operation:     func(old int) int { return old + 6 },
	// 		Test:          func(value int) bool { return value%19 == 0 },
	// 		IfTrue:        func() int { return 2 },
	// 		IfFalse:       func() int { return 0 },
	// 	},
	// 	{
	// 		StartingItems: []int{79, 60, 97},
	// 		Operation:     func(old int) int { return old * old },
	// 		Test:          func(value int) bool { return value%13 == 0 },
	// 		IfTrue:        func() int { return 1 },
	// 		IfFalse:       func() int { return 3 },
	// 	},
	// 	{
	// 		StartingItems: []int{74},
	// 		Operation:     func(old int) int { return old + 3 },
	// 		Test:          func(value int) bool { return value%17 == 0 },
	// 		IfTrue:        func() int { return 0 },
	// 		IfFalse:       func() int { return 1 },
	// 	},
	// }

	// monkeys := []*Monkey{
	// 	{
	// 		StartingItems: []*big.Int{big.NewInt(79), big.NewInt(98)},
	// 		Operation:     func(old *big.Int) *big.Int { return old.Mul(old, big.NewInt(19)) },
	// 		Test:          func(value *big.Int) bool { return IsDivisible(value, 23) },
	// 		IfTrue:        func() int { return 2 },
	// 		IfFalse:       func() int { return 3 },
	// 	},
	// 	{
	// 		StartingItems: []*big.Int{big.NewInt(54), big.NewInt(65), big.NewInt(75), big.NewInt(74)},
	// 		Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(6)) },
	// 		Test:          func(value *big.Int) bool { return IsDivisible(value, 19) },
	// 		IfTrue:        func() int { return 2 },
	// 		IfFalse:       func() int { return 0 },
	// 	},
	// 	{
	// 		StartingItems: []*big.Int{big.NewInt(79), big.NewInt(60), big.NewInt(97)},
	// 		Operation:     func(old *big.Int) *big.Int { return old.Mul(old, old) },
	// 		Test:          func(value *big.Int) bool { return IsDivisible(value, 13) },
	// 		IfTrue:        func() int { return 1 },
	// 		IfFalse:       func() int { return 3 },
	// 	},
	// 	{
	// 		StartingItems: []*big.Int{big.NewInt(74)},
	// 		Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(3)) },
	// 		// Test:          func(value *big.Int) bool { return value.Mod(value, big.NewInt(17)).Int64() == 0 },
	// 		Test:    func(value *big.Int) bool { return IsDivisible(value, 17) },
	// 		IfTrue:  func() int { return 0 },
	// 		IfFalse: func() int { return 1 },
	// 	},
	// }

	// monkeys := []*Monkey{
	// 	{
	// 		StartingItems: []*big.Int{big.NewInt(65), big.NewInt(78)},
	// 		Operation:     func(old int) int { return old * 3 },
	// 		Test:          func(value int) bool { return value%5 == 0 },
	// 		IfTrue:        func() int { return 2 },
	// 		IfFalse:       func() int { return 3 },
	// 	},
	// 	{
	// 		StartingItems: []int{54, 78, 86, 79, 73, 64, 85, 88},
	// 		Operation:     func(old int) int { return old + 8 },
	// 		Test:          func(value int) bool { return value%11 == 0 },
	// 		IfTrue:        func() int { return 4 },
	// 		IfFalse:       func() int { return 7 },
	// 	},
	// 	{
	// 		StartingItems: []int{69, 97, 77, 88, 87},
	// 		Operation:     func(old int) int { return old + 2 },
	// 		Test:          func(value int) bool { return value%2 == 0 },
	// 		IfTrue:        func() int { return 5 },
	// 		IfFalse:       func() int { return 3 },
	// 	},
	// 	{
	// 		StartingItems: []int{99},
	// 		Operation:     func(old int) int { return old + 4 },
	// 		Test:          func(value int) bool { return value%13 == 0 },
	// 		IfTrue:        func() int { return 1 },
	// 		IfFalse:       func() int { return 5 },
	// 	},
	// 	{
	// 		StartingItems: []int{60, 57, 52},
	// 		Operation:     func(old int) int { return old * 19 },
	// 		Test:          func(value int) bool { return value%7 == 0 },
	// 		IfTrue:        func() int { return 7 },
	// 		IfFalse:       func() int { return 6 },
	// 	},
	// 	{
	// 		StartingItems: []int{91, 82, 85, 73, 84, 53},
	// 		Operation:     func(old int) int { return old + 5 },
	// 		Test:          func(value int) bool { return value%3 == 0 },
	// 		IfTrue:        func() int { return 4 },
	// 		IfFalse:       func() int { return 1 },
	// 	},
	// 	{
	// 		StartingItems: []int{88, 74, 68, 56},
	// 		Operation:     func(old int) int { return old * old },
	// 		Test:          func(value int) bool { return value%17 == 0 },
	// 		IfTrue:        func() int { return 0 },
	// 		IfFalse:       func() int { return 2 },
	// 	},
	// 	{
	// 		StartingItems: []int{54, 82, 72, 71, 53, 99, 67},
	// 		Operation:     func(old int) int { return old + 1 },
	// 		Test:          func(value int) bool { return value%19 == 0 },
	// 		IfTrue:        func() int { return 6 },
	// 		IfFalse:       func() int { return 0 },
	// 	},
	// }

	monkeys := []*Monkey{
		{
			StartingItems: []*big.Int{big.NewInt(65), big.NewInt(78)},
			Operation:     func(old *big.Int) *big.Int { return old.Mul(old, big.NewInt(3)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 5) },
			IfTrue:        func() int { return 2 },
			IfFalse:       func() int { return 3 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(54), big.NewInt(78), big.NewInt(86), big.NewInt(79), big.NewInt(73), big.NewInt(64), big.NewInt(85), big.NewInt(88)},
			Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(8)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 11) },
			IfTrue:        func() int { return 4 },
			IfFalse:       func() int { return 7 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(69), big.NewInt(97), big.NewInt(77), big.NewInt(88), big.NewInt(87)},
			Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(2)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 2) },
			IfTrue:        func() int { return 5 },
			IfFalse:       func() int { return 3 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(99)},
			Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(4)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 13) },
			IfTrue:        func() int { return 1 },
			IfFalse:       func() int { return 5 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(60), big.NewInt(57), big.NewInt(52)},
			Operation:     func(old *big.Int) *big.Int { return old.Mul(old, big.NewInt(19)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 7) },
			IfTrue:        func() int { return 7 },
			IfFalse:       func() int { return 6 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(91), big.NewInt(82), big.NewInt(85), big.NewInt(73), big.NewInt(84), big.NewInt(53)},
			Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(5)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 3) },
			IfTrue:        func() int { return 4 },
			IfFalse:       func() int { return 1 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(88), big.NewInt(74), big.NewInt(68), big.NewInt(56)},
			Operation:     func(old *big.Int) *big.Int { return old.Mul(old, old) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 17) },
			IfTrue:        func() int { return 0 },
			IfFalse:       func() int { return 2 },
		},
		{
			StartingItems: []*big.Int{big.NewInt(54), big.NewInt(82), big.NewInt(72), big.NewInt(71), big.NewInt(53), big.NewInt(99), big.NewInt(67)},
			Operation:     func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(1)) },
			Test:          func(value *big.Int) bool { return IsDivisible(value, 19) },
			IfTrue:        func() int { return 6 },
			IfFalse:       func() int { return 0 },
		},
	}

	inspectCount := make(map[int]int)
	for c := 0; c < 10000; c++ {
		for mi, m := range monkeys {
			inspectCount[mi] += len(m.StartingItems)
			for i, item := range m.StartingItems {
				//fmt.Println(mi, m.StartingItems[i].String())
				m.StartingItems[i] = m.Operation(item)
				//fmt.Println(mi, m.StartingItems[i].String(), "op")
				// m.StartingItems[i].Div(m.StartingItems[i], big.NewInt(3))
				// m.StartingItems[i] = int(m.StartingItems[i] / 3)
				m.StartingItems[i].Mod(m.StartingItems[i], big.NewInt(9699690))
				//fmt.Println(mi, m.StartingItems[i].String(), "div3", m.Test(m.StartingItems[i]))
				var dst int
				if m.Test(m.StartingItems[i]) {
					dst = m.IfTrue()
				} else {
					dst = m.IfFalse()
				}
				//fmt.Println(mi, m.StartingItems[i].String(), "append")
				monkeys[dst].StartingItems = append(monkeys[dst].StartingItems, m.StartingItems[i])
			}
			m.StartingItems = nil
		}
	}

	for i, m := range monkeys {
		// fmt.Println(i, m.StartingItems)
		print(i, " [ ")
		for _, c := range m.StartingItems {
			print(c.String(), " ")
		}
		println("]")
	}
	println()

	for i, m := range inspectCount {
		fmt.Println(i, m)
	}

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
