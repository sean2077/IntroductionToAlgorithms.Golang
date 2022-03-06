package sort_algorithms

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	QuickTestSort(MergeSort, t)
}

func TestInvertPairs(t *testing.T) {
	arr := []int{2, 3, 8, 6, 1}
	res := InvertPairs(arr)
	fmt.Printf("res: %d\n", res)
}
