# Dynamic programming

## Background

Let’s start with a topic~

As titled [triangle](https://leetcode-cn.com/problems/triangle/)

> Given a triangle, find the minimum path sum from top to bottom. Each step can only move to the adjacent node in the next row.

For example, given a triangle:

```text
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
```

The minimum path sum from top to bottom is 11 (ie, 2 + 3 + 5 + 1 = 11).

Use DFS (traversal or divide and conquer)

Traverse

![image.png](https://img.fuiboom.com/img/dp_triangle.png)

Divide and conquer

![image.png](https://img.fuiboom.com/img/dp_dc.png)

Optimize DFS, cache the calculated value (called: memory search, in essence: dynamic programming)

![image.png](https://img.fuiboom.com/img/dp_memory_search.png)

Dynamic programming is the method of turning big problems into small problems and solving small problems and recalculating is called dynamic programming.

The difference between dynamic programming and DFS

- Binary tree The sub-problem is that there is no intersection, so most binary trees can be solved by recursive or divide-and-conquer method, namely DFS
- Like triangle, there are repeated walks, **sub-problems have intersection**, so it can be solved by dynamic programming

Dynamic programming, bottom-up

```go
func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 || len(triangle[0]) == 0 {
        return 0
    }
    // 1. State definition: f[i][j] represents the shortest path starting from i, j to the last layer
    var l = len(triangle)
    var f = make([][]int, l)
    // 2. Initialization
    for i := 0; i <l; i++ {
        for j := 0; j <len(triangle[i]); j++ {
            if f[i] == nil {
                f[i] = make([]int, len(triangle[i]))
            }
            f[i][j] = triangle[i][j]
        }
    }
    // 3. Recursive solution
    for i := len(triangle)-2; i >= 0; i-- {
        for j := 0; j <len(triangle[i]); j++ {
            f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
        }
    }
    // 4. Answer
    return f[0][0]
}

func min(a, b int) int {
    if a> b {
        return b
    }
    return a
}

```

Dynamic planning, top-down

```go
// Test case:
// [
// [2],
// [3,4],
// [6,5,7],
// [4,1,8,3]
//]
func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 || len(triangle[0]) == 0 {
        return 0
    }
    // 1. State definition: f[i][j] represents the shortest path starting from 0,0 to i,j
    var l = len(triangle)
    var f = make([][]int, l)
    // 2. Initialization
    for i := 0; i <l; i++ {
        for j := 0; j <len(triangle[i]); j++ {
            if f[i] == nil {
                f[i] = make([]int, len(triangle[i]))
            }
            f[i][j] = triangle[i][j]
        }
    }
    // Recursive solution
    for i := 1; i <l; i++ {
        for j := 0; j <len(triangle[i]); j++ {
            // There are two situations here:
            // 1, the upper layer has no left value
            // 2. The upper layer has no right value
            if j-1 <0 {
                f[i][j] = f[i-1][j] + triangle[i][j]
            } else if j >= len(f[i-1]) {
                f[i][j] = f[i-1][j-1] + triangle[i][j]
            } else {
                f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
            }
        }
    }
    result := f[l-1][0]
    for i := 1; i <len(f[l-1]); i++ {
        result = min(result, f[l-1][i])
    }
    return result
}
func min(a, b int) int {
    if a> b {
        return b
    }
    return a
}
```

## Recursion and Dynamic Programming  relationship

Recursion is a way of implementing programs: self-calling of functions

```go
Function(x) {
...
Funciton(x-1);
...
}
```

Dynamic programming: is a problem-solving idea. The results of large-scale problems are calculated from the results of small-scale problems. Dynamic programming can be implemented recursively (Memorization Search)

## scenes to be used

Meet two conditions

- Meet one of the following conditions
  - Find the maximum/minimum value (Maximum/Minimum)
  - Is it feasible (Yes/No)
  - Find the feasible number (Count(\*))
- Can not sort or swap (Cannot sort/swap)

Such as the title: [longest-consecutive-sequence](https://leetcode-cn.com/problems/longest-consecutive-sequence/) The location can be exchanged, so there is no need for dynamic planning

## Four elements

1. **State**
   - Inspiration, creativity, storing the results of small-scale problems
2. Equation Function
   - The connection between states, how to use the small state to calculate the big state
3. Initialization
   - What is the most extreme small state, starting point
4. Answer
   - What is the biggest state, the end

## Common four types

1. Matrix DP (10%)
1. Sequence (40%)
1. Two Sequences DP (40%)
1. Backpack (10%)

> Note
>
>-Greedy algorithms have back-to-back answers to most questions, so if you can use dynamic programming, try to use dynamic rules instead of greedy algorithms

## 1. Matrix type (10%)

### [minimum-path-sum](https://leetcode-cn.com/problems/minimum-path-sum/)

> Given a *m* x *n* grid containing non-negative integers, please find a path from the upper left corner to the lower right corner so that the sum of the numbers on the path is the smallest.

Idea: Dynamic programming
1. State: f[x][y] The shortest path from the starting point to x,y
2. function: f[x][y] = min(f[x-1][y], f[x][y-1]) + A[x][y]
3. Intialize: f[0][0] = A[0][0], f[i][0] = sum(0,0 -> i,0), f[0][i] = sum( 0,0 -> 0,i)
4. answer: f[n-1][m-1]

```go
func minPathSum(grid [][]int) int {
    // Idea: dynamic programming
    // f[i][j] represents the smallest sum of i, j to 0, 0
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    // Reuse the original matrix list
    // Initialization: f[i][0], f[0][j]
    for i := 1; i <len(grid); i++ {
        grid[i][0] = grid[i][0] + grid[i-1][0]
    }
    for j := 1; j <len(grid[0]); j++ {
        grid[0][j] = grid[0][j] + grid[0][j-1]
    }
    for i := 1; i <len(grid); i++ {
        for j := 1; j <len(grid[i]); j++ {
            grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
    if a> b {
        return b
    }
    return a
}
```

### [unique-paths](https://leetcode-cn.com/problems/unique-paths/)

> A robot is located in the upper left corner of an m x n grid (the starting point is marked as "Start" in the figure below).
> The robot can only move one step down or right at a time. The robot tries to reach the bottom right corner of the grid (labeled "Finish" in the image below).
> Ask how many different paths are there in total?

```go
func uniquePaths(m int, n int) int {
    // f[i][j] represents the number of paths from i,j to 0,0
    f := make([][]int, m)
    for i := 0; i <m; i++ {
        for j := 0; j <n; j++ {
            if f[i] == nil {
                f[i] = make([]int, n)
            }
            f[i][j] = 1
        }
    }
    for i := 1; i <m; i++ {
        for j := 1; j <n; j++ {
            f[i][j] = f[i-1][j] + f[i][j-1]
        }
    }
    return f[m-1][n-1]
}
```

### [unique-paths-ii](https://leetcode-cn.com/problems/unique-paths-ii/)

> A robot is located in the upper left corner of an m x n grid (the starting point is marked as "Start" in the figure below).
> The robot can only move one step down or right at a time. The robot tries to reach the bottom right corner of the grid (labeled "Finish" in the image below).
> Ask how many different paths are there in total?
> Now consider that there are obstacles in the grid. How many different paths will there be from the upper left corner to the lower right corner?

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // f[i][j] = f[i-1][j] + f[i][j-1] and check for obstacles
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    f := make([][]int, m)
    for i := 0; i <m; i++ {
        for j := 0; j <n; j++ {
            if f[i] == nil {
                f[i] = make([]int, n)
            }
            f[i][j] = 1
        }
    }
    for i := 1; i <m; i++ {
        if obstacleGrid[i][0] == 1 || f[i-1][0] == 0 {
            f[i][0] = 0
        }
    }
    for j := 1; j <n; j++ {
        if obstacleGrid[0][j] == 1 || f[0][j-1] == 0 {
            f[0][j] = 0
        }
    }
    for i := 1; i <m; i++ {
        for j := 1; j <n; j++ {
            if obstacleGrid[i][j] == 1 {
                f[i][j] = 0
            } else {
                f[i][j] = f[i-1][j] + f[i][j-1]
            }
        }
    }
    return f[m-1][n-1]
}
```

## 2. Sequence type (40%)

### [climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs/)

> Suppose you are climbing stairs. It takes *n* steps to reach the top of the building.

```go
func climbStairs(n int) int {
    // f[i] = f[i-1] + f[i-2]
    if n == 1 || n == 0 {
        return n
    }
    f := make([]int, n+1)
    f[1] = 1
    f[2] = 2
    for i := 3; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}
```

### [jump-game](https://leetcode-cn.com/problems/jump-game/)

> Given an array of non-negative integers, you are initially at the first position of the array.
> Each element in the array represents the maximum length you can jump at that position.
> Determine whether you can reach the last position.

```go
func canJump(nums []int) bool {
    // Idea: Look at the last jump
    // Status: f[i] indicates whether it can jump from 0 to i
    // Derivation: f[i] = OR(f[j],j<i&&j can jump to i) Determine whether the last jump of all previous points can jump to the current point
    // Initialization: f[0] = 0
    // Result: f[n-1]
    if len(nums) == 0 {
        return true
    }
    f := make([]bool, len(nums))
    f[0] = true
    for i := 1; i <len(nums); i++ {
        for j := 0; j <i; j++ {
            if f[j] == true && nums[j]+j >= i {
                f[i] = true
            }
        }
    }
    return f[len(nums)-1]
}
```

### [jump-game-ii](https://leetcode-cn.com/problems/jump-game-ii/)

> Given an array of non-negative integers, you are initially at the first position of the array.
> Each element in the array represents the maximum length you can jump at that position.
> Your goal is to use the least number of jumps to reach the last position of the array.

```go
// v1 dynamic programming (for other languages, refer to v2 for timeout)
func jump(nums []int) int {
    // Status: f[i] represents the minimum number of times from the starting point to the current position
    // Derivation: f[i] = f[j],a[j]+j >=i,min(f[j]+1)
    // Initialization: f[0] = 0
    // Result: f[n-1]
    f := make([]int, len(nums))
    f[0] = 0
    for i := 1; i <len(nums); i++ {
        // f[i] maximum value is i
        f[i] = i
        // The result before traversal takes a minimum value +1
        for j := 0; j <i; j++ {
            if nums[j]+j >= i {
                f[i] = min(f[j]+1,f[i])
            }
        }
    }
    return f[len(nums)-1]
}
func min(a, b int) int {
    if a> b {
        return b
    }
    return a
}
```

```go
// v2 dynamic programming + greedy optimization
func jump(nums []int) int {
    n:=len(nums)
    f := make([]int, n)
    f[0] = 0
    for i := 1; i <n; i++ {
        // Just take the first point that can jump to the current position
        // Because the result set of the number of jumps is monotonically increasing, the greedy idea is correct
        idx:=0
        for idx<n&&idx+nums[idx]<i{
            idx++
        }
        f[i]=f[idx]+1
    }
    return f[n-1]
}

```

### [palindrome-partitioning-ii](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)

> Given a string _s_, split _s_ into some substrings so that each substring is a palindrome.
> Return the minimum number of divisions that meet the requirements.

```go
func minCut(s string) int {
    // state: f[i] The substring composed of "before i" characters needs to be cut at least several times (number-1 is the index)
    // function: f[i] = MIN{f[j]+1}, j <i && [j+1 ~ i] This paragraph is a palindrome
    // intialize: f[i] = i-1 (f[0] = -1)
    // answer: f[s.length()]
    if len(s) == 0 || len(s) == 1 {
        return 0
    }
    f := make([]int, len(s)+1)
    f[0] = -1
    f[1] = 0
    for i := 1; i <= len(s); i++ {
        f[i] = i-1
        for j := 0; j <i; j++ {
            if isPalindrome(s, j, i-1) {
                f[i] = min(f[i], f[j]+1)
            }
        }
    }
    return f[len(s)]
}
func min(a, b int) int {
    if a> b {
        return b
    }
    return a
}
func isPalindrome(s string, i, j int) bool {
    for i <j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
```

- When judging the palindrome string, you can use dynamic programming to calculate it in advance to reduce time complexity

### [longest-increasing-subsequence](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

> Given an unordered integer array, find the length of the longest ascending subsequence.

```go
func lengthOfLIS(nums []int) int {
    // f[i] represents the longest sequence length from 0 to i
    // f[i] = max(f[j])+1 ,a[j]<a[i]
    // f[0...n-1] = 1
    // max(f[0]...f[n-1])
    if len(nums) == 0 || len(nums) == 1 {
        return len(nums)
    }
    f := make([]int, len(nums))
    f[0] = 1
    for i := 1; i <len(nums); i++ {
        f[i] = 1
        for j := 0; j <i; j++ {
            if nums[j] <nums[i] {
                f[i] = max(f[i], f[j]+1)
            }
        }
    }
    result := f[0]
    for i := 1; i <len(nums); i++ {
        result = max(result, f[i])
    }
    return result

}
func max(a, b int) int {
    if a> b {
        return a
    }
    return b
}
```

### [word-break](https://leetcode-cn.com/problems/word-break/)

> Given a **non-empty** string *s* and a dictionary *wordDict* containing a list of **non-empty** words, determine whether *s* can be split into one or more by spaces in the dictionary Words that appear.

```go
func wordBreak(s string, wordDict []string) bool {
    // f[i] indicates whether the first i characters can be divided
    // f[i] = f[j] && s[j+1~i] in wordDict
    // f[0] = true
    // return f[len]

    if len(s) == 0 {
        return true
    }
    f := make([]bool, len(s)+1)
    f[0] = true
    max,dict := maxLen(wordDict)
    for i := 1; i <= len(s); i++ {
        l := 0
        if i-max> 0 {
            l = i-max
        }
        for j := l; j <i; j++ {
            if f[j] && inDict(s[j:i],dict) {
                f[i] = true
                break
            }
        }
    }
    return f[len(s)]
}



func maxLen(wordDict []string) (int,map[string]bool) {
    dict := make(map[string]bool)
    max := 0
    for _, v := range wordDict {
        dict[v] = true
        if len(v)> max {
            max = len(v)
        }
    }
    return max,dict
    }

func inDict(s string,dict map[string]bool) bool {
    _, ok := dict[s]
    return ok
}

```

summary

The common processing method is to place the 0 position, so that the problem is treated equally, the initialization is based on the original length+1, and the result is returned f[n]

- Status can be the first i
- Initialize length+1
- Value index=i-1
- Return value: f[n] or f[m][n]

## Two Sequences DP (40%)

### [longest-common-subsequence](https://leetcode-cn.com/problems/longest-common-subsequence/)

> Given two strings text1 and text2, return the longest common subsequence of these two strings.
> A subsequence of a character string refers to a new character string: it is a new character string formed by deleting certain characters (or without deleting any characters) from the original character string without changing the relative order of the characters .
> For example, "ace" is a subsequence of "abcde", but "aec" is not a subsequence of "abcde". The "common subsequence" of two strings is the subsequence shared by the two strings.

```go
func longestCommonSubsequence(a string, b string) int {
    // dp[i][j] The longest common subsequence of the first i characters of a and the first j characters of b
    // dp[m+1][n+1]
    // 'a d c e
    // '0 0 0 0 0
    // a 0 1 1 1 1
    // c 0 1 1 2 1
    //
    dp:=make([][]int,len(a)+1)
    for i:=0;i<=len(a);i++ {
        dp[i]=make([]int,len(b)+1)
    }
    for i:=1;i<=len(a);i++ {
        for j:=1;j<=len(b);j++ {
            // Equal to take the upper left element + 1, otherwise take the larger value of the left or upper
            if a[i-1]==b[j-1] {
                dp[i][j]=dp[i-1][j-1]+1
            } else {
                dp[i][j]=max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[len(a)][len(b)]
}
func max(a,b int)int {
    if a>b{
        return a
    }
    return b
}
```

- go slice initialization

```go
dp:=make([][]int,len(a)+1)
for i:=0;i<=len(a);i++ {
    dp[i]=make([]int,len(b)+1)
}
```

- Traverse from 1 to the maximum length
- Index needs to be minus one

### [edit-distance](https://leetcode-cn.com/problems/edit-distance/)

> Give you two words word1 and word2, please calculate the minimum number of operations used to convert word1 to word2
> You can perform the following three operations on a word:
> Insert a character
> Delete a character
> Replace a character

Idea: Similar to the previous question, no operation is required if it is equal, otherwise the minimum number of operations for deletion, insertion, and replacement is +1

```go
func minDistance(word1 string, word2 string) int {
    // dp[i][j] indicates how many operations are required to edit the first i characters of a string to the first j characters of b string
    // dp[i][j] = OR(dp[i-1][j-1], a[i]==b[j],min(dp[i-1][j],dp[i ][j-1],dp[i-1][j-1])+1)
    dp:=make([][]int,len(word1)+1)
    for i:=0;i<len(dp);i++{
        dp[i]=make([]int,len(word2)+1)
    }
    for i:=0;i<len(dp);i++{
        dp[i][0]=i
    }
    for j:=0;j<len(dp[0]);j++{
        dp[0][j]=j
    }
    for i:=1;i<=len(word1);i++{
        for j:=1;j<=len(word2);j++{
            // No operation is required for equality
            if word1[i-1]==word2[j-1] {
                dp[i][j]=dp[i-1][j-1]
            }else{ // Otherwise, take the value of the minimum number of delete, insert, and replace operations +1
                dp[i][j]=min(min(dp[i-1][j],dp[i][j-1]),dp[i-1][j-1])+1
            }
        }
    }
    return dp[len(word1)][len(word2)]
}
func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}
```

Description

> Another approach: MAXLEN(a,b)-LCS(a,b)

## Change and backpack (10%)

### [coin-change](https://leetcode-cn.com/problems/coin-change/)

> Given coins of different denominations and a total amount amount. Write a function to calculate the minimum number of coins required to make up the total amount. If no coin combination can make up the total amount, return -1.

Idea: Different from other DPs, i means money or capacity

```go
func coinChange(coins []int, amount int) int {
    // State dp[i] indicates the minimum number of coins formed when the amount is i
    // Derive dp[i] = min(dp[i-1], dp[i-2], dp[i-5])+1, premise i-coins[j]> 0
    // Initialize to the maximum value dp[i]=amount+1
    // Return value dp[n] or dp[n]>amount =>-1
    dp:=make([]int,amount+1)
    for i:=0;i<=amount;i++{
        dp[i]=amount+1
    }
    dp[0]=0
    for i:=1;i<=amount;i++{
        for j:=0;j<len(coins);j++{
            if i-coins[j]>=0 {
                dp[i]=min(dp[i],dp[i-coins[j]]+1)
            }
        }
    }
    if dp[amount]> amount {
        return -1
    }
    return dp[amount]

}
func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}
```

note

> dp[i-a[j]] decide whether a[j] will participate

### [backpack](https://www.lintcode.com/problem/backpack/description)

> Choose a number of items from n items to pack into the backpack, how full is the maximum? Suppose the size of the backpack is m, and the size of each item is A[i]

```go
func backPack (m int, A []int) int {
    // write your code here
    // f[i][j] The first i items can be loaded with j
    // f[i][j] = f[i-1][j] f[i-1][j-a[i] j>a[i]
    // f[0][0]=true f[...][0]=true
    // f[n][X]
    f:=make([][]bool,len(A)+1)
    for i:=0;i<=len(A);i++{
        f[i]=make([]bool,m+1)
    }
    f[0][0]=true
    for i:=1;i<=len(A);i++{
        for j:=0;j<=m;j++{
            f[i][j]=f[i-1][j]
            if j-A[i-1]>=0 && f[i-1][j-A[i-1]]{
                f[i][j]=true
            }
        }
    }
    for i:=m;i>=0;i--{
        if f[len(A)][i] {
            return i
        }
    }
    return 0
}

```

### [backpack-ii](https://www.lintcode.com/problem/backpack-ii/description)

> There are `n` items and a backpack of size `m`. Given array `A` represents the size of each item and array `V` represents the value of each item.
> Ask what is the total value that can be loaded into the backpack?

Idea: f[i][j] the first i items, put the maximum value of j backpack

```go
func backPackII (m int, A []int, V []int) int {
    // write your code here
    // f[i][j] the first i items, put into j backpack maximum value
    // f[i][j] =max(f[i-1][j] ,f[i-1][j-A[i]]+V[i]) Whether to add item A[i]
    // f[0][0]=0 f[0][...]=0 f[...][0]=0
    f:=make([][]int,len(A)+1)
    for i:=0;i<len(A)+1;i++{
        f[i]=make([]int,m+1)
    }
    for i:=1;i<=len(A);i++{
        for j:=0;j<=m;j++{
            f[i][j]=f[i-1][j]
            if j-A[i-1] >= 0{
                f[i][j]=max(f[i-1][j],f[i-1][j-A[i-1]]+V[i-1])
            }
        }
    }
    return f[len(A)][m]
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```

## Exercise

Matrix DP (10%)

- [ ] [triangle](https://leetcode-cn.com/problems/triangle/)
- [ ] [minimum-path-sum](https://leetcode-cn.com/problems/minimum-path-sum/)
- [ ] [unique-paths](https://leetcode-cn.com/problems/unique-paths/)
- [ ] [unique-paths-ii](https://leetcode-cn.com/problems/unique-paths-ii/)

Sequence (40%)

- [ ] [climbing-stairs](https://leetcode-cn.com/problems/climbing-stairs/)
- [ ] [jump-game](https://leetcode-cn.com/problems/jump-game/)
- [ ] [jump-game-ii](https://leetcode-cn.com/problems/jump-game-ii/)
- [ ] [palindrome-partitioning-ii](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)
- [ ] [longest-increasing-subsequence](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
- [ ] [word-break](https://leetcode-cn.com/problems/word-break/)

Two Sequences DP (40%)

- [ ] [longest-common-subsequence](https://leetcode-cn.com/problems/longest-common-subsequence/)
- [ ] [edit-distance](https://leetcode-cn.com/problems/edit-distance/)

Backpack & Coin Change (10%)

- [ ] [coin-change](https://leetcode-cn.com/problems/coin-change/)
- [ ] [backpack](https://www.lintcode.com/problem/backpack/description)
- [ ] [backpack-ii](https://www.lintcode.com/problem/backpack-ii/description)