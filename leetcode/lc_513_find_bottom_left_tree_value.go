package leetcode

func findBottomLeftValue(root *TreeNode) int {
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		if len(newNodes) == 0 {
			// 最后一层
			return nodes[0].Val
		}
		nodes = newNodes
	}
	return 0
}
