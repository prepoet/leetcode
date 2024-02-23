package leetcode

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
