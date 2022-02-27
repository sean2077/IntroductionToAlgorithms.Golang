package sort_algorithms

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type sortfunc func([]int)

func init() {
	rand.Seed(time.Now().Unix())
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func CheckSorted(arr []int) bool {
	for i := range arr {
		if i == 0 {
			continue
		}
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func TestSort(fn sortfunc, arr []int, t *testing.T) {
	fmt.Printf("%s\n", GetFunctionName(fn))
	fmt.Printf(" [input]: %v \n", arr)
	fn(arr)
	fmt.Printf(" [output]: %v \n", arr)
	assert.True(t, CheckSorted(arr))
}

func QuickTestSort(fn sortfunc, t *testing.T) {
	arr := []int{4, 5, 2, 3, 1}
	TestSort(fn, arr, t)

	bigArr := rand.Perm(100)
	TestSort(fn, bigArr, t)
}
