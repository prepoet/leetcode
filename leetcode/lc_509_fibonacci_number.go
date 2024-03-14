package leetcode

// 暴力解
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

// 递归 + 备忘录
func fib1(n int) int {
	memo := make([]int, n+1)
	return fib1Helper(n, memo)
}

func fib1Helper(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] > 0 {
		return memo[n]
	}
	result := fib1Helper(n-2, memo) + fib1Helper(n-1, memo)
	memo[n] = result
	return result
}

// dp table + 循环
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	dpTable := make([]int, n+1)
	dpTable[0], dpTable[1] = 0, 1
	for i := 2; i <= n; i++ {
		dpTable[i] = dpTable[i-1] + dpTable[i-2]
	}
	return dpTable[n]
}

// 优化为size=2的数组
func fib3(n int) int {
	if n == 0 {
		return n
	}
	dpTable := []int{0, 1}
	var result int
	for i := 2; i < n; i++ {
		result = dpTable[1] + dpTable[0]
		dpTable[0] = dpTable[1]
		dpTable[1] = result
	}
	return dpTable[1]
}

// 优化为两个变量 不需要保持连续
func fib4(n int) int {
	if n <= 1 {
		return n
	}
	dp0, dp1 := 0, 1
	var result int
	for i := 2; i <= n; i++ {
		result = dp1 + dp0
		dp0 = dp1
		dp1 = result
	}
	return dp1
}
