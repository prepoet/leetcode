package leetcode

// O(n) 时间复杂度
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	count := 0
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}
		count += len(nodes)
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		nodes = newNodes
	}
	return count
}
