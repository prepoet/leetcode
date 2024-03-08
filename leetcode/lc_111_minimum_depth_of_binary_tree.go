package leetcode

import "math"

var depth111 int = math.MaxInt32

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth111 = math.MaxInt32
	traversal111(root, 1)
	return depth111
}

func traversal111(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if level < depth111 {
			depth111 = level
		}
		return
	}
	level++
	traversal111(root.Left, level)
	traversal111(root.Right, level)
}
