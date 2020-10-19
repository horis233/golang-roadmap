# Backtracking

## Background

Backtracking is often used to traverse all subsets of the list. It is a kind of DFS deep search. It is generally used for full permutation to exhaust all possibilities. The traversal process is actually a decision tree traversal process. The time complexity is generally O(N!). Unlike dynamic programming, there are overlapping sub-problems that can be optimized. The backtracking algorithm is purely brute force and the complexity is generally high.

## Template

```go
result = []
func backtrack (selection list, path):
    if meets the end condition:
        result.add(path)
        return
    for select in select list:
        make decisions
        backtrack (selection list, path)
        Cancel selection
```

The core is to make a selection from the selection list, and then search for the answer recursively downwards. If the path fails, return to cancel the selection.

## Example

### [subsets](https://leetcode-cn.com/problems/subsets/)

> Given a set of integer array nums without repeated elements, return all possible subsets (power sets) of the array.

Traversal process

![image.png](https://img.fuiboom.com/img/backtrack.png)

```go
func subsets(nums []int) [][]int {
    // save the final result
    result := make([][]int, 0)
    // save intermediate results
    list := make([]int, 0)
    backtrack(nums, 0, list, &result)
    return result
}

// nums given set
// pos The index of the element position added to the collection next
// list temporary result set (copy and save each time)
// result final result
func backtrack(nums []int, pos int, list []int, result *[][]int) {
    // Copy the temporary result and save it to the final result
    ans := make([]int, len(list))
    copy(ans, list)
    *result = append(*result, ans)
    // Select, process the result, and then cancel the selection
    for i := pos; i <len(nums); i++ {
        list = append(list, nums[i])
        backtrack(nums, i+1, list, result)
        list = list[0: len(list)-1]
    }
}
```

### [subsets-ii](https://leetcode-cn.com/problems/subsets-ii/)

> Given an integer array nums that may contain repeated elements, return all possible subsets (power sets) of the array. Note: The solution set cannot contain duplicate subsets.

```go
import (
    "sort"
)

func subsetsWithDup(nums []int) [][]int {
    // save the final result
    result := make([][]int, 0)
    // save intermediate results
    list := make([]int, 0)
    // sort first
    sort.Ints(nums)
    backtrack(nums, 0, list, &result)
    return result
}

// nums given set
// pos The index of the element position added to the collection next
// list temporary result set (copy and save each time)
// result final result
func backtrack(nums []int, pos int, list []int, result *[][]int) {
    // Copy the temporary result and save it to the final result
    ans := make([]int, len(list))
    copy(ans, list)
    *result = append(*result, ans)
    // Need to prune, process, and cancel the selection when selecting
    for i := pos; i <len(nums); i++ {
        // After sorting, if you encounter duplicate elements again, do not select this element
        if i != pos && nums[i] == nums[i-1] {
            continue
        }
        list = append(list, nums[i])
        backtrack(nums, i+1, list, result)
        list = list[0: len(list)-1]
    }
}
```

### [permutations](https://leetcode-cn.com/problems/permutations/)

> Given a sequence without repeated numbers, return all possible permutations.

Idea: You need to record the selected elements, and only return the results that meet the conditions

```go
func permute(nums []int) [][]int {
    result := make([][]int, 0)
    list := make([]int, 0)
    // Mark whether this element has been added to the result set
    visited := make([]bool, len(nums))
    backtrack(nums, visited, list, &result)
    return result
}

// nums input collection
// visited the currently recursively marked element
// list temporary result set (path)
// result final result
func backtrack(nums []int, visited []bool, list []int, result *[][]int) {
    // Return condition: the temporary result is the same as the length of the input collection, it is the full arrangement
    if len(list) == len(nums) {
        ans := make([]int, len(list))
        copy(ans, list)
        *result = append(*result, ans)
        return
    }
    for i := 0; i <len(nums); i++ {
        // Elements that have been added, skip directly
        if visited[i] {
            continue
        }
        // add elements
        list = append(list, nums[i])
        visited[i] = true
        backtrack(nums, visited, list, result)
        // remove element
        visited[i] = false
        list = list[0: len(list)-1]
    }
}
```

### [permutations-ii](https://leetcode-cn.com/problems/permutations-ii/)

> Given a sequence that can contain repeated numbers, return all non-repeating full permutations.

```go
import (
    "sort"
)

func permuteUnique(nums []int) [][]int {
    result := make([][]int, 0)
    list := make([]int, 0)
    // Mark whether this element has been added to the result set
    visited := make([]bool, len(nums))
    sort.Ints(nums)
    backtrack(nums, visited, list, &result)
    return result
}

// nums input collection
// visited the currently recursively marked element
// list temporary result set
// result final result
func backtrack(nums []int, visited []bool, list []int, result *[][]int) {
    // The temporary result is the same as the length of the input collection, which is the full arrangement
    if len(list) == len(nums) {
        subResult := make([]int, len(list))
        copy(subResult, list)
        *result = append(*result, subResult)
    }
    for i := 0; i <len(nums); i++ {
        // Elements that have been added, skip directly
        if visited[i] {
            continue
        }
        // The previous element is the same as the current one, and skipped if not visited
        if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
            continue
        }
        list = append(list, nums[i])
        visited[i] = true
        backtrack(nums, visited, list, result)
        visited[i] = false
        list = list[0: len(list)-1]
    }
}
```

## Exercise

- [ ] [subsets](https://leetcode-cn.com/problems/subsets/)
- [ ] [subsets-ii](https://leetcode-cn.com/problems/subsets-ii/)
- [ ] [permutations](https://leetcode-cn.com/problems/permutations/)
- [ ] [permutations-ii](https://leetcode-cn.com/problems/permutations-ii/)

Challenge topic

- [ ] [combination-sum](https://leetcode-cn.com/problems/combination-sum/)
- [ ] [letter-combinations-of-a-phone-number](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)
- [ ] [palindrome-partitioning](https://leetcode-cn.com/problems/palindrome-partitioning/)
- [ ] [restore-ip-addresses](https://leetcode-cn.com/problems/restore-ip-addresses/)
- [ ] [permutations](https://leetcode-cn.com/problems/permutations/)
