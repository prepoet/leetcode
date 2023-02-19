package leetcode

func candy(ratings []int) int {
	length := len(ratings)
	if length < 2 {
		// 注意边界值处理
		return length
	}
	// go slice 变长数组初始化，go 数组初始0值，不再初始赋1，后面汇总再操作+1
	result := make([]int, length)
	// 从左到右遍历 右分数>左 右=左+1
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			result[i] = result[i-1] + 1
		}
	}
	// 从右到左 左分数>右 且 左分到的糖果小于右 左=右+1
	for i := length - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && result[i-1] <= result[i] {
			result[i-1] = result[i] + 1
		}
	}
	// 汇总总糖果数
	sum := 0
	for i := 0; i < length; i++ {
		sum += result[i] + 1
	}
	return sum
}
