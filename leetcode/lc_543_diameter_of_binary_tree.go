package leetcode

import "math"

var max543 int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	max543 = 0
	traversal543(root)
	return max543
}

func traversal543(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := traversal543(root.Left)
	right := traversal543(root.Right)
	sum := left + right
	if sum > max543 {
		max543 = sum
	}
	return 1 + int(math.Max(float64(left), float64(right)))
}
