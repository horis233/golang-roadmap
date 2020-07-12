# Stacks and Queues

## Introduction

The stack is characterized by last in first out

According to this feature, you can temporarily save some data, and then use it to insert it in turn. It is often used for DFS.

It is usually used for BFS, similar to layer by layer search

## Stack

[Minimal Stack]

> Design a stack that supports push, pop, and top operations and can retrieve the smallest element in a constant time.

Method: Implemented with two stacks, a minimum stack is always guaranteed to terminate at the top

```go
Enter MinStack struct {
    Minutes[] int
    Stack[] int
}


/ **Initialize your data structure here. * /
func Constructor() MinStack {
    Back to MinStack {
        Minimum value: make([] int, 0),
        Stack: make([] int, 0),
    }
}


func (this * MinStack) Push (x int) {
    min: = this.GetMin()
    If x <min {
        this.min = append(this.min, x)
    }other{
        this.min = append(this.min, min)
    }
    this.stack = append(this.stack, x)
}


func (this * MinStack) Pop() {
    If len (this.stack) == 0 {
        return
    }
    this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}


func (this * MinStack) Top () int {
    If len (this.stack) == 0 {
        Returns 0
    }
    Returns this.stack [len(this.stack)-1]
}


func (this * MinStack) GetMin() int {
    If len(this.min) == 0 {
        Returns 1 << 31
    }
    min: = this.min [len(this.min)-1]
    Returns the minimum value
}


/ **
 *Your MinStack object will be instantiated and called like this:
 * obj: = Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3: = obj.Top();
 * param_4: = obj.GetMin();
 * /
```

[evaluate-reverse-polish-notation]

> **Polish expression calculation**> **Input:** [“ 2”, “1”,“ +”, “3”,“ *”]> **Output:** 9
> **Explanation:** ((2+1)*3) = 9

Method: Save the original element through the stack, encounter the expression increment operation, then push the result, repeat this process

```go
func evalRPN (token[] string) int {
    if len (tokens) == 0 {
        Returns 0
    }
    stack: = make([] int, 0)
    for i: = 0; i <len (tokens); i ++ {
        switch token [i] {
        case "+", "-", "*", "/":
            if len(stack)<2 {
                Returns -1
            }
            //Note: a is the dividend, b is the divisor
            b: = stack [len(stack)-1]
            a: = stack [len(stack)-2]
            stack = stack [:len(stack)-2]
            var result int
            switch token [i] {
            case "+":
                result = a + b
            case "-":
                Result = a-b
            case "*":
                Result = a * b
            case "/":
                Result = a/b
            }
            stack = append (stack, result)
        default:
            //Turn to numbers
            val, _: = strconv.Atoi (tokens [i])
            stack = append(stack, val)
        }
    }
    return stack [0]
}
```

[Decode String]

> Given an encoded string, return the decoded string.
> s = "3[a]2[bc]", return" aaabcbc".
> s = "3[a2[c]]", return "accaccacc".
> s = "2[abc]3[cd] ef", return" abcabccdcdcdef".

Method: operation through stack assistance

```go
func decodeString(s string) string {
    if len(s) == 0 {
        return ""
    }
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case ']':
            temp := make([]byte, 0)
            for len(stack) != 0 && stack[len(stack)-1] != '[' {
                v := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                temp = append(temp, v)
            }
            // pop '['
            stack = stack[:len(stack)-1]
            // pop num
            idx := 1
            for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
                idx++
            }
            // mention index edge
            num := stack[len(stack)-idx+1:]
            stack = stack[:len(stack)-idx+1]
            count, _ := strconv.Atoi(string(num))
            for j := 0; j < count; j++ {
                // put character back to stack
                for j := len(temp) - 1; j >= 0; j-- {
                    stack = append(stack, temp[j])
                }
            }
        default:
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}
```

Recursive DFS search template using stack

```go
boolean DFS(int root, int target) {
    Set<Node> visited;
    Stack<Node> s;
    add root to s;
    while (s is not empty) {
        Node cur = the top element in s;
        return true if cur is target;
        for (Node next : the neighbors of cur) {
            if (next is not in visited) {
                add next to s;
                add next to visited;
            }
        }
        remove cur from s;
    }
    return false;
}
```

[binary-tree-inorder-traversal]

> Given a binary tree, return its *middle order* traversal.

```go
//Thinking: Save the elements that have been accessed through the stack for the original return
func inorderTraversal (root * TreeNode) [] int {
    Result: = make([] int, 0)
    If root == nil {
        Return result
    }
    Stack: = make([] * TreeNode, 0)
    For len(stack)> 0 || root! = nil {
        For root! = nil {
            Stack = append (stack, root)
            root = root.Left //Always left
        }
        //pop up
        val:= stack [len(stack)-1]
        Stack=stack[:len(stack)-1]
        Result = append(result, val.Val)
        Root = rvalue
    }
    Return result
}
```

[Clone graph]

> Give you a reference to a routine in the undirected connected graph, please return to a deep copy (clone) of the graph.

```go
func cloneGraph(node ​​* Node)* Node {
    Visited: = make (map [* Node] * Node)
    Return to clone (visited node)
}
// 1 2
// 4 3
//Recursive cloning, booking elements that have been visited as filter conditions
func clone(node ​​* Node, visited map [* Node] * Node)*Node{
    if node==nil{
        return nil
    }
    // Have already visited and returned directly
    if v,ok:=visited[node];ok{
        return v
    }
    newNode:=&Node{
        Val:node.Val,
        Neighbors:make([]*Node,len(node.Neighbors)),
    }
    visited[node]=newNode
    for i:=0;i<len(node.Neighbors);i++{
        newNode.Neighbors[i]=clone(node.Neighbors[i],visited)
    }
    return newNode
}
```

