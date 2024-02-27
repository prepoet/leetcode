package leetcode

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		len := len(q)
		max := math.MinInt32
		for i := 0; i < len; i++ {
			cur := q[0]
			if cur.Val > max {
				max = cur.Val
			}
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		arr = append(arr, max)
	}
	return arr
}
