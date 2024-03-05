package leetcode

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	arr = append(arr, left...)
	arr = append(arr, right...)
	arr = append(arr, root.Val)
	return arr
}

var postorderArr []int = []int{}

func postorderTraversal1(root *TreeNode) []int {
	postorderArr = []int{}
	postorder(root)
	return postorderArr
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	postorderArr = append(postorderArr, root.Val)
}
