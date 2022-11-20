package main

import "testing"

func TestCountNegatives(t *testing.T) {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	num := countNegatives(grid)
	t.Logf("num1: %d", num)
	grid2 := [][]int{{3, 2}, {1, 0}}
	num2 := countNegatives(grid2)
	t.Logf("num2: %d", num2)
	grid3 := [][]int{{-1}}
	num3 := countNegatives(grid3)
	t.Logf("num2: %d", num3)
}
