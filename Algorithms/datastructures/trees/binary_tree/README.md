# Binary tree

## Knowledge

### Binary tree traversal

**Preorder traversal**: **First visit the root node**, then traverse the left subtree first, then traverse the right subtree first
**Sequence traversal**: Traverse the left subtree first, **then visit the root node**, then traverse the right subtree first
**Post-order traversal**: Traverse the left subtree in sequence, then traverse the right subtree in sequence, **visit the root node**

**Notes**

- Decide what traversal is in root access order
- Left subtrees are prioritized right subtrees

#### Pre-order recursion

```go
func preorderTraversal(root *TreeNode) {
    if root==nil{
        return
    }
    // Visit the root first and then the left and right
    fmt.Println(root.Val)
    preorderTraversal(root.Left)
    preorderTraversal(root.Right)
}
```

#### Preorder non-recursive

```go
// non-recursive traversal
func preorderTraversal(root *TreeNode) []int {
    // non-recursive
    if root == nil{
        return nil
    }
    result:=make([]int,0)
    stack:=make([]*TreeNode,0)

    for root!=nil || len(stack)!=0{
        for root !=nil{
            // preorder traversal, so save the result first
            result=append(result,root.Val)
            stack=append(stack,root)
            root=root.Left
        }
        // pop
        node:=stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        root=node.Right
    }
    return result
}
```

#### inorder non-recursive

```go
// Method: Save the elements that have been accessed through the stack for the original return
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    stack := make([]*TreeNode, 0)
    for len(stack)> 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left // all the way to the left
        }
        // pop up
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)
        root = val.Right
    }
    return result
}
```

#### Postorder non-recursive

```go
func postorderTraversal(root *TreeNode) []int {
    // Use lastVisit to identify whether the right child node has popped up
    if root == nil {
        return nil
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    var lastVisit *TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // Look here first, don’t pop it first
        node:= stack[len(stack)-1]
        // The root node must pop up after the right node pops up
        if node.Right == nil || node.Right == lastVisit {
            stack = stack[:len(stack)-1] // pop
            result = append(result, node.Val)
            // Mark the current node has popped up
            lastVisit = node
        } else {
            root = node.Right
        }
    }
        return result
}
```

**Notes:**

- The core is: the root node must pop up after the right node pops up

#### DFS Deep Search-From top to bottom

```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    dfs(root, &result)
    return result
}

// V1: deep traversal, the result pointer is passed into the function as a parameter
func dfs(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    *result = append(*result, root.Val)
    dfs(root.Left, result)
    dfs(root.Right, result)
}
```

#### DFS in-depth search-from bottom to top (divide and conquer)

```go
// traverse through divide and conquer
func preorderTraversal(root *TreeNode) []int {
    result := divideAndConquer(root)
    return result
}
func divideAndConquer(root *TreeNode) []int {
    result := make([]int, 0)
    // Return condition (null & leaf)
    if root == nil {
        return result
    }
    // Divide and conquer (Divide)
    left := divideAndConquer(root.Left)
    right := divideAndConquer(root.Right)
    // merge result (Conquer)
    result = append(result, root.Val)
    result = append(result, left...)
    result = append(result, right...)
    return result
}
```

**Notes:**

> The difference between DFS deep search (from top to bottom) and divide and conquer method: the former generally passes the final result through the pointer parameter, the latter generally recursively returns the result and finally merges

#### BFS level traversal

```go
func levelOrder(root *TreeNode) [][]int {
    // Determine the elements of the next layer by the length of the previous layer
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue)> 0 {
        list := make([]int, 0)
        // Why should we take length?
        // Record how many elements there are in the current layer (traverse the current layer and add the next layer)
        l := len(queue)
        for i := 0; i <l; i++ {
            // Dequeue
            level := queue[0]
            queue = queue[1:]
            list = append(list, level.Val)
            if level.Left != nil {
                queue = append(queue, level.Left)
            }
            if level.Right != nil {
                queue = append(queue, level.Right)
            }
        }
        result = append(result, list)
    }
    return result
}
```

### Divide and conquer application

Process the parts separately before merging the results

Applicable scene

- Quick sort
- Merge sort
- Binary tree related issues

Divide and conquer template

