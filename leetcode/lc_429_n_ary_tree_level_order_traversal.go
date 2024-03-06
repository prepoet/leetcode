package leetcode

func levelOrderNaryTree(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	nodes := []*Node{root}
	for len(nodes) > 0 {
		newNodes := []*Node{}
		nums := []int{}
		for _, node := range nodes {
            nums = append(nums, node.Val)
			for _, child := range node.Children {
				if child != nil {
					newNodes = append(newNodes, child)
				}
			}
		}
		result = append(result, nums)
		nodes = newNodes
	}
	return result
}
