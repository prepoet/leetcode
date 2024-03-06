package leetcode

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	result := []float64{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		sum := 0
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
			sum += node.Val
		}
		result = append(result, float64(sum)/float64(len(nodes)))
		nodes = newNodes
	}
	return result
}
