package leetcode

var eqTargetSum bool = false

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	eqTargetSum = false
	traversal112(root, 0, targetSum)
	return eqTargetSum
}

func traversal112(root *TreeNode, val int, targetSum int) {
	if root == nil {
		return
	}
	val += root.Val
	if root.Left == nil && root.Right == nil {
		if val == targetSum && !eqTargetSum {
			eqTargetSum = true
		}
		return
	}
	traversal112(root.Left, val, targetSum)
	traversal112(root.Right, val, targetSum)
}
