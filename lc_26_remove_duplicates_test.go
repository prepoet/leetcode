package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, test := range tests {
		if got := removeDuplicates(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("removeDuplicates(%v) = %v", test.input, got)
		}
	}
}
