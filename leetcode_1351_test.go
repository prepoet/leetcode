package main

import (
	"reflect"
	"testing"
)

func TestCountNegatives(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  int
	}{
		{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
		{[][]int{{3, 2}, {1, 0}}, 0},
		{[][]int{{-1}}, 1},
	}
	for _, test := range tests {
		if got := countNegatives(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("countNegatives(%v) = %v", test.input, got)
		}
	}
}
