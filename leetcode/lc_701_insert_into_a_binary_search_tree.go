package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	traversalInsertIntoBST(root, val, nil)
	return root
}

func traversalInsertIntoBST(root *TreeNode, val int, pre *TreeNode) {
	if root == nil {
		if val < pre.Val {
			pre.Left = &TreeNode{Val: val}
		} else {
			pre.Right = &TreeNode{Val: val}
		}
		return
	}
	if val < root.Val {
		traversalInsertIntoBST(root.Left, val, root)
	} else if val > root.Val {
		traversalInsertIntoBST(root.Right, val, root)
	}
}
