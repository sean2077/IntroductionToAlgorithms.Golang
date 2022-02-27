package sort_algorithms

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	QuickTestSort(InsertionSort, t)
}
