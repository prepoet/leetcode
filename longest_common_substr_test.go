package leetcode

import (
	"reflect"
	"testing"
)

func TestLongestCommonSubstr(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"reflower", "flow", "flight"}, "fl"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"d", "d", "d2"}, "d"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"cir", "car"}, "c"},
	}
	for _, test := range tests {
		if got := longestCommonSubstr(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("longestCommonSubstr(%v) = %v", test.input, got)
		}
	}
}
