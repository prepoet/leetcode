package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}


var pa, qa []*TreeNode
var result *TreeNode

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	pa, qa = []*TreeNode{}, []*TreeNode{}
	result = nil
	preorder(root, p, q, []*TreeNode{})
	return result
}

func preorder(root, p, q *TreeNode, arr []*TreeNode) {
	if root == nil {
		return
	}
	arr = append(arr, root)
	if p == root {
		pa = copySlice(arr)
	}
	if q == root {
		qa = copySlice(arr)
	}
	pcount, qcount := len(pa), len(qa)
	if pcount > 0 && qcount > 0 && result == nil {
		count := pcount
		if pcount > qcount {
			count = qcount
		}
		for i := 0; i < count; i++ {
			if pa[i] != qa[i] {
				if i >= 1 {
					result = pa[i-1]
				}
				break
			}
		}
		if result == nil {
			if count == pcount {
				result = pa[count-1]
			} else {
				result = qa[count-1]
			}
		}
	}
	preorder(root.Left, p, q, arr)
	preorder(root.Right, p, q, arr)
}

func copySlice(arr []*TreeNode) []*TreeNode {
	newArr := make([]*TreeNode, len(arr))
	copy(newArr, arr)
	return newArr
}
