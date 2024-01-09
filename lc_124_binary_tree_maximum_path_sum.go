package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxPathSum(root *TreeNode) int {
	// 这里每次刷新结果跑 因为leetcode的上一个测试用例的结果会带到下一个中
	res = math.MinInt
	oneSlideMax(root)
	return res
}

func oneSlideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 这里为什么要过滤负数 负数一定是拖累的 可以不带负数子节点 保持总值最大
	left := math.Max(0, float64(oneSlideMax(root.Left)))
	right := math.Max(0, float64(oneSlideMax(root.Right)))
	res = int(math.Max(float64(res), left+right+float64(root.Val)))
	// 只选择其中一条路径 左或右 返回该路径的最大值
	return int(math.Max(float64(left), float64(right)) + float64(root.Val))
}
