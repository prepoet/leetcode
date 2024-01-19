package leetcode

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	l13 := ListNode{3, nil}
	l12 := ListNode{2, &l13}
	l11 := ListNode{1, &l12}
	l1 := ListNode{1, &l11}
	r := deleteDuplicates(&l1)
	fmt.Printf("xxx: %v", r)
}
