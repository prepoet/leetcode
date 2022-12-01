package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	hp := head
	p := head
	if head == nil {
		return head
	}
	lastVal := p.Val
	for p.Next != nil {
		next := p.Next
		if next.Val == lastVal {
			p.Next = next.Next
		} else {
			lastVal = next.Val
			p = p.Next
		}
	}
	return hp
}
