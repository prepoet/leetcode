package leetcode

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	arr = append(arr, root.Val)
	arr = append(arr, preorderTraversal(root.Left)...)
	arr = append(arr, preorderTraversal(root.Right)...)
	return arr
}

var array []int = []int{}

func preorderTraversal1(root *TreeNode) []int {
	array = []int{}
	traversal(root)
	return array
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	array = append(array, root.Val)
	traversal(root.Left)
	traversal(root.Right)
}
