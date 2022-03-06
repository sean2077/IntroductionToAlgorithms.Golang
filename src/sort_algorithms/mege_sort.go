package sort_algorithms

// 使用下标的写法（类C++/Java）

// // 合并 [l,m] 和 [m+1,r]
// func merge(arr []int, l, m, r int) {

// 	raw := make([]int, r-l+1)
// 	copy(raw, arr[l:r+1])

// 	i := 0
// 	j := m - l + 1
// 	m = m - l
// 	sz := len(raw)

// 	for k := l; k <= r; k++ {
// 		if j >= sz || i <= m && raw[i] < raw[j] {
// 			arr[k] = raw[i]
// 			i++
// 		} else {
// 			arr[k] = raw[j]
// 			j++
// 		}
// 	}
// }

// // 归并排序 arr [l,r]
// func mergeSort(arr []int, l, r int) {
// 	if l < r {
// 		m := (l + r) / 2
// 		mergeSort(arr, l, m)
// 		mergeSort(arr, m+1, r)
// 		merge(arr, l, m, r)
// 	}
// }

// func MergeSort(arr []int) {
// 	mergeSort(arr, 0, len(arr)-1)
// }

// 使用 golang slice 的写法

func merge(L, R []int) []int {
	result := []int{}
	i := 0
	j := 0
	lsz := len(L)
	rsz := len(R)
	for i < lsz || j < rsz {
		if j >= rsz || i < lsz && L[i] < R[j] {
			result = append(result, L[i])
			i++
		} else {
			result = append(result, R[j])
			j++
		}
	}

	return result
}

// 归并排序
func MergeSort(arr []int) []int {
	sz := len(arr)
	if sz < 2 {
		return arr
	}

	left := MergeSort(arr[:sz/2])
	right := MergeSort(arr[sz/2:])

	return merge(left, right)
}

// 逆序对
func mergeInvertParis(L, R []int) ([]int, int) {
	lsz := len(L)
	rsz := len(R)
	M := make([]int, 0, lsz+rsz)
	i := 0
	j := 0
	res := 0
	for i < lsz || j < rsz {
		if j >= rsz || i < lsz && L[i] < R[j] {
			M = append(M, L[i])
			i++
		} else {
			M = append(M, R[j])
			j++
			res += lsz - i
		}
	}
	return M, res
}

func invertParis(arr []int) ([]int, int) {
	sz := len(arr)
	if sz < 2 {
		return arr, 0
	}
	half := sz / 2
	L, l := invertParis(arr[:half])
	R, r := invertParis(arr[half:])
	M, m := mergeInvertParis(L, R)
	return M, l + r + m
}

func InvertPairs(arr []int) int {
	_, res := invertParis(arr)
	return res
}
