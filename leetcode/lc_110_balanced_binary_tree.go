package leetcode

import "math"

var maxDiff int = 0

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	maxDiff = 0
	getDepth(root)
	return maxDiff <= 1
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	diff := int(math.Abs(float64(left) - float64(right)))
	if diff > maxDiff {
		maxDiff = diff
	}
	return int(math.Max(float64(left), float64(right))) + 1
}
