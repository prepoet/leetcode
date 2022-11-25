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
		{
			// 暂时没想到测试用例的快速构建 使用json反序列化是否可行
		},
	}
	for _, test := range tests {
		if got := mergeTwoLists(test.input1, test.input2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("mergeTwoLists(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}
