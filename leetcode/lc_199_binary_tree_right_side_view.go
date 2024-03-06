package leetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		num := 0
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
			num = node.Val
		}
		nodes = newNodes
		arr = append(arr, num)
	}
	return arr
}
