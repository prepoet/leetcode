package leetcode

import "math"

func sumRootToLeaf(root *TreeNode) int {
	var sum int
	traversalSumRootToLeaf(root, []int{}, &sum)
	return sum
}

func traversalSumRootToLeaf(root *TreeNode, vals []int, sum *int) {
	if root == nil {
		return
	}
	vals = append(vals, root.Val)
	if root.Left == nil && root.Right == nil {
		// 叶子结点
		count := len(vals)
		for i := count - 1; i >= 0; i-- {
			*sum += vals[i] * int(math.Pow(2, float64(count-1-i)))
		}
		return
	}
	traversalSumRootToLeaf(root.Left, vals, sum)
	traversalSumRootToLeaf(root.Right, vals, sum)
}
