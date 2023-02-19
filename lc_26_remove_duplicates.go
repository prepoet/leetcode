package leetcode

import "sort"

func removeDuplicates(nums []int) int {
	ln := 0
	l := len(nums)
	for i := 1; i < l; i++ {
		before := nums[i-1]
		cur := nums[i]
		if before == cur {
			// max num is 100
			nums[i-1] = 101
		} else {
			ln++
		}
	}
	sort.Ints(nums)
	return ln + 1
}
