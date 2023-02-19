package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2
	// 第一个节点
	var first *ListNode
	// 前一个节点
	var before *ListNode
	for true {
		var small *ListNode
		if p1 == nil && p2 == nil {
			break
		} else if p1 != nil && p2 == nil {
			small = p1
			p1 = p1.Next
		} else if p2 != nil && p1 == nil {
			small = p2
			p2 = p2.Next
		} else {
			if p1.Val <= p2.Val {
				small = p1
				p1 = p1.Next
			} else {
				small = p2
				p2 = p2.Next
			}
		}
		if before != nil {
			before.Next = small
			before = small
		} else {
			first = small
			before = small
		}
	}
	return first
}
