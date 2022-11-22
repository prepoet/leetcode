package main

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"reflower", "flow", "flight"}, ""},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"d", "d", "d2"}, "d"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"cir", "car"}, "c"},
	}
	for _, test := range tests {
		if got := longestCommonPrefix(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("longestCommonPrefix(%v) = %v", test.input, got)
		}
	}
}
