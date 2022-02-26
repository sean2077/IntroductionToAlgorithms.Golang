package sort_algorithms

import (
	"fmt"
	"reflect"
	"runtime"
)

type sortfunc func([]int)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestSort(fn sortfunc, arr []int) {
	fmt.Printf("%s\n", GetFunctionName(fn))
	fmt.Printf(" [input]: %v \n", arr)
	fn(arr)
	fmt.Printf(" [output]: %v \n", arr)
}

var QuickTestArr = []int{4, 5, 2, 3, 1}

func QuickTestSort(fn sortfunc) {
	arr := make([]int, len(QuickTestArr))
	copy(arr, QuickTestArr)
	TestSort(fn, arr)
}
