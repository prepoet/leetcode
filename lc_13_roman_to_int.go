package main

var roman = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

// roman-to-integer
func romanToInt(s string) int {
	l := len(s)
	if l == 1 {
		// 单字符
		return roman[s]
	}
	// 2+长度的字符
	rlt, ok := roman[s]
	if ok {
		// 优先返回一些特殊值
		return rlt
	}
	num := 0
	last := 1000
	for i := 0; i < l; i++ {
		x := roman[string(s[i])]
		if x > last {
			// 现在的符号值大于前一个符号值
			num -= last
			x = roman[s[i-1:i+1]]
		}
		last = x
		num += x
	}
	return num
}
