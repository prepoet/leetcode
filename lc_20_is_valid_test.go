package leetcode

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"{[]}", true},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}
	for _, test := range tests {
		if got := isValid(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("isValid(%v) = %v", test.input, got)
		}
	}
}