- Recursive return conditions
- Segment processing
- Merger result

```go
func traversal(root *TreeNode) ResultType {
    // nil or leaf
    if root == nil {
        // do something and return
    }

    // Divide
    ResultType left = traversal(root.Left)
    ResultType right = traversal(root.Right)

    // Conquer
    ResultType result = Merge from left and right

    return result
}
```

#### Examples

```go
// Traverse the binary tree by dividing and conquering
func preorderTraversal(root *TreeNode) []int {
    result := divideAndConquer(root)
    return result
}
func divideAndConquer(root *TreeNode) []int {
    result := make([]int, 0)
    // Return condition (null & leaf)
    if root == nil {
        return result
    }
    // Divide and conquer (Divide)
    left := divideAndConquer(root.Left)
    right := divideAndConquer(root.Right)
    // merge result (Conquer)
    result = append(result, root.Val)
    result = append(result, left...)
    result = append(result, right...)
    return result
}
```

#### Merge sort

```go
func MergeSort(nums []int) []int {
    return mergeSort(nums)
}
func mergeSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    // Divide and conquer: divide is divided into two sections
    mid := len(nums) / 2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])
    // merge two pieces of data
    result := merge(left, right)
    return result
}
func merge(left, right []int) (result []int) {
    // merge cursor on both sides of the array
    l := 0
    r := 0
    // Not to cross the border
    for l <len(left) && r <len(right) {
        // Who merges who
        if left[l]> right[r] {
            result = append(result, right[r])
            r++
        } else {
            result = append(result, left[l])
            l++
        }
    }
    // merge the rest
    result = append(result, left[l:]...)
    result = append(result, right[r:]...)
    return
}
```

**Notes:**

> Recursion needs to return results for merging

#### Quick Sort

```go
func QuickSort(nums []int) []int {
    // Idea: Divide an array into left and right segments, the left segment is smaller than the right segment, similar to the divide and conquer method without merge process
    quickSort(nums, 0, len(nums)-1)
    return nums

    }
    // Exchange in place, so the exchange index is passed in
    func quickSort(nums []int, start, end int) {
        if start <end {
            // Divide and conquer: divide
            pivot := partition(nums, start, end)
            quickSort(nums, 0, pivot-1)
            quickSort(nums, pivot+1, end)
        }
    }
    // partition
    func partition(nums []int, start, end int) int {
        p := nums[end]
        i := start
        for j := start; j <end; j++ {
            if nums[j] <p {
                swap(nums, i, j)
                i++
            }
        }
            // Change the intermediate value to the reference value for comparison
        swap(nums, i, end)
        return i
    }
    func swap(nums []int, i, j int) {
        t := nums[i]
        nums[i] = nums[j]
        nums[j] = t
    }
```

**Notes:**

> The quick queue has no merge process because it is exchanged in place
> The index passed in is an existing index (eg: 0, length-1, etc.), out of bounds may cause a crash

Examples of common topics

#### maximum-depth-of-binary-tree

