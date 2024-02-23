package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]int{3, 2, 1, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("soted: %v", tt.args.array)
			quickSort1(tt.args.array)
			fmt.Printf("soted: %v", tt.args.array)
		})
	}
}
