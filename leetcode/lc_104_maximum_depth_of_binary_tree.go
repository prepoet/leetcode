package leetcode

import "math"

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1 := maxDepth(root.Left)
	d2 := maxDepth(root.Right)
	return int(math.Max(float64(d1), float64(d2))) + 1
}

// 记录深度
var depth int = 0
// 最大深度维护在外部变量
var maxDepthRes int = 0

func maxDepth1(root *TreeNode) int {
	depth = 0
    maxDepthRes = 0
	traverse104(root)
	return maxDepthRes
}

func traverse104(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		if depth > maxDepthRes {
			maxDepthRes = depth
		}
	}
	traverse104(root.Left)
	traverse104(root.Right)
	depth--
}
