package other

import (
	"strings"
)

func longestCommonSubstr(strs []string) string {
	minL := len(strs[0])
	var minLStrIndex int
	for i, str := range strs {
		l := len(str)
		if len(str) < minL {
			minLStrIndex = i
			minL = l
		}
	}
	result := ""
	tmp := ""
	for i, c := range strs[minLStrIndex] {
		subStr := tmp + string(c)
		find := true
		for j, str := range strs {
			if j == minLStrIndex {
				continue
			}
			if !strings.Contains(str, subStr) {
				find = false
				break
			}
		}
		if find {
			tmp = subStr
			if i == minL-1 {
				// 最后一个字符
				if len(tmp) > len(result) {
					result = tmp
				}
			}
		} else {
			if len(tmp) > len(result) {
				result = tmp
			}
		}
	}
	return result
}
