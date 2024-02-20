package leetcode

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	len := len(nums)
	for fast < len {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow < len {
		nums[slow] = 0
		slow++
	}
}
