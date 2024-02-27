package tree

import "algorithms/leetcode"

func count(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := count(root.Left)
	rightCount := count(root.Right)
	return leftCount + rightCount + 1
}
