package leetcode

var paths []string = []string{}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths = []string{}
	traversal257(root, []string{})
	return paths
}

func traversal257(root *TreeNode, path []string) {
	if root == nil {
		return
	}
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		paths = append(paths, strings.Join(path, "->"))
		return
	}
	traversal257(root.Left, path)
	traversal257(root.Right, path)
}