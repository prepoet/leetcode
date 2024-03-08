package lc559

type Node struct {
	Val      int
	Children []*Node
}

var depth int = 0

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth = 0
	traversal(root, 1)
	return depth
}

func traversal(root *Node, level int) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		traversal(child, level+1)
	}
	if level > depth {
		depth = level
	}
}
