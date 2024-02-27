package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummay := &ListNode{0, nil}
	dummay.Next = head
	fast, slow := dummay, dummay
	for i := 1; i <= n; i++ {
		// 快指针先走n步
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	// 快指针慢指针同时走，直至快指针走完，降同时走len-n个节点，慢指针走到了倒数第n个节点前
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummay.Next
}
