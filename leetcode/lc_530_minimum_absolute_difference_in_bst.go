package leetcode

import "math"

var minDifference int

func getMinimumDifference(root *TreeNode) int {
	minDifference = math.MaxInt
	inorderGetMinimumDifference(root, nil)
	return minDifference
}

func inorderGetMinimumDifference(root *TreeNode, last *TreeNode) {
	if root == nil {
		return
	}
	inorderGetMinimumDifference(root.Left, last)
	if last != nil {
		diff := root.Val - last.Val
		if diff < minDifference {
			minDifference = diff
		}
	}
	last = root
	inorderGetMinimumDifference(root.Right, last)
}
