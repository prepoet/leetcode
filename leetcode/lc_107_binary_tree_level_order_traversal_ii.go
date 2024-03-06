package leetcode

var result107 [][]int = [][]int{}

func levelOrderBottom(root *TreeNode) [][]int {
	result107 = [][]int{}
	if root == nil {
		return result107
	}
	nodes := []*TreeNode{root}
	traversal107(nodes)
	return result107
}

func traversal107(nodes []*TreeNode) {
	if len(nodes) <= 0 {
		return
	}
	newNodes := []*TreeNode{}
	nums := []int{}
	for _, node := range nodes {
		if node.Left != nil {
			newNodes = append(newNodes, node.Left)
		}
		if node.Right != nil {
			newNodes = append(newNodes, node.Right)
		}
		nums = append(nums, node.Val)
	}
	traversal107(newNodes)
	result107 = append(result107, nums)
}
