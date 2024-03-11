package leetcode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	val := 0
	var left1, left2, right1, right2 *TreeNode
	if root1 != nil {
		val += root1.Val
		left1 = root1.Left
		right1 = root1.Right
	}
	if root2 != nil {
		val += root2.Val
		left2 = root2.Left
		right2 = root2.Right
	}
	// 重叠相加
	node := &TreeNode{Val: val}
	node.Left = mergeTrees(left1, left2)
	node.Right = mergeTrees(right1, right2)
	return node
}
