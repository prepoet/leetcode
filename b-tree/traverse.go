package btree

import "leetcode"

func preTraverse(root *leetcode.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// root.Val
	array := []int{}
	array = append(array, root.Val)
	array = append(array, preTraverse(root.Left)...)
	array = append(array, preTraverse(root.Right)...)
	return array
}

func inTraverse(root *leetcode.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// root.Val
	array := []int{}
	array = append(array, inTraverse(root.Left)...)
	array = append(array, root.Val)
	array = append(array, inTraverse(root.Right)...)
	return array
}


func postTraverse(root *leetcode.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// root.Val
	array := []int{}
	array = append(array, postTraverse(root.Left)...)
	array = append(array, postTraverse(root.Right)...)
	array = append(array, root.Val)
	return array
}
