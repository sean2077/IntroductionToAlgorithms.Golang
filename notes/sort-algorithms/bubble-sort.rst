========
冒泡排序
========

Golang 实现
-----------

.. code-block:: go

    // 冒泡排序
    func BubbleSort(arr []int) {
        size := len(arr)
        for i := 0; i < size; i++ {
            for j := size - 1; j > i; j-- {
                if arr[j-1] > arr[j] {
                    arr[j-1], arr[j] = arr[j], arr[j-1]
                }
            }
        }
    }