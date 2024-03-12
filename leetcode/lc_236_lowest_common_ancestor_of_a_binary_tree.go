package leetcode

var pa, qa []*TreeNode
var result *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pa, qa = []*TreeNode{}, []*TreeNode{}
	result = nil
	preorder(root, p, q)
	return result
}

func preorder(root, p, q *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	arr := []*TreeNode{root}
	if p == root {
		pa = arr
	}
	if p == root {
		qa = arr
	}
	pcount, qcount := len(pa), len(qa)
	if pcount > 0 && qcount > 0 && result == nil {
		count := pcount
		if pcount > qcount {
			count = qcount
		}
		for i := 0; i < count; i++ {
			if pa[i] != qa[i] {
				result = pa[i-1]
				break
			}
		}
		if result == nil {
			result = pa[count-1]
		}
	}
	arr = append(arr, preorder(root.Left, p, q)...)
	arr = append(arr, preorder(root.Right, p, q)...)
	return arr
}