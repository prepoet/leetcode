package leetcode

//迭代
func searchBST(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if node.Val == val {
			return node
		}
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}

// 递归
func searchBST1(root *TreeNode, val int) *TreeNode {
	return traversalInSearchBst(root, val)
}

func traversalInSearchBst(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return traversalInSearchBst(root.Left, val)
	} else {
		return traversalInSearchBst(root.Right, val)
	}
}
