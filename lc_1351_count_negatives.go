// leetcode_1351
package leetcode

func countNegatives(grid [][]int) int {
	m := len(grid[0])
	n := len(grid)
	if m <= 0 || n <= 0 {
		return 0
	}
	num := 0
	x := m - 1
	for i := m - 1; i >= 0; i-- {
		if grid[0][i] < 0 {
			num++
			x = i
		}
	}
	if n <= 1 {
		return num
	}
	for j := 1; j < n; j++ {
		plus := true
		for k := x; k >= 0; k-- {
			if grid[j][k] >= 0 {
				if k == m-1 {
					plus = false
				} else {
					x = k + 1
				}
				break
			} else if k == 0 {
				x = 0
			}
		}
		if plus {
			num += m - x
		}
	}
	return num
}
