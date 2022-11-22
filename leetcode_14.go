package main

import (
	"strings"
)

//Write a function to find the longest common prefix string amongst an array of strings.
//If there is no common prefix, return an empty string "".
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] consists of only lowercase English letters.
func longestCommonPrefix(strs []string) string {
	minL := len(strs[0])
	var minLStrIndex int
	for i, str := range strs {
		l := len(str)
		if len(str) < minL {
			minLStrIndex = i
			minL = l
		}
	}
	minLStr := strs[minLStrIndex]
	result := ""
	tmp := ""
	for i, c := range minLStr {
		subStr := tmp + string(c)
		find := true
		for j, str := range strs {
			if j == minLStrIndex {
				continue
			}
			// 这里替换为 strings.Contains 方法 且 find = false 的break去掉 就为 最长公共子串
			if !strings.HasPrefix(str, subStr) {
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
			break
		}
	}
	return result
}