[maximum-depth-of-binary-tree](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

> Given a binary tree, find its maximum depth.

Method: Divide and conquer

```go
func maxDepth(root *TreeNode) int {
    // return condition processing
    if root == nil {
        return 0
    }
    // divide: divide the left and right subtrees separately
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)

    // conquer: merge left and right subtree results
    if left> right {
        return left + 1
    }
    return right + 1
}
```

#### balanced-binary-tree

[balanced-binary-tree](https://leetcode-cn.com/problems/balanced-binary-tree/)

> Given a binary tree, determine whether it is a highly balanced binary tree.

Method: Divide and conquer, balance on the left && balance on the right && height of left and right <= 1,
Because it is necessary to return whether it is balanced and high, either return two data or merge two data,
So -1 is used for unbalance, and >0 for tree height (ambiguity: a variable has two meanings).

```go
func isBalanced(root *TreeNode) bool {
    if maxDepth(root) == -1 {
        return false
    }
    return true
}
func maxDepth(root *TreeNode) int {
    // check
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)

    // Why does it return -1? (Variable has ambiguity)
    if left == -1 || right == -1 || left-right> 1 || right-left> 1 {
        return -1
    }
    if left> right {
        return left + 1
    }
    return right + 1
}
```

**Notes:**

> In general engineering, the result is returned by two variables, it is not recommended to use a variable to represent two meanings

#### binary-tree-maximum-path-sum

[binary-tree-maximum-path-sum](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)

> Given a **non-empty** binary tree, return its maximum path sum.

Method: divide-and-conquer method is divided into three cases: left sub-tree maximum path and maximum, right sub-tree maximum path and maximum, left and right sub-tree maximum plus root node maximum, two variables need to be saved: one to save sub-tree maximum path and , Save the left and right plus the root node sum, and then compare the two variables to select the maximum value

```go
type ResultType struct {
    SinglePath int // save the single maximum
    MaxPath int // save the maximum value (one-sided or two single-sided + root values)
}
func maxPathSum(root *TreeNode) int {
    result := helper(root)
    return result.MaxPath
}
func helper(root *TreeNode) ResultType {
    // check
    if root == nil {
        return ResultType{
            SinglePath: 0,
            MaxPath: -(1 << 31),
        }
    }
    // Divide
    left := helper(root.Left)
    right := helper(root.Right)

    // Conquer
    result := ResultType{}
    // Find the unilateral maximum
    if left.SinglePath> right.SinglePath {
        result.SinglePath = max(left.SinglePath + root.Val, 0)
    } else {
        result.SinglePath = max(right.SinglePath + root.Val, 0)
    }
    // Find the maximum value on both sides
    maxPath := max(right.MaxPath, left.MaxPath)
    result.MaxPath = max(maxPath,left.SinglePath+right.SinglePath+root.Val)
    return result
}
func max(a,b int) int {
    if a> b {
        return a
    }
    return b
}
```

#### lowest-common-ancestor-of-a-binary-tree

[lowest-common-ancestor-of-a-binary-tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

> Given a binary tree, find the nearest common ancestor of the two specified nodes in the tree.

Method: divide and conquer method, if there is a common ancestor of the left subtree or a common ancestor of the right subtree, the ancestor of the subtree is returned, otherwise the root node is returned

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // check
    if root == nil {
        return root
    }
    // Equal to return directly to the root node
    if root == p || root == q {
        return root
    }
    // Divide
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)


    // Conquer
    // The left and right sides are not empty, then the root node is an ancestor
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    if right != nil {
        return right
    }
    return nil
}
```

### BFS level application

#### binary-tree-level-order-traversal

[binary-tree-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)

> Give you a binary tree, please return the node value obtained by traversing ** in sequence**. (That is, access all nodes from left to right layer by layer)

Method: use a queue to record the elements of one layer, and then scan this layer of elements to add the next layer of elements to the queue (a number goes in and out once, so the complexity is O(logN))

```go
func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue)> 0 {
        list := make([]int, 0)
        // Why should we take length?
        // Record how many elements there are in the current layer (traverse the current layer and add the next layer)
        l := len(queue)
        for i := 0; i <l; i {
            // Dequeue
            level := queue[0]
            queue = queue[1:]
            list = append(list, level.Val)
            if level.Left != nil {
                queue = append(queue, level.Left)
            }
            if level.Right != nil {
                queue = append(queue, level.Right)
            }
        }
        result = append(result, list)
    }
    return result
}
```

#### binary-tree-level-order-traversal-ii

[binary-tree-level-order-traversal-ii](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)

> Given a binary tree, return the traversal of its node values ​​from bottom to top. (That is, from the layer where the leaf node is located to the layer where the root node is located, traversing from left to right layer by layer)

Method: On the basis of traversal, you can flip the result

```go
func levelOrderBottom(root *TreeNode) [][]int {
    result := levelOrder(root)
    // flip result
    reverse(result)
    return result
}
func reverse(nums [][]int) {
    for i, j := 0, len(nums)-1; i <j; i, j = i 1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
    return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue)> 0 {
    list := make([]int, 0)
            // Why should we take length?
            // Record how many elements there are in the current layer (traverse the current layer and add the next layer)
    l := len(queue)
    for i := 0; i <l; i {
        // Dequeue
        level := queue[0]
        queue = queue[1:]
        list = append(list, level.Val)
        if level.Left != nil {
            queue = append(queue, level.Left)
        }
        if level.Right != nil {
            queue = append(queue, level.Right)
        }
    }
    result = append(result, list)
    }
    return result
}
```

#### binary-tree-zigzag-level-order-traversal

[binary-tree-zigzag-level-order-traversal] (https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)

> Given a binary tree, return the zigzag hierarchy traversal of its node values. Z-shaped traversal

```go
func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    toggle := false
    for len(queue)> 0 {
        list := make([]int, 0)
        // Record how many elements are in the current layer (traverse the current layer, and then add the next layer)
        l := len(queue)
        for i := 0; i <l; i {
            // Dequeue
            level := queue[0]
            queue = queue[1:]
            list = append(list, level.Val)
            if level.Left != nil {
                queue = append(queue, level.Left)
            }
            if level.Right != nil {
                queue = append(queue, level.Right)
            }
        }
        if toggle {
            reverse(list)
        }
        result = append(result, list)
        toggle = !toggle
    }
    return result
}
func reverse(nums []int) {
    for i := 0; i <len(nums)/2; i {
        t := nums[i]
        nums[i] = nums[len(nums)-1-i]
        nums[len(nums)-1-i] = t
    }
}
```

### Binary search tree application

#### validate-binary-search-tree

[validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree/)

> Given a binary tree, determine whether it is a valid binary search tree.

Method 1: Traverse in order, check whether the result list is already ordered

Method 2: Divide and conquer, judge left MAX <root <right MIN

```go
// v1
func isValidBST(root *TreeNode) bool {
    result := make([]int, 0)
    inOrder(root, &result)
    // check order
    for i := 0; i <len(result)-1; i {
        if result[i] >= result[i 1] {
            return false
        }
    }
    return true
}

