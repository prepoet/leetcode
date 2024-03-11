package leetcode

var sum404 int = 0

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum404 = 0
	traversal404(root, false)
	return sum404
}

func traversal404(root *TreeNode, isLeft bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if isLeft {
			sum404 += root.Val
		}
		return
	}
	traversal404(root.Left, true)
	traversal404(root.Right, false)
}
