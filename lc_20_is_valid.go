package leetcode

import "container/list"

// right -> left
var brackets = map[uint8]uint8{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	l := len(s)
	if l%2 == 1 {
		// 成对出现 奇数直接返回
		return false
	}
	stack := list.New()
	for i := 0; i < l; i++ {
		bracket := s[i]
		left, ok := brackets[bracket]
		if ok {
			// right pop
			tail := stack.Back()
			if tail != nil && left == tail.Value {
				// match
				stack.Remove(tail)
			} else {
				return false
			}
		} else {
			// left push
			stack.PushBack(bracket)
		}
	}
	return stack.Len() == 0
}
