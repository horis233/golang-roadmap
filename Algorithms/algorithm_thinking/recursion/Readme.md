# Recursion

## Introduction

Turn big problems into small problems and solve each small problem in turn through recursion

## Example

[reverse-string](https://leetcode-cn.com/problems/reverse-string/)

> Write a function whose role is to reverse the input string. The input string is given as a character array `char[]`.

```go
func reverseString(s []byte) {
    res := make([]byte, 0)
    reverse(s, 0, &res)
    for i := 0; i <len(s); i++ {
        s[i] = res[i]
    }
}
func reverse(s []byte, i int, res *[]byte) {
    if i == len(s) {
        return
    }
    reverse(s, i+1, res)
    *res = append(*res, s[i])
}
```

[swap-nodes-in-pairs](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

> Given a linked list, exchange adjacent nodes in it in pairs, and return the exchanged linked list.
> **You can't just change the value inside the node**, but need to actually exchange the node.

```go
func swapPairs(head *ListNode) *ListNode {
    // Idea: Turn the linked list into a sub-problem, and then solve it in turn by recursive
    // First flip two, and then continue to flip the following nodes like this, and then connect these flipped nodes
    return helper(head)
}
func helper(head *ListNode)*ListNode{
    if head==nil||head.Next==nil{
        return head
    }
    // Save the head pointer for the next stage
    nextHead:=head.Next.Next
    // Flip the current stage pointer
    next:=head.Next
    next.Next=head
    head.Next=helper(nextHead)
    return next
}
```

[unique-binary-search-trees-ii](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)

> Given an integer n, generate all binary search trees composed of 1 ... n nodes.

```go
func generateTrees(n int) []*TreeNode {
    if n==0{
        return nil
    }
    return generate(1,n)

}
func generate(start,end int)[]*TreeNode{
    if start>end{
        return []*TreeNode{nil}
    }
    ans:=make([]*TreeNode,0)
    for i:=start;i<=end;i++{
        // Recursively generate all left and right subtrees
        lefts:=generate(start,i-1)
        rights:=generate(i+1,end)
        // Return after splicing left and right subtrees
        for j:=0;j<len(lefts);j++{
            for k:=0;k<len(rights);k++{
                root:=&TreeNode{Val:i}
                root.Left=lefts[j]
                root.Right=rights[k]
                ans=append(ans,root)
            }
        }
    }
    return ans
}
```

## Recursion + Memo

[fibonacci-number](https://leetcode-cn.com/problems/fibonacci-number/)

> Fibonacci number, usually expressed by F(n), and the resulting sequence is called Fibonacci sequence. The number sequence starts with 0 and 1, and each number after it is the sum of the previous two numbers. That is:
> F(0) = 0, F(1) = 1
> F(N) = F(N-1) + F(N-2), where N> 1.
> Given N, calculate F(N).

```go
func fib(N int) int {
    return dfs(N)
}
var m map[int]int=make(map[int]int)
func dfs(n int)int{
    if n <2{
        return n
    }
    // read cache
    if m[n]!=0{
        return m[n]
    }
    ans:=dfs(n-2)+dfs(n-1)
    // Cache the calculated value
    m[n]=ans
    return ans
}
```

## Exercise

- [ ] [reverse-string](https://leetcode-cn.com/problems/reverse-string/)
- [ ] [swap-nodes-in-pairs](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)
- [ ] [unique-binary-search-trees-ii](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)
- [ ] [fibonacci-number](https://leetcode-cn.com/problems/fibonacci-number/)