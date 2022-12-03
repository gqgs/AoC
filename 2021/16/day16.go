package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/gqgs/AoC2021/generic"
)

type Token struct {
	OP                 string
	Value              int
	Length             int
	SubpacketsLength   int
	RequiredSubpackets int
}

func tokenize(input string) ([]Token, int) {
	var tokens []Token
	var versionSum int
	for i := 0; i < len(input)-6; {
		version := input[i : i+3]
		id := input[i+3 : i+6]

		versionInt, _ := strconv.ParseUint(version, 2, len(version))
		versionSum += int(versionInt)

		i += 6
		if id == "100" {
			// literal value
			var valueString string
			parts := 1
			for input[i] == '1' {
				valueString += input[i+1 : i+5]
				i += 5
				parts++
			}
			valueString += input[i+1 : i+5]
			value, _ := strconv.ParseUint(valueString, 2, len(valueString))
			tokens = append(tokens, Token{
				OP:     "literal",
				Value:  int(value),
				Length: 6 + parts + len(valueString),
			})
			i += 5
		} else {
			lengthTypeID := input[i]
			i++
			if lengthTypeID == '0' {
				// know the subpackets length but not how many there are
				totalLength, _ := strconv.ParseUint(input[i:i+15], 2, 15)
				tokens = append(tokens, Token{
					OP:               id,
					SubpacketsLength: int(totalLength),
					Length:           22,
				})
				i += 15
			} else {
				// know how many subpackets there are but not their length
				subPackets, _ := strconv.ParseUint(input[i:i+11], 2, 11)
				tokens = append(tokens, Token{
					OP:                 id,
					RequiredSubpackets: int(subPackets),
					Length:             18,
				})
				i += 11
			}
		}
	}
	return tokens, versionSum
}

func eval(token []Token) int {
	cmpFunc := map[string]func(t1, t2 Token) bool{
		"=":   func(t1, t2 Token) bool { return t1.Value == t2.Value },
		">":   func(t1, t2 Token) bool { return t1.Value > t2.Value },
		"<":   func(t1, t2 Token) bool { return t1.Value < t2.Value },
		"max": func(t1, t2 Token) bool { return t1.Value > t2.Value },
		"min": func(t1, t2 Token) bool { return t1.Value < t2.Value },
	}

	stack := make(generic.Stack[Token], 0)
	for i := len(token) - 1; i >= 0; i-- {
		switch decodecOP := decode(token[i].OP); decodecOP {
		case "literal":
			stack.Push(token[i])
		case "=", ">", "<":
			updated := Token{}
			next1 := stack.Pop()
			next2 := stack.Pop()
			updated.Value = 0
			if cmpFunc[decodecOP](next1, next2) {
				updated.Value = 1
			}
			updated.Length = token[i].Length + next1.Length + next2.Length
			stack.Push(updated)
		case "min", "max":
			updated := Token{}
			updated.Length = token[i].Length
			next := stack.Pop()
			updated.Value = next.Value
			updated.Length += next.Length
			token[i].RequiredSubpackets--
			token[i].SubpacketsLength -= next.Length
			for generic.Max(token[i].RequiredSubpackets, token[i].SubpacketsLength) > 0 {
				next := stack.Pop()
				updated.Length += next.Length
				if cmpFunc[decodecOP](next, updated) {
					updated.Value = next.Value
				}
				token[i].RequiredSubpackets--
				token[i].SubpacketsLength -= next.Length
			}
			stack.Push(updated)
		case "+":
			updated := Token{}
			updated.Length = token[i].Length
			updated.Value = 0
			for generic.Max(token[i].RequiredSubpackets, token[i].SubpacketsLength) > 0 {
				next := stack.Pop()
				updated.Value += next.Value
				updated.Length += next.Length
				token[i].RequiredSubpackets--
				token[i].SubpacketsLength -= next.Length
			}
			stack.Push(updated)
		case "*":
			updated := Token{}
			updated.Value = 1
			updated.Length = token[i].Length
			for generic.Max(token[i].RequiredSubpackets, token[i].SubpacketsLength) > 0 {
				next := stack.Pop()
				updated.Value *= next.Value
				updated.Length += next.Length
				token[i].RequiredSubpackets--
				token[i].SubpacketsLength -= next.Length
			}
			stack.Push(updated)
		}
	}

	return stack.Pop().Value
}

func decode(code string) string {
	switch code {
	case "000":
		return "+"
	case "001":
		return "*"
	case "010":
		return "min"
	case "011":
		return "max"
	case "101":
		return ">"
	case "110":
		return "<"
	case "111":
		return "="
	default:
		return code
	}
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	decoded, err := hex.DecodeString(scanner.Text())
	if err != nil {
		return err
	}

	var input string
	for _, d := range decoded {
		input += fmt.Sprintf("%08b", d)
	}

	tokens, versionSum := tokenize(input)
	println("silver:", versionSum)
	println("gold:", eval(tokens))
	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
