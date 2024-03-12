package leetcode

var isBST bool = true

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	isBST = true
	traversalCheckBST(root)
	return isBST
}

func traversalCheckBST(root *TreeNode) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	var leftMin, leftMax int
	if root.Left != nil {
		// 左子树最大值比root小
		leftMin, leftMax = traversalCheckBST(root.Left)
		if isBST && leftMax >= root.Val {
			isBST = false
		}
	} else {
		leftMin = root.Val
	}
	var rightMin, rightMax int
	if root.Right != nil {
		// 右子树最小值比root大
		rightMin, rightMax = traversalCheckBST(root.Right)
		if isBST && rightMin <= root.Val {
			isBST = false
		}
	} else {
		rightMax = root.Val
	}
	return leftMin, rightMax
}
