# Sort

## Sorting Algorithms

### Quick Sort

```go
func QuickSort(nums []int) []int {
    // Method: divide an array into left and right segments, the left segment is smaller than the right segment
    quickSort(nums, 0, len(nums)-1)
    return nums
}

// In-place exchange, so pass in exchange index
func quickSort(nums []int, start, end int) {
    if start <end {
        // Divide and conquer: divide
        pivot := partition(nums, start, end)
        quickSort(nums, 0, pivot-1)
        quickSort(nums, pivot 1, end)
    }
}

// partition
func partition(nums []int, start, end int) int {
    // Select the last element as the reference pivot
    p := nums[end]
    i := start
    // The last value is the benchmark so there is no need to compare
    for j := start; j <end; j {
        if nums[j] <p {
            swap(nums, i, j)
            i
        }
    }
    // Change the reference value to the middle
    swap(nums, i, end)
    return i
}

// swap two elements
func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}
```

### Merge Sort

```go
func MergeSort(nums []int) []int {
    return mergeSort(nums)
}

func mergeSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    // Divide and conquer: divide into two sections
    mid := len(nums)/2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])
    // Combine two pieces of data
    result := merge(left, right)
    return result
}

func merge(left, right []int) (result []int) {
    // merge cursor on both sides of the array
    l := 0
    r := 0
    // Be careful not to cross the boundary
    for l <len(left) && r <len(right) {
        // whoever merges
        if left[l]> right[r] {
            result = append(result, right[r])
            r
        } else {
            result = append(result, left[l])
            l
        }
    }
    // The remaining parts are merged
    result = append(result, left[l:]...)
    result = append(result, right[r:]...)
    return
}
```

### Heap sort

The full binary tree represented by an array complete binary tree

> Full Binary Tree VS Other Binary Trees

![image.png](https://img.fuiboom.com/img/tree_type.png)

[Animation Show](https://www.bilibili.com/video/av18980178/)

![image.png](https://img.fuiboom.com/img/heap.png)

```go
func HeapSort(a []int) []int {
    // 1. Unordered array a
    // 2. Construct the unordered array a as a big root heap
    for i := len(a)/2-1; i >= 0; i-- {
        sink(a, i, len(a))
    }
    // 3. Exchange a[0] and a[len(a)-1]
    // 4. Then continue to sink the previous section of the array to maintain the heap structure, and loop like this
    for i := len(a)-1; i >= 1; i-- {
        // Fill in the value from back to front
        swap(a, 0, i)
        // The length in front is also reduced by one
        sink(a, 0, i)
    }
    return a
}

func sink(a []int, i int, length int) {
    for {
        // Left node index (starting from 0, so the left node is i*2 1)
        l := i*2 1
        // has node index
        r := i*2 2
        // idx saves the index of the larger value between the root, left and right
        idx := i
        // There is a left node, the value of the left node is larger, then the left node is taken
        if l <length && a[l]> a[idx] {
            idx = l
        }
        // There are nodes, and the value is larger, take the right node
        if r <length && a[r]> a[idx] {
            idx = r
        }
        // If the root node is larger, there is no need to sink
        if idx == i {
            break
        }
        // If the root node is smaller, exchange values ​​and continue to sink
        swap(a, i, idx)
        // Continue to sink idx nodes
        i = idx
    }
}

func swap(a []int, i, j int) {
    a[i], a[j] = a[j], a[i]
}

```

## Reference

[Top Ten Classic Sort](https://www.cnblogs.com/onepixel/p/7674659.html)

[Two fork pile](https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/er-cha-dui-xiang-jie-shi-xian-you-xian-ji-dui-lie)

## Exercise

- [ ] Handwriting quick sort, merge, heap sort
