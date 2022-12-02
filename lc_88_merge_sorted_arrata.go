package main

import "fmt"

//nums1.length == m + n
//nums2.length == n
//O(m + n)

// [1,2,9,10,0,0,0,0] [2,3,4,11]
func merge(nums1 []int, m int, nums2 []int, n int) {
	var p1, p2 int
	for p1 < m || p2 < n {
		n1 := nums1[p1]
		n2 := nums2[p2]
		if n1 > n2 {
			start := p2
			for p2+1 < n && nums2[p2+1] < n1 {
				p2++
			}
			end := p2 + 1
			// num2:[start:end]插到左边数组
			skip := end - start
			fmt.Printf("start: %d, end: %d\n", start, end)
			var tmp = p1
			for tmp < m {
				nums1[tmp+skip] = nums1[tmp]
				tmp++
			}
			for i := start; i < end; i++ {
				nums1[p1+i] = nums2[i]
			}
			p1 += skip
			m += skip
			p2 += skip
		} else {
			p1++
		}
		fmt.Printf("nums1: %v, p1: %d, p2: %d, m: %d, n: %d\n", nums1, p1, p2, m, n)
	}
}
