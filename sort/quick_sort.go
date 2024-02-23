package sort

func quickSort1(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, low int, high int) {
	if low < high {
		p := partition(array, low, high)
		quickSort(array, low, p-1)
		quickSort(array, p+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	// 选择第一个元素作为基准元素
	pivot := arr[low]
	i, j := low, high
	for i < j {
		// 从右往左找到第一个小于基准元素的元素
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		// 从左往右找到第一个大于基准元素的元素
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	// 将基准元素放到正确的位置
	arr[i] = pivot
	return i
}
