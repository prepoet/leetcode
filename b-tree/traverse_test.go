package btree

import (
	"leetcode"
	"reflect"
	"testing"
)

var n2 *leetcode.TreeNode = &leetcode.TreeNode{Val: 2, Left: nil, Right: nil}
var n3 *leetcode.TreeNode = &leetcode.TreeNode{Val: 3, Left: nil, Right: nil}
var root *leetcode.TreeNode = &leetcode.TreeNode{Val: 1, Left: n2, Right: n3}

func Test_preTraverse(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{root}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preTraverse(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Logf("got: %v\n", got)
				t.Errorf("preTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inTraverse(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{root}, []int{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inTraverse(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Logf("%v\n", got)
				t.Errorf("inTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postTraverse(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{root}, []int{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postTraverse(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Logf("%v\n", got)
				t.Errorf("postTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
