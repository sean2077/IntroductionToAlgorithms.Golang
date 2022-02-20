package sort_algorithms

import (
	"fmt"
	"testing"
)

var Arr = []int{4, 5, 2, 3, 1}

func TestInsertionSort(t *testing.T) {
	copyArr := make([]int, len(Arr))
	copy(copyArr, Arr)
	fmt.Printf("【input】:%v       【output】:%v\n", Arr, insertionSort(copyArr))
}
