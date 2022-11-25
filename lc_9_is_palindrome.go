package main

import (
	"math"
)

/*
Given an integer x, return true if x is a palindrome, and false otherwise.
Follow up: Could you solve it without converting the integer to a string?
*/
func isPalindrome(x int) bool {
	if x < 0 {
		// 负数均不满足
		return false
	}
	digit := int(math.Log10(float64(x))) + 1
	if digit == 1 {
		// 个位数均满足
		return true
	}
	pairNum := digit / 2
	for i := 1; i <= pairNum; i++ {
		left := x / int(math.Pow10(digit-i)) % 10
		right := x / int(math.Pow10(i-1)) % 10
		if left != right {
			return false
		}
	}
	return true
}
