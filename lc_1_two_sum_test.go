package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input  []int
		input2 int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, test := range tests {
		if got := twoSum(test.input, test.input2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v, %v) = %v", test.input, test.input2, got)
		}
	}
}
