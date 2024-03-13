package leetcode

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := isUnivalTree(root.Left)
	right := isUnivalTree(root.Right)
	return left && right && isUnival(root, root.Left) && isUnival(root, root.Right)
}

func isUnival(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	return a.Val == b.Val
}
