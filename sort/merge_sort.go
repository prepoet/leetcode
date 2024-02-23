package sort

func mergeSort(array []int, low int, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(array, low, mid)
		mergeSort(array, mid+1, high)
		merge(array, low, mid, high)
	}
}

// 将有序数组 nums[lo..mid] 和有序数组 nums[mid+1..hi]
func merge(array []int, low int, mid int, high int) {

}
