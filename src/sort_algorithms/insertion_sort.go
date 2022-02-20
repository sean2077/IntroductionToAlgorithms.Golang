package sort_algorithms

// 插入排序
func insertionSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		val := arr[j]
		i := j - 1
		for ; i >= 0; i-- {
			if arr[i] <= val {
				break
			}
			arr[i+1] = arr[i]
		}
		arr[i+1] = val
	}
	return arr
}
