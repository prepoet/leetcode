package leetcode

func lengthOfLastWord(s string) int {
	l := len(s)
	num := 0
	for i := l - 1; i >= 0; i-- {
		char := s[i]
		if char != ' ' {
			num++
		} else if num > 0 {
			break
		}
	}
	return num
}
