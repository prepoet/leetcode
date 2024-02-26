package leetcode

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		// 抛弃该节点的左子树
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		// 抛弃该节点的右子树
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
