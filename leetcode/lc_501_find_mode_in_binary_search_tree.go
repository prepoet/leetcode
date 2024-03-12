package leetcode

func findMode(root *TreeNode) []int {
	maxCount, count := 0, 0
	var prev *TreeNode
	var modes []int

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if prev != nil {
			if prev.Val == root.Val {
				count++
			} else {
				count = 1
			}
		}
		if count > maxCount {
			maxCount = count
			modes = []int{root.Val}
		} else if count == maxCount {
			modes = append(modes, root.Val)
		}
		prev = root
		inorder(root.Right)
	}
	inorder(root)
	return modes
}

func findMode1(root *TreeNode) []int {
	numCount := make(map[int]int)
	inOrderFindMod(root, numCount)
	maxCount := 0
	countNums := make(map[int][]int)
	for num, count := range numCount {
		if count > maxCount {
			maxCount = count
		}
		if arr, ok := countNums[count]; ok {
			arr = append(arr, num)
			countNums[count] = arr
		} else {
			countNums[count] = []int{num}
		}
	}
	return countNums[maxCount]
}

func inOrderFindMod(root *TreeNode, numCount map[int]int) {
	if root == nil {
		return
	}
	inOrderFindMod(root.Left, numCount)
	if count, ok := numCount[root.Val]; ok {
		count++
		numCount[root.Val] = count
	} else {
		numCount[root.Val] = 1
	}
	inOrderFindMod(root.Right, numCount)
}