func inOrder(root *TreeNode, result *[]int) {
    if root == nil{
        return
    }
    inOrder(root.Left, result)
    *result = append(*result, root.Val)
    inOrder(root.Right, result)
}


```

```go
// divide and conquer
type ResultType struct {
    IsValid bool
    // Record the maximum and minimum values ​​on the left and right sides, and compare with the root node
    Max *TreeNode
    Min *TreeNode
}

func isValidBST2(root *TreeNode) bool {
    result := helper(root)
    return result.IsValid
}
func helper(root *TreeNode) ResultType {
    result := ResultType{}
    // check
    if root == nil {
        result.IsValid = true
        return result
    }

    left := helper(root.Left)
    right := helper(root.Right)

    if !left.IsValid || !right.IsValid {
        result.IsValid = false
        return result
    }
    if left.Max != nil && left.Max.Val >= root.Val {
        result.IsValid = false
        return result
    }
    if right.Min != nil && right.Min.Val <= root.Val {
        result.IsValid = false
        return result
    }

    result.IsValid = true
        // If there is a smaller 3 on the left, use a smaller node instead of 4
        //   5
        //  / \
        // 1   4
        //    / \
        //   3   6
    result.Min = root
    if left.Min != nil {
        result.Min = left.Min
    }
    result.Max = root
    if right.Max != nil {
        result.Max = right.Max
    }
    return result
}
```

#### insert-into-a-binary-search-tree

[insert-into-a-binary-search-tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)

> Given the root node of the binary search tree (BST) and the value to be inserted into the tree, insert the value into the binary search tree. Returns the root node of the binary search tree after insertion.

Method: Find the last leaf node that meets the insertion condition

```go
// DFS finds the insertion position
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        root = &TreeNode{Val: val}
        return root
    }
    if root.Val> val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
```

## to sum up

- Master recursive and non-recursive traversal of binary trees
- Understand DFS preorder traversal and divide and conquer method
- Understand BFS level traversal

## Exercise

- [ ] [maximum-depth-of-binary-tree](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
- [ ] [balanced-binary-tree](https://leetcode-cn.com/problems/balanced-binary-tree/)
- [ ] [binary-tree-maximum-path-sum](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)
- [ ] [lowest-common-ancestor-of-a-binary-tree] (https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
- [ ] [binary-tree-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
- [ ] [binary-tree-level-order-traversal-ii] (https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)
- [ ] [binary-tree-zigzag-level-order-traversal] (https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)
- [ ] [validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree/)
- [ ] [insert-into-a-binary-search-tree] (https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)