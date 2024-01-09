package leetcode

import (
	"sort"
)

// 贪心策略 区间尽可能小 能保证移除最少的区间
func eraseOverlapIntervals(intervals [][]int) int {
	length := len(intervals)
	if length <= 0 {
		return 0
	}
	// 先将数组按照右边界线排序
	// 自定义比较器的数组排序函数
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	removed := 0
	before := intervals[0][1]
	for i := 1; i < length; i++ {
		if intervals[i][0] < before {
			removed++
		} else {
			before = intervals[i][1]
		}
	}
	return removed
}
