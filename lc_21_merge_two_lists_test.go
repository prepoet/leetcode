package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	node11 := ListNode{1, nil}
	node12 := ListNode{2, nil}
	node14 := ListNode{4, nil}
	node11.Next = &node12
	node12.Next = &node14

	node21 := ListNode{1, nil}
	node23 := ListNode{3, nil}
	node24 := ListNode{4, nil}
	node21.Next = &node23
	node23.Next = &node24
	node3 := mergeTwoLists(&node11, &node21)
	fmt.Printf("%v", node3)
}
