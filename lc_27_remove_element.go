package main

func removeElement(nums []int, val int) int {
	same := 0
	l := len(nums)
	j := l - 1
	for i := 0; i < l; i++ {
		if nums[i] == val {
			same++
			for j > i {
				if nums[j] == val {
					same++
					nums[j] = 101
					j--
				} else {
					nums[i] = nums[j]
					nums[j] = 101
					j--
					break
				}
			}
		}
	}
	return l - same
}
