package main

/**
O(log n) 二分查找
*/
func searchInsert(nums []int, target int) int {
	l := len(nums)
	left := 0
	right := l - 1
	for {
		if right-left <= 1 {
			// 相邻
			if target <= nums[left] {
				return left
			} else if target <= nums[right] {
				return right
			} else {
				return right + 1
			}
		}
		i := (left + right) / 2
		num := nums[i]
		if target < num {
			right = max(left+1, i-1)
		} else if target > num {
			left = min(right-1, i+1)
		} else {
			return i
		}
	}
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
