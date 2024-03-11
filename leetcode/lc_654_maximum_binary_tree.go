package leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	max, maxIdx := findMax(nums)
	root := &TreeNode{Val: max}
	buildBinaryTree(root, nums[0:maxIdx], nums[maxIdx+1:len(nums)])
	return root
}

func findMax(nums []int) (int, int) {
	max := nums[0]
	maxIdx := 0
	for i, count := 1, len(nums); i < count; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}
	return max, maxIdx
}

func buildBinaryTree(root *TreeNode, left []int, right []int) {
	if len(left) > 0 {
		max, maxIdx := findMax(left)
		root.Left = &TreeNode{Val: max}
		buildBinaryTree(root.Left, left[0:maxIdx], left[maxIdx+1:len(left)])
	}
	if len(right) > 0 {
		max, maxIdx := findMax(right)
		root.Right = &TreeNode{Val: max}
		buildBinaryTree(root.Right, right[0:maxIdx], right[maxIdx+1:len(right)])
	}
}
