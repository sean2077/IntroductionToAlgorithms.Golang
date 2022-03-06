========
归并排序
========

.. _分治模式:

分治模式
--------

**分治模式** 在每层递归时都有三个步骤：

1. **分解** 原问题为若干子问题，这些子问题是原问题的规模较小的实例。
2. **解决** 这些子问题，递归地求解各子问题（若子问题规模足够小，则直接求解）。
3. **合并** 这些子问题的解成原问题的解。

算法
----

**归并排序** 算法完全遵循 :ref:`分治模式` :

1. **分解** 分解排序 n 个元素的序列成排序 n/2 个元素的两个子序列。
2. **解决** 使用归并排序递归地排序两个子序列。
3. **合并** 合并两个已排序的子序列。

Golang 实现
-----------

.. code-block:: go

    func merge(left, right []int) []int {
        result := []int{}
        i := 0
        j := 0
        lsz := len(left)
        rsz := len(right)
        for i < lsz && j < rsz {
            if j >= rsz || i < lsz && left[i] < right[j] {
                result = append(result, left[i])
                i++
            } else {
                result = append(result, right[j])
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