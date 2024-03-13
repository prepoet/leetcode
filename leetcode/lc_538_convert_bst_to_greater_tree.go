package leetcode

func convertBST(root *TreeNode) *TreeNode {
    sum := 0
	traversalConvertBST(root, &sum)
	return root
}

func traversalConvertBST(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traversalConvertBST(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traversalConvertBST(root.Left, sum)
}