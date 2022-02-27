package sort_algorithms

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	QuickTestSort(MergeSort, t)
}
