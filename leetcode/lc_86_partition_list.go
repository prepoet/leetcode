package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// 存放小于 x 的链表的虚拟头结点
	dummy1 := &ListNode{-1, nil}
	// 存放大于等于 x 的链表的虚拟头结点
	dummy2 := &ListNode{-1, nil}
	// p1, p2 指针负责生成结果链表
	p1, p2 := dummy1, dummy2
	// p 负责遍历原链表，类似合并两个有序链表的逻辑
	// 这里是将一个链表分解成两个链表
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		// 不能直接让 p 指针前进
		nextNode := p.Next
		// 断开原链表中的每个节点的 Next 指针
		p.Next = nil
		p = nextNode
	}
	// 连接两个链表
	p1.Next = dummy2.Next
	return dummy1.Next
}
