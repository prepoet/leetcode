package leetcode

import "sort"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			// 维护 nums[0..slow] 无重复
			nums[slow] = nums[fast]
		}
	}
	// 数组长度为索引 + 1
	return slow + 1
}

// 解法2
func removeDuplicates1(nums []int) int {
	ln := 0
	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[i-1] == nums[i] {
			// max num is 100
			nums[i-1] = 101
		} else {
			ln++
		}
	}
	sort.Ints(nums)
	return ln + 1
}
