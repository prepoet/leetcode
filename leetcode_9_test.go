package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{1, true},
	}
	for _, test := range tests {
		if got := isPalindrome(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("isPalindrome(%v) = %v", test.input, got)
		}
	}
}
