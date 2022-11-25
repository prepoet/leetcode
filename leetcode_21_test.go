package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		input1 *ListNode
		input2 *ListNode
		want   *ListNode
	}{
		{},
	}
	for _, test := range tests {
		if got := mergeTwoLists(test.input1, test.input2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("mergeTwoLists(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}
