# Binary search

## Binary search template

Given an **ordered array** and target value, find the index of the first/last/any occurrence, if there is no return -1

Four elements of the template

- 1. Initialization: start=0, end=len-1
- 2. Loop exit condition: start + 1 <end
- 3. Compare midpoint and target value: A[mid] ==, <,> target
- 4. Determine whether the last two elements match: A[start], A[end]? Target

Time complexity O(logn), the usage scenario is generally the search of ordered arrays

Examples

[binary-search](https://leetcode-cn.com/problems/binary-search/)

> Given a n-element ordered (ascending) integer array nums and a target value target, write a function to search for the target in nums, and return the subscript if the target value exists, otherwise return -1.

```go
// The most common template for binary search
func search(nums []int, target int) int {
    // 1. Initialize start and end
    start := 0
    end := len(nums)-1
    // 2, deal with for loop
    for start+1 <end {
        mid := start + (end-start)/2
        // 3. Compare a[mid] and target value
        if nums[mid] == target {
            end = mid
        } else if nums[mid] <target {
            start = mid
        } else if nums[mid]> target {
            end = mid
        }
    }
    // 4. The last two elements are left, manual judgment
    if nums[start] == ​​target {
        return start
    }
    if nums[end] == target {
        return end
    }
    return -1
}
```

Most of the binary search questions can use this template, and then do a little special logic

In addition, there are some other templates for binary search as shown below. Most of the scene templates #3 can solve the problem, and they can also find the first/last position, which is more widely used.

![binary_search_template](https://img.fuiboom.com/img/binary_search_template.png)

So just use template #3, the detailed comparison can be introduced in this article: [Binary Search Template](https://leetcode-cn.com/explore/learn/card/binary-search/212/template-analysis/847/)

If it is the simplest binary search, you don’t need to find the first, last position, or there are no duplicate elements, you can use template #1, the code is more concise

```go
// More convenient when searching without duplicate elements
func search(nums []int, target int) int {
    start := 0
    end := len(nums)-1
    for start <= end {
        mid := start + (end-start)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] <target {
            start = mid+1
        } else if nums[mid]> target {
            end = mid-1
        }
    }
    // If not found, start is the first index greater than target
    // If you search in binary in the B+ tree structure, you can return start
    // This way you can continue to search for child nodes, such as: node:=node.Children[start]
    return -1
}
```

## Common Questions

### [search-for-range](https://www.lintcode.com/problem/search-for-a-range/description)

> Given a sorted array containing n integers, find the start and end of the given target value target.
> If the target value is not in the array, it returns `[-1, -1]`

Idea: The core point is to find the index of the first target and the index of the last target, so use two binary searches to find the first and last positions respectively

```go
func searchRange (A []int, target int) []int {
    if len(A) == 0 {
        return []int{-1, -1}
    }
    result := make([]int, 2)
    start := 0
    end := len(A)-1
    for start+1 <end {
        mid := start + (end-start)/2
        if A[mid]> target {
            end = mid
        } else if A[mid] <target {
            start = mid
        } else {
            // If they are equal, you should continue looking to the left to find the position of the first target value
            end = mid
        }
    }
    // Search the left index
    if A[start] == ​​target {
        result[0] = start
    } else if A[end] == target {
        result[0] = end
    } else {
        result[0] = -1
        result[1] = -1
        return result
    }
    start = 0
    end = len(A)-1
    for start+1 <end {
        mid := start + (end-start)/2
        if A[mid]> target {
            end = mid
        } else if A[mid] <target {
            start = mid
        } else {
            // If they are equal, you should continue looking to the right to find the position of the last target value
            start = mid
        }
    }
    // Search the index on the right
    if A[end] == target {
        result[1] = end
    } else if A[start] == ​​target {
        result[1] = start
    } else {
        result[0] = -1
        result[1] = -1
        return result
    }
    return result
}
```

### [search-insert-position](https://leetcode-cn.com/problems/search-insert-position/)

> Given a sorted array and a target value, find the target value in the array and return its index. If the target value does not exist in the array, return the position where it will be inserted in order.

```go
func searchInsert(nums []int, target int) int {
    // Idea: find the position of the first element >= target
    start := 0
    end := len(nums)-1
    for start+1 <end {
        mid := start + (end-start)/2
        if nums[mid] == target {
            // mark the start position
            start = mid
        } else if nums[mid]> target {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start] >= target {
        return start
    } else if nums[end] >= target {
        return end
    } else if nums[end] <target {// target value is greater than all values
        return end + 1
    }
    return 0
}
```

### [search-a-2d-matrix](https://leetcode-cn.com/problems/search-a-2d-matrix/)

> Write an efficient algorithm to judge whether there is a target value in the m x n matrix. The matrix has the following characteristics:
>
>-The integers in each row are sorted in ascending order from left to right.
>-The first integer on each line is greater than the last integer on the previous line.

```go
func searchMatrix(matrix [][]int, target int) bool {
    // Idea: Convert 2 latitude array to 1 dimension array for binary search
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    row := len(matrix)
    col := len(matrix[0])
    start := 0
    end := row*col-1
    for start+1 <end {
        mid := start + (end-start)/2
        // Get the corresponding value of 2 latitude array
        val := matrix[mid/col][mid%col]
        if val> target {
            end = mid
        } else if val <target {
            start = mid
        } else {
            return true
        }
    }
    if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target{
        return true
    }
    return false
}
```

### [first-bad-version](https://leetcode-cn.com/problems/first-bad-version/)

> Suppose you have n versions [1, 2, ..., n], and you want to find the first wrong version that caused all subsequent versions to go wrong.
> You can call bool isBadVersion(version) interface to determine whether the version number is wrong in the unit test. Implement a function to find the first wrong version. You should minimize the number of API calls.

```go
func firstBadVersion(n int) int {
    // Idea: binary search
    start := 0
    end := n
    for start+1 <end {
        mid := start + (end-start)/2
        if isBadVersion(mid) {
            end = mid
        } else if isBadVersion(mid) == false {
            start = mid
        }
    }
    if isBadVersion(start) {
        return start
    }
    return end
}
```

### [find-minimum-in-rotated-sorted-array](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

> Suppose that the array sorted in ascending order is rotated at a point unknown in advance (for example, the array [0,1,2,4,5,6,7] may become [4,5,6,7,0 ,1,2] ).
> Please find the smallest element.

```go
func findMin(nums []int) int {
    // Idea: // The last value is used as the target, then move to the left, and finally compare the values ​​of start and end
    if len(nums) == 0 {
        return -1
    }
    start := 0
    end := len(nums)-1

    for start+1 <end {
        mid := start + (end-start)/2
        // The last element value is target
        if nums[mid] <= nums[end] {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start]> nums[end] {
        return nums[end]
    }
    return nums[start]
}
```

### [find-minimum-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

> Suppose the array sorted in ascending order is rotated at a point unknown in advance
> (For example, the array [0,1,2,4,5,6,7] may become [4,5,6,7,0,1,2] ).
> Please find the smallest element. (Contains repeating elements)

```go
func findMin(nums []int) int {
    // Idea: Skip repeated elements, compare mid value and end value, and deal with it in two cases
    if len(nums) == 0 {
        return -1
    }
    start := 0
    end := len(nums)-1
    for start+1 <end {
        // remove duplicate elements
        for start <end && nums[end] == nums[end-1] {
            end--
        }
        for start <end && nums[start] == ​​nums[start+1] {
            start++
        }
        mid := start + (end-start)/2
        // The middle element is compared with the last element (to determine whether the middle point falls in the left rising area or the right rising area)
        if nums[mid] <= nums[end] {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start]> nums[end] {
        return nums[end]
    }
    return nums[start]
}
```

### [search-in-rotated-sorted-array](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

> Suppose that the array sorted in ascending order is rotated at a point unknown in advance.
> (For example, the array [0,1,2,4,5,6,7] may become [4,5,6,7,0,1,2] ).
> Search for a given target value, if the target value exists in the array, then return its index, otherwise it returns -1.
> You can assume that there are no duplicate elements in the array.

```go
func search(nums []int, target int) int {
    // Idea: // Two rising straight lines, four cases to judge
    if len(nums) == 0 {
        return -1
    }
    start := 0
    end := len(nums)-1
    for start+1 <end {
        mid := start + (end-start)/2
        // return directly
        if nums[mid] == target {
            return mid
        }
        // Judging in that interval, it may be divided into four cases
        if nums[start] <nums[mid] {
            if nums[start] <= target && target <= nums[mid] {
                end = mid
            } else {
                start = mid
            }
        } else if nums[end]> nums[mid] {
            if nums[end] >= target && nums[mid] <= target {
                start = mid
            } else {
                end = mid
            }
        }
    }
    if nums[start] == ​​target {
        return start
    } else if nums[end] == target {
        return end
    }
    return -1
}
```

be careful

> During the interview, you can directly draw a picture for auxiliary explanation, empty talk is easy to make everyone more circled

### [search-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)

> Suppose that the array sorted in ascending order is rotated at a point unknown in advance.
> (For example, the array [0,0,1,2,2,5,6] may become [2,5,6,0,0,1,2]).
> Write a function to determine whether the given target value exists in the array. Returns true if it exists, otherwise returns false. (Contains repeating elements)

```go
func search(nums []int, target int) bool {
    // Idea: // Two rising straight lines, judged in four cases, and deal with repeated numbers
    if len(nums) == 0 {
        return false
    }
    start := 0
    end := len(nums)-1
    for start+1 <end {
        // handle duplicate numbers
        for start <end && nums[start] == ​​nums[start+1] {
            start++
        }
        for start <end && nums[end] == nums[end-1] {
            end--
        }
        mid := start + (end-start)/2
        // return directly
        if nums[mid] == target {
            return true
        }
        // Judging in that interval, it may be divided into four cases
        if nums[start] <nums[mid] {
            if nums[start] <= target && target <= nums[mid] {
                end = mid
            } else {
                start = mid
            }
        } else if nums[end]> nums[mid] {
            if nums[end] >= target && nums[mid] <= target {
                start = mid
            } else {
                end = mid
            }
        }
    }
    if nums[start] == ​​target || nums[end] == target {
        return true
    }
    return false
}
```

## to sum up

Binary search core four elements (must back & understand)

- 1. Initialization: start=0, end=len-1
- 2. Loop exit condition: start + 1 <end
- 3. Compare midpoint and target value: A[mid] ==, <,> target
- 4. Determine whether the last two elements match: A[start], A[end]? Target

## Exercises

- [ ] [search-for-range](https://www.lintcode.com/problem/search-for-a-range/description)
- [ ] [search-insert-position](https://leetcode-cn.com/problems/search-insert-position/)
- [ ] [search-a-2d-matrix](https://leetcode-cn.com/problems/search-a-2d-matrix/)
- [ ] [first-bad-version](https://leetcode-cn.com/problems/first-bad-version)