package leetcode

var arr113 [][]int = [][]int{}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	arr113 = [][]int{}
	traversal113(root, 0, []int{}, targetSum)
	return arr113
}

func traversal113(root *TreeNode, val int, arr []int, targetSum int) {
	if root == nil {
		return
	}
	val += root.Val
	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil {
		if val == targetSum {
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			arr113 = append(arr113, newArr)
		}
		return
	}
	traversal113(root.Left, val, arr, targetSum)
	traversal113(root.Right, val, arr, targetSum)
}
