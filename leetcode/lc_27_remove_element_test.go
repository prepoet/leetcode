package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		input1 []int
		input2 int
		want   int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},//1
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5}, //4
		{[]int{1}, 1, 0},
		{[]int{4, 5}, 5, 1},
		{[]int{3, 3}, 3, 0},
	}
	for _, test := range tests {
		// if got := removeElement(test.input1, test.input2); !reflect.DeepEqual(got, test.want) {
		// 	t.Errorf("removeElement(%v, %v) = %v", test.input1, test.input2, got)
		// }
		// if got := removeElement1(test.input1, test.input2); !reflect.DeepEqual(got, test.want) {
		// 	t.Errorf("removeElement1(%v, %v) = %v", test.input1, test.input2, got)
		// }
		if got := removeElement2(test.input1, test.input2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("removeElement2(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}
