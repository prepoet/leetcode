package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	len := len(lists)
	if len < 1 {
		return nil
	}
	if len == 1 {
		return lists[0]
	}
	num := len / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeKLists1(left, right)
}

func mergeKLists1(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeKLists1(left.Next, right)
		return left
	}
	right.Next = mergeKLists1(left, right.Next)
	return right
}
