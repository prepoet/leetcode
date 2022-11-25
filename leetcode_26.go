package main

import "sort"

const max = 101

func removeDuplicates(nums []int) int {
	ln := 0
	l := len(nums)
	for i := 1; i < l; i++ {
		before := nums[i-1]
		cur := nums[i]
		if before == cur {
			nums[i-1] = max
		} else {
			ln++
		}
	}
	sort.Ints(nums)
	return ln + 1
}
