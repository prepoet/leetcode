package leetcode

var inorderArr []int = []int{}

func inorderTraversal(root *TreeNode) []int {
	inorderArr = []int{}
	inorder(root)
	return inorderArr
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	inorderArr = append(inorderArr, root.Val)
	inorder(root.Right)
}
