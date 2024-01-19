package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	p := head
	for p != nil {
		if p.Next != nil && p.Next.Val == p.Val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return dummy.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	hp := head
	p := head
	if head == nil {
		return head
	}
	lastVal := p.Val
	// 后值与前值比较
	for p.Next != nil {
		if p.Next.Val == lastVal {
			p.Next = p.Next.Next
		} else {
			lastVal = p.Next.Val
			p = p.Next
		}
	}
	return hp
}
