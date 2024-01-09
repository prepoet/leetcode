package leetcode

import "sort"

func findMinArrowShots(points [][]int) int {
	length := len(points)
	if length <= 1 {
		return length
	}
	// 先将数组按照右边界线排序
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	num := 1
	before := []int{points[0][0], points[0][1]}
	for i := 1; i < length; i++ {
		if points[i][0] <= before[1] {
			// 相交 缩小区间
			if points[i][0] > before[0] {
				before[0] = points[i][0]
			}
		} else {
			// 不相交
			before[0] = points[i][0]
			before[1] = points[i][1]
			num++
		}
	}
	return num
}
