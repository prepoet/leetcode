package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	cookies := 0
	for child < len(g) && cookies < len(s) {
		if g[child] <= s[cookies] {
			child++
		}
		cookies++
	}
	return child
}
