package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	arr := [][]int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}
		nums := []int{}
		for _, node := range nodes {
			nums = append(nums, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		arr = append(arr, nums)
		nodes = newNodes
	}
	return arr
}
