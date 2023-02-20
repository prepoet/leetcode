package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	var before, after int
	// TODO 优化为步长为2遍历数组
	for i := 0; i < length && n > 0; i++ {
		if i+1 > length-1 {
			after = 0
		} else {
			after = flowerbed[i+1]
		}
		if flowerbed[i] == 0 && before == 0 && after == 0 {
			flowerbed[i] = 1
			n--
		}
		before = flowerbed[i]
	}
	return n == 0
}
