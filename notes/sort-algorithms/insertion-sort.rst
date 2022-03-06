========
插入排序
========

Golang 实现
-----------

.. code-block:: go

    // 插入排序
    func InsertionSort(arr []int) []int {
        for j := 1; j < len(arr); j++ {
            val := arr[j]
            for i := j - 1; i >= 0; i-- {
                if arr[i] <= val {
                    arr[i+1] = val
                    break
                }
                arr[i+1] = arr[i]
            }
        }
        return arr
    }
