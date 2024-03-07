package leetcode

// 纯遍历
func isSymmetric(root *TreeNode) bool {
	nodes := []*TreeNode{root.Left, root.Right}
	size := len(nodes)
	for size > 0 {
		newNodes := []*TreeNode{}
		notNilNum := 0
		for i, node := range nodes {
			if (nodes[i] == nil && nodes[size-1-i] == nil) || (nodes[i] != nil && nodes[size-1-i] != nil && nodes[i].Val == nodes[size-1-i].Val) {
				if node != nil {
					notNilNum++
					newNodes = append(newNodes, node.Left)
					newNodes = append(newNodes, node.Right)
				} else {
					newNodes = append(newNodes, nil)
					newNodes = append(newNodes, nil)
				}
			} else {
				return false
			}
		}
		if notNilNum == 0 {
			break
		}
		nodes = newNodes
		size = len(newNodes)
	}
	return true
}

// 链表的思路
func isSymmetric1(root *TreeNode) bool {
	// 左右子树中序遍历
	left := inorder101(root.Left, true)
	right := inorder101(root.Right, false)
	if len(left) != len(right) {
		return false
	}
	for i, node := range left {
		if (node == nil && right[i] == nil) || (node != nil && right[i] != nil && node.Val == right[i].Val) {
			continue
		} else {
			return false
		}
	}
	return true
}

func inorder101(root *TreeNode, revert bool) []*TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		// 叶子结点
		return []*TreeNode{root}
	}
	arr := []*TreeNode{}
	if !revert {
		arr = append(arr, inorder101(root.Left, revert)...)
		arr = append(arr, root)
		arr = append(arr, inorder101(root.Right, revert)...)
	} else {
		arr = append(arr, inorder101(root.Left, revert)...)
		arr = append(arr, root)
		arr = append(arr, inorder101(root.Right, revert)...)
	}
	return arr
}
