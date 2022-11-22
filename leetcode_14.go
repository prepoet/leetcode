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
	result := ""
	for _, c := range strs[minLStrIndex] {
		subStr := result + string(c)
		find := true
		for j, str := range strs {
			if j == minLStrIndex {
				continue
			}
			if !strings.HasPrefix(str, subStr) {
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
