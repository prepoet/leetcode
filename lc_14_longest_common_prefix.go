package leetcode

import "strings"

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.
func longestCommonPrefix(strs []string) string {
	result := ""
	sl := len(strs)
	for _, c := range strs[0] {
		subStr := result + string(c)
		find := true
		for i := 1; i < sl; i++ {
			if !strings.HasPrefix(strs[i], subStr) {
				find = false
				break
			}
		}
		if find {
			result = subStr
		} else {
			break
		}
	}
	return result
}
