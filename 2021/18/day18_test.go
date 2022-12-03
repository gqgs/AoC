package main

import (
	"strings"
	"testing"
)

func Test_reduce(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"[[[[[9,8],1],2],3],4]",
			"[[[[0,9],2],3],4]",
		},
		{
			"[7,[6,[5,[4,[3,2]]]]]",
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"[[6,[5,[4,[3,2]]]],1]",
			"[[6,[5,[7,0]]],3]",
		},
		{
			"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			"[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]",
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			number := decodeReader(strings.NewReader(tt.input))
			if got := reduce(number).String(); got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
		})
	}

}
