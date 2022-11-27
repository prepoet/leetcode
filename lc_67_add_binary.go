package main

import (
	"strconv"
	"strings"
)

// TODO 可以考虑优化下 是否有更trick的方法
func addBinary(a string, b string) string {
	al := len(a)
	bl := len(b)
	var maxL int
	if al > bl {
		maxL = al
	} else {
		maxL = bl
	}
	addOne := false
	array := make([]int, maxL+1)
	for i := 1; i <= maxL+1; i++ {
		av := getValue(a, al, i)
		bv := getValue(b, bl, i)
		num := av + bv
		if addOne {
			num += 1
		}
		if num == 2 {
			addOne = true
			array[i-1] = 0
		} else if num == 3 {
			addOne = true
			array[i-1] = 1
		} else {
			addOne = false
			array[i-1] = num
		}
	}
	var reverse []string
	j := 0
	for i := maxL; i >= 0; i-- {
		x := array[i]
		if reverse == nil {
			if x != 0 {
				reverse = make([]string, i+1)
			} else if i == 0 {
				return strconv.Itoa(x)
			} else {
				continue
			}
		}
		reverse[j] = strconv.Itoa(x)
		j++
	}
	return strings.Join(reverse, "")
}

func getValue(str string, length int, i int) int {
	if i > length {
		return 0
	} else {
		return int(str[length-i]) - 48
	}
}
