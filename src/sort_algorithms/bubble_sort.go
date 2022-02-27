package sort_algorithms

// 冒泡排序
func BubbleSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := size - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
