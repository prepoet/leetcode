package leetcode

func permute(nums []int) [][]int {
	result := [][]int{}
	bpPermute(nums, []int{}, &result)
	return result
}

func bpPermute(nums []int, result []int, results *[][]int) {
	if len(nums) == 1 {
		result = append(result, nums[0])
		*results = append(*results, result)
		return
	}
	for i, num := range nums {
		if contains(result, num) {
			// 从不包含的里边挑选一个
			continue
		}
		newResult := append(result, num)

		remain := []int{}
		remain = append(remain, nums[:i]...)
		remain = append(remain, nums[i+1:]...)
		bpPermute(remain, newResult, results)
	}
}

func contains(nums []int, num int) bool {
	for _, x := range nums {
		if x == num {
			// 不能重复
			return true
		}
	}
	return false
}

func permute1(nums []int) [][]int {
	result := [][]int{}
	used := make([]bool, len(nums))
	bpPermute1(nums, []int{}, used, &result)
	return result
}

func bpPermute1(nums []int, track []int, used []bool, result *[][]int) {
	if len(track) == len(nums) {
		*result = append(*result, track)
		return
	}
	for i, num := range nums {
		if used[i] {
			continue
		}
		newResult := append(track, num)
		used[i] = true
		bpPermute1(nums, newResult, used, result)
		used[i] = false
	}
}
