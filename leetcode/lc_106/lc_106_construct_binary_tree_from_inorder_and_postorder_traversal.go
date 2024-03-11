package lc106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	count := len(postorder)
	root := &TreeNode{Val: postorder[count-1]}
	build(root, inorder, postorder)
	return root
}

func build(root *TreeNode, inorder []int, postorder []int) {
	count := len(inorder)
	var rootIdx int
	for i := 0; i < count; i++ {
		if inorder[i] == root.Val {
			rootIdx = i
			break
		}
	}
	if rootIdx > 0 {
		// 左子树有节点
		inLeft := inorder[0:rootIdx]
		rightLen := count - 1 - len(inLeft)
		postLeft := postorder[0 : count-1-rightLen]
		root.Left = &TreeNode{Val: postLeft[len(postLeft)-1]}
		build(root.Left, inLeft, postLeft)
	}
	if count-1 != rootIdx {
		// 右子树还有节点
		inRight := inorder[rootIdx+1 : count]
		leftLen := count - 1 - len(inRight)
		postRight := postorder[leftLen : count-1]
		root.Right = &TreeNode{Val: postRight[len(postRight)-1]}
		build(root.Right, inRight, postRight)
	}
}
