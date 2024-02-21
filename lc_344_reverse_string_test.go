package leetcode

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]byte{'h', 'e', 'l', 'l', 'o'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.s)
			reverseString(tt.args.s)
			fmt.Println(tt.args.s)
		})
	}
}
