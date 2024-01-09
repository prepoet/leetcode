package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res1 int
var rank int

func kthSmallest(root *TreeNode, k int) int {
	res1 = 0
	rank = 0
	traverse(root, k)
	return res1
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	// 本质是中序遍历，因为输入是二叉搜索树，父节点在中间，左子节点比父节点小，右子节点比父节点大，所以二叉搜索树的中序遍历本质是有序的
	rank++
	if k == rank {
		res1 = root.Val
		return
	}
	traverse(root.Right, k)
}
