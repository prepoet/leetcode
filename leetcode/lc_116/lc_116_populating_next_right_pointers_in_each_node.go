package lc116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// perfect binary tree // 层序遍历连接指针
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	nodes := []*Node{root}
	for len(nodes) > 0 {
		newNodes := []*Node{}
		for i, node := range nodes {
			if i+1 < len(nodes) {
				node.Next = nodes[i+1]
			}
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
				newNodes = append(newNodes, node.Right)
			}
		}
		nodes = newNodes
	}
	return root
}
