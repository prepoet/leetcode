package leetcode

import "math"

// 暴力破解 超时
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt
	for _, coin := range coins {
		subResult := coinChange(coins, amount-coin)
		if subResult == -1 {
			// 不成立
			continue
		}
		res = int(math.Min(float64(res), float64(1+subResult)))
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

// 动态规划 增加备忘录 尝试分解
func coinChange1(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -2
	}
	return dpCoinChange1(coins, amount, memo)
}

func dpCoinChange1(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != -2 {
		return memo[amount]
	}
	res := math.MaxInt
	for _, coin := range coins {
		subResult := dpCoinChange1(coins, amount-coin, memo)
		if subResult == -1 {
			// 不成立
			continue
		}
		res = int(math.Min(float64(res), float64(1+subResult)))
	}
	if res == math.MaxInt {
		memo[amount] = -1
		return -1
	}
	memo[amount] = res
	return res
}

// 基于base case 从子问题递推
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 数组大小为 amount + 1，初始值也为 amount + 1
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < amount+1; i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			if i >= coin {
				// 金额>=硬币大小的才能继续 否则跳过
				dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 基于base case 从子问题递推
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 数组大小为 amount + 1，初始值也为 amount + 1
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < amount+1; i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			if i >= coin {
				// 金额>=硬币大小的才能继续 否则跳过
				dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
