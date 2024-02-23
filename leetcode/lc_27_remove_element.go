package leetcode

// 双指针
func removeElement(nums []int, val int) int {
	len := len(nums)
	num, left, right := 0, 0, len-1
	for left <= right {
		if nums[left] == val {
			if nums[right] != val {
				// 交换
				nums[left] = nums[right]
				nums[right] = val
				left++
			}
			right--
			num++
		} else {
			left++
		}
	}
	return len - num
}

func removeElement1(nums []int, val int) int {
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

// 快慢指针
func removeElement2(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
