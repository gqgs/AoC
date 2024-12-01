package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strings"
)

type Number struct {
	Value  float64
	Left   *Number
	Right  *Number
	Parent *Number

	Prev *Number
	Next *Number
}

func (n *Number) String() string {
	if n.Left == nil && n.Right == nil {
		return fmt.Sprint(n.Value)
	}
	var str []string
	if n.Left != nil {
		str = append(str, n.Left.String())
	}
	if n.Right != nil {
		str = append(str, n.Right.String())
	}
	return "[" + strings.Join(str, ",") + "]"
}

func (n *Number) updateList(prev *Number) *Number {
	if n.Left == nil && n.Right == nil {
		if prev == nil {
			return n
		}
		n.Prev = prev
		prev.Next = n
		return n
	}

	if n.Left != nil {
		prev = n.Left.updateList(prev)
	}

	if n.Right != nil {
		prev = n.Right.updateList(prev)
	}

	return prev
}

func (n Number) Max() float64 {
	if n.Left == nil && n.Right == nil {
		return n.Value
	}

	var values []float64

	if n.Left != nil {
		values = append(values, n.Left.Max())
	}

	if n.Right != nil {
		values = append(values, n.Right.Max())
	}

	return slices.Max(values)
}

func (n Number) Nested() int {
	if n.Left == nil && n.Right == nil {
		return 0
	}
	var depths []int
	if n.Left != nil {
		depths = append(depths, n.Left.Nested())
	}
	if n.Right != nil {
		depths = append(depths, n.Right.Nested())
	}
	return 1 + slices.Max(depths)
}

func (n Number) magnitude() float64 {
	if n.Left == nil && n.Right == nil {
		return n.Value
	}
	var left, right float64
	if n.Left != nil {
		left = n.Left.magnitude()
	}
	if n.Right != nil {
		right = n.Right.magnitude()
	}
	return 3*left + 2*right
}

func decodeReader(reader io.Reader) *Number {
	numbers := make([]interface{}, 0)
	json.NewDecoder(reader).Decode(&numbers)
	number := decode(numbers)
	number.updateList(nil)
	return number
}

func decode(numbers []interface{}) *Number {
	number := new(Number)
	if value, ok := numbers[0].(float64); ok {
		number.Left = &Number{
			Value:  value,
			Parent: number,
		}
	} else {
		number.Left = decode(numbers[0].([]interface{}))
		number.Left.Parent = number
	}

	if value, ok := numbers[1].(float64); ok {
		number.Right = &Number{
			Value:  value,
			Parent: number,
		}
	} else {
		number.Right = decode(numbers[1].([]interface{}))
		number.Right.Parent = number
	}
	return number
}

func reduce(n *Number) *Number {
	for n.Nested() > 4 || n.Max() > 9 {
		if n.Left.Nested() >= 4 {
			explode(n.Left, 4)
			n.updateList(nil)
			continue
		}
		if n.Right.Nested() >= 4 {
			explode(n.Right, 4)
			n.updateList(nil)
			continue
		}

		if n.Left.Max() > 9 {
			split(n.Left)
			n.updateList(nil)
			continue
		}

		if n.Right.Max() > 9 {
			split(n.Right)
			n.updateList(nil)
			continue
		}
	}

	return n
}

func split(n *Number) bool {
	if n.Left == nil && n.Right == nil {
		if n.Value <= 9 {
			return false
		}
		n.Left = &Number{
			Value:  math.Floor(n.Value / 2.0),
			Parent: n,
		}
		n.Right = &Number{
			Value:  math.Ceil(n.Value / 2.0),
			Parent: n,
		}
		n.Value = 0
		return true
	}

	if n.Left != nil {
		if split(n.Left) {
			return true
		}
	}

	if n.Right != nil {
		return split(n.Right)
	}

	return false
}

func explode(n *Number, depth int) bool {
	if depth == 0 {
		left := n.Parent.Left.Prev
		for left != nil {
			if left.Nested() == 0 {
				left.Value += n.Parent.Left.Value
				break
			}
			left = left.Prev
		}
		right := n.Parent.Right.Next
		for right != nil {
			if right.Nested() == 0 {
				right.Value += n.Parent.Right.Value
				break
			}
			right = right.Next
		}
		if n.Parent.Parent.Left.Nested() == 1 {
			n.Parent.Parent.Left.Left = nil
			n.Parent.Parent.Left.Right = nil
		} else {
			n.Parent.Parent.Right.Left = nil
			n.Parent.Parent.Right.Right = nil
		}
		return true
	}

	if n.Left == nil && n.Right == nil {
		return false
	}

	if n.Left != nil {
		if explode(n.Left, depth-1) {
			return true
		}
	}

	if n.Right != nil {
		return explode(n.Right, depth-1)
	}

	return false
}

func add(left, right *Number) *Number {
	parent := new(Number)
	left.Parent = parent
	right.Parent = parent
	parent.Left = left
	parent.Right = right
	parent.updateList(nil)
	return reduce(parent)
}

func silver(input []string) {
	var numbers []*Number
	for _, s := range input {
		numbers = append(numbers, decodeReader(strings.NewReader(s)))
	}

	sum := numbers[0]
	for _, number := range numbers[1:] {
		sum = add(sum, number)
	}

	fmt.Println("silver:", sum, sum.magnitude())
}

func gold(input []string) {
	var magnitudes []int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			number1 := decodeReader(strings.NewReader(input[i]))
			number2 := decodeReader(strings.NewReader(input[j]))
			magnitudes = append(magnitudes, int(add(number1, number2).magnitude()))
		}
	}

	println("gold:", slices.Max(magnitudes))
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

	silver(input)
	gold(input)

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
