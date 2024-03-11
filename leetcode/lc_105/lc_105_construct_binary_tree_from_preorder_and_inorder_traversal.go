package lc105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	build(root, preorder, inorder)
	return root
}

func build(root *TreeNode, preorder []int, inorder []int) {
	count := len(inorder)
	rootIdx := 0
	for i := 0; i < count; i++ {
		if root.Val == inorder[i] {
			rootIdx = i
		}
	}
	if rootIdx > 0 {
		// 有左子树
		inLeft := inorder[0:rootIdx]
		preLeft := preorder[1 : 1+len(inLeft)]
		root.Left = &TreeNode{Val: preLeft[0]}
		build(root.Left, preLeft, inLeft)
	}
	if rootIdx != count-1 {
		// 有右子树
		inRight := inorder[rootIdx+1 : count]
		leftLen := count - 1 - len(inRight)
		preRight := preorder[1+leftLen : count]
		root.Right = &TreeNode{Val: preRight[0]}
		build(root.Right, preRight, inRight)
	}
}
