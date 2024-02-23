package leetcode

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		num := digits[i]
		if num+1 == 10 {
			digits[i] = 0
			if i == 0 {
				//生成新数组 这种情况是 每一位都是9
				var array = make([]int, l+1)
				array[0] = 1
				return array
			}
		} else {
			digits[i] = num + 1
			return digits
		}
	}
	return digits
}