[number-of-islands]

> Given a two-dimensional grid consisting of '1' (land) and '0' (water), calculate the number of islands. An island is surrounded by water, and it is connected by adjacent land in the horizontal or vertical direction. You can assume that all four sides of the grid are surrounded by water.

Method: traverse the possibility through deep search (note that the visited element is marked)

```go

func numIslands(grid [][]byte) int {
    var count int
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            if grid[i][j]=='1' && dfs(grid,i,j)>=1{
                count++
            }
        }
    }
    return count
}
func dfs(grid [][]byte,i,j int)int{
    if i<0||i>=len(grid)||j<0||j>=len(grid[0]){
        return 0
    }
    if grid[i][j]=='1'{
        // Mark has been visited (each point only needs to be visited once)
        grid[i][j]=0
        return dfs(grid,i-1,j)+dfs(grid,i,j-1)+dfs(grid,i+1,j)+dfs(grid,i,j+1)+1
    }
    return 0
}
```

[largest-rectangle-in-histogram]

> Given n non-negative integers, used to represent the height of each column in the histogram. Each pillar is adjacent to each other and has a width of 1.
> Find the maximum area of ​​the rectangle that can be outlined in this histogram.

Method: Find the area with the current column as the height, which is converted to find the left and right values ​​that are less than the current value

Use the stack to save the left element less than the current value

```go
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    stack := make([]int, 0)
    max := 0
    for i := 0; i <= len(heights); i++ {
        var cur int
        if i == len(heights) {
            cur = 0
        } else {
            cur = heights[i]
        }
        // The current height is less than the stack, then pop the elements in the stack to calculate the area
        for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
            pop := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            h := heights[pop]
            // calculate width
            w := i
            if len(stack) != 0 {
                peek := stack[len(stack)-1]
                w = i-peek-1
            }
            max = Max(max, h*w)
        }
        // Record the index to get the corresponding element
        stack = append(stack, i)
    }
    return max
}
func Max(a, b int) int {
    if a> b {
        return a
    }
        return b
}
```

## Queue

Commonly used for BFS width-first search

[implement-queue-using-stacks]

> Use stack to implement queue

```go
type MyQueue struct {
    stack []int
    back []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
        stack: make([]int, 0),
        back: make([]int, 0),
    }
}

// 1
// 3
// 5

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
    for len(this.back) != 0 {
        val := this.back[len(this.back)-1]
        this.back = this.back[:len(this.back)-1]
        this.stack = append(this.stack, val)
    }
    this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    for len(this.stack) != 0 {
        val := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        this.back = append(this.back, val)
    }
    if len(this.back) == 0 {
        return 0
    }
    val := this.back[len(this.back)-1]
    this.back = this.back[:len(this.back)-1]
    return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
    for len(this.stack) != 0 {
        val := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        this.back = append(this.back, val)
    }
    if len(this.back) == 0 {
        return 0
    }
    val := this.back[len(this.back)-1]
    return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.stack) == 0 && len(this.back) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
```

Binary tree level traversal

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

[01-matrix]

> Given a matrix of 0 and 1, find the distance from each element to the nearest 0.
> The distance between two adjacent elements is 1

```go
// BFS enters the queue from 0. After popping, the results of up, down, left, and right are calculated.
// 0 0 0 0
// 0 x 0 0
// x x x 0
// 0 x 0 0

// 0 0 0 0
// 0 1 0 0// 1 x 1 0
// 0 1 0 0

// 0 0 0 0
// 0 1 0 0
// 1 2 1 0
// 0 1 0 0
func updateMatrix(matrix [][]int) [][]int {
    q:=make([][]int,0)
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j]==0{
                // enter the queue
                point:=[]int{i,j}
                q=append(q,point)
            }else{
                matrix[i][j]=-1
            }
        }
    }
    directions:=[][]int{{0,1},{0,-1},{-1,0},{1,0}}
    for len(q)!=0{
        // Dequeue
        point:=q[0]
        q=q[1:]
        for _,v:=range directions{
            x:=point[0]+v[0]
            y:=point[1]+v[1]
            if x>=0&&x<len(matrix)&&y>=0&&y<len(matrix[0])&&matrix[x][y]==-1{
                matrix[x][y]=matrix[point[0]][point[1]]+1
                // Put the current element into the queue and perform a BFS
                q=append(q,[]int{x,y})
            }
        }
    }
    return matrix

}
```

## to sum up

- Familiar with stack usage scenarios
  - Last in first out, save temporary values
  - In-depth search using stack DFS
- Familiar with the usage scenarios of the queue
  - Use queue BFS breadth search

## Exercise

- [ ] [min-stack](https://leetcode-cn.com/problems/min-stack/)
- [ ] [evaluate-reverse-polish-notation](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)
- [ ] [decode-string](https://leetcode-cn.com/problems/decode-string/)
- [ ] [binary-tree-inorder-traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
- [ ] [clone-graph](https://leetcode-cn.com/problems/clone-graph/)
- [ ] [number-of-islands](https://leetcode-cn.com/problems/number-of-islands/)
- [ ] [largest-rectangle-in-histogram](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)
- [ ] [implement-queue-using-stacks](https://leetcode-cn.com/problems/implement-queue-using-stacks/)
- [ ] [01-matrix](https://leetcode-cn.com/problems/01-matrix/)
