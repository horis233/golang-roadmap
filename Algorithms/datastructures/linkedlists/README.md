# Linked list

## basic skills

Linked core points

- null/nil exception handling
- dummy node
- Fast and slow pointer
- Insert a node into the sorted linked list
- Remove a node from a linked list
- Flip linked list
- Merge two linked lists
- Find the middle node of the linked list

## Typical questions

### [remove-duplicates-from-sorted-list](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)

> Given a sorted linked list, delete all duplicate elements so that each element appears only once.

```go
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil {
        // Delete all and move to the next element
        for current.Next != nil && current.Val == current.Next.Val {
            current.Next = current.Next.Next
        }
        current = current.Next
    }
    return head
}
```

### [remove-duplicates-from-sorted-list-ii](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)

> Given a sorted linked list, delete all nodes that contain duplicate numbers, and keep only the numbers that are not repeated in the original linked list.

Method: The head of the linked list may be deleted, so use the dummy node to delete

```go
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    dummy := &ListNode{Val: 0}
    dummy.Next = head
    head = dummy

    var rmVal int
    for head.Next != nil && head.Next.Next != nil {
        if head.Next.Val == head.Next.Next.Val {
            // Record the deleted value for subsequent node judgment
            rmVal = head.Next.Val
            for head.Next != nil && head.Next.Val == rmVal {
                head.Next = head.Next.Next
            }
        } else {
            head = head.Next
        }
    }
    return dummy.Next
}
```

**Notes:**
- A->B->C delete B, A.next = C
- Delete assisted with a Dummy Node (allow head node to be variable)
- Access X.next and X.value must ensure that X != nil

### [reverse-linked-list](https://leetcode-cn.com/problems/reverse-linked-list/)

> Reverse a singly linked list.

Method: Use a prev node to save the forward pointer, temp to save the backward temporary pointer

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        // Save the current head.Next node to prevent it from being overwritten after reassignment
        // State after one round: nil<-1 2->3->4
        // prev head
        temp := head.Next
        head.Next = prev
        // pre move
        prev = head
        // head moves
        head = temp
    }
    return prev
}
```

### [reverse-linked-list-ii](https://leetcode-cn.com/problems/reverse-linked-list-ii/)

> Reverse the linked list from position *m* to *n*. Please use one scan to complete the reversal.

Method: traverse to m first, flip, and then splice the follow-up, pay attention to pointer processing

```go
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    // Method: first traverse to m, flip, then splice subsequent, pay attention to pointer processing
    // Enter: 1->2->3->4->5->NULL, m = 2, n = 4
    if head == nil {
        return head
    }
    // Head changes so use dummy node
    dummy := &ListNode{Val: 0}
    dummy.Next = head
    head = dummy
    // At the beginning: 0->1->2->3->4->5->nil
    var pre *ListNode
    var i = 0
    for i <m {
        pre = head
        head = head.Next
        i++
    }
    // After traversal: 1(pre)->2(head)->3->4->5->NULL
    // i = 1
    var j = i
    var next *ListNode
    // Used for intermediate node connection
    var mid = head
    for head != nil && j <= n {
        // The first cycle: 1 nil<-2 3->4->5->nil
        temp := head.Next
        head.Next = next
        next = head
        head = temp
        j++
    }
    // The loop needs to be executed four times
    // End of the loop: 1 nil<-2<-3<-4 5(head)->nil
    pre.Next = next
    mid.Next = head
    return dummy.Next
}
```

### [merge-two-sorted-lists](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

> Combine two ascending linked lists into a new ascending linked list and return. The new linked list is composed of all the nodes of the given two linked lists.

Method: Connect each element through the linked list of dummy node

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    head := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <l2.Val {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        head = head.Next
    }
    // Connection to node l1 has not been processed
    for l1 != nil {
        head.Next = l1
        head = head.Next
        l1 = l1.Next
    }
    // Connect l2 unprocessed node
    for l2 != nil {
        head.Next = l2
        head = head.Next
        l2 = l2.Next
    }
    return dummy.Next
}
```

### [partition-list](https://leetcode-cn.com/problems/partition-list/)

> Given a linked list and a specific value x, separate the linked list so that all nodes less than *x* are before nodes greater than or equal to *x*.

Method: Put the node greater than x into another linked list, and finally connect the two linked lists

```go
func partition(head *ListNode, x int) *ListNode {
    // Method: Put the node greater than x into another linked list, and finally connect the two linked lists
    // check
    if head == nil {
        return head
    }
    headDummy := &ListNode{Val: 0}
    tailDummy := &ListNode{Val: 0}
    tail := tailDummy
    headDummy.Next = head
    head = headDummy
    for head.Next != nil {
        if head.Next.Val <x {
            head = head.Next
        } else {
            // remove <x node
            t := head.Next
            head.Next = head.Next.Next
            // Put in another linked list
            tail.Next = t
            tail = tail.Next
        }
    }
    // Splice two linked lists
    tail.Next = nil
    head.Next = tailDummy.Next
    return headDummy.Next
}
```

Dumb node usage scenarios

> When the head node is uncertain, use the dumb node

### [sort-list](https://leetcode-cn.com/problems/sort-list/)

> Sort the linked list under *O*(*n* log *n*) time complexity and constant space complexity.

Method: merge sort, find midpoint and merge operation

```go
func sortList(head *ListNode) *ListNode {
    // Method: merge sort, find midpoint and merge operation
    return mergeSort(head)
}
func findMiddle(head *ListNode) *ListNode {
    // 1->2->3->4->5
    slow := head
    fast := head.Next
    // The fast pointer is nil first
    for fast !=nil && fast.Next! = nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    head := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <l2.Val {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        head = head.Next
    }
    // Connection to node l1 has not been processed
    for l1 != nil {
        head.Next = l1
        head = head.Next
        l1 = l1.Next
    }
    // Connect l2 unprocessed node
    for l2 != nil {
        head.Next = l2
        head = head.Next
        l2 = l2.Next
    }
    return dummy.Next
}
func mergeSort(head *ListNode) *ListNode {
    // If there is only one node, return this node directly
    if head == nil || head.Next == nil{
        return head
    }
    // find middle
    middle := findMiddle(head)
    // Disconnect the intermediate node
    tail := middle.Next
    middle.Next = nil
    left := mergeSort(head)
    right := mergeSort(tail)
    result := mergeTwoLists(left, right)
    return result
}
```

**Notes:**

- Fast and slow pointer to determine whether fast and fast.Next are nil values
- Recursive mergeSort requires disconnecting intermediate nodes
- The recursive return condition is head is nil or head.Next is nil

### [reorder-list](https://leetcode-cn.com/problems/reorder-list/)

> Given a singly linked list *L*: *L*→*L*→…→*L\_\_n*→*L*
> After rearranging it into: *L*→*L\_\_n*→*L*→*L\_\_n*→*L*→*L\_\_n*→…

Method: Find the midpoint and disconnect, flip the back part, and then merge the two linked lists

```go
func reorderList(head *ListNode) {
    // Method: Find the midpoint and disconnect, flip the back part, and then merge the two linked lists
    if head == nil {
        return
    }
    mid := findMiddle(head)
    tail := reverseList(mid.Next)
    mid.Next = nil
    head = mergeTwoLists(head, tail)
}
func findMiddle(head *ListNode) *ListNode {
    fast := head.Next
    slow := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    head := dummy
    toggle := true
    for l1 != nil && l2 != nil {
        // Node switching
        if toggle {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        toggle = !toggle
        head = head.Next
    }
    // Connection to node l1 has not been processed
    for l1 != nil {
        head.Next = l1
        head = head.Next
        l1 = l1.Next
    }
    // Connect l2 unprocessed node
    for l2 != nil {
        head.Next = l2
        head = head.Next
        l2 = l2.Next
    }
    return dummy.Next
}
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        // Save the current head.Next node to prevent it from being overwritten after reassignment
        // State after one round: nil<-1 2->3->4
        // prev head
        temp := head.Next
        head.Next = prev
        // pre move
        prev = head
        // head moves
        head = temp
    }
    return prev
}
```

### [linked-list-cycle](https://leetcode-cn.com/problems/linked-list-cycle/)

> Given a linked list, determine whether there are rings in the linked list.

Method: Fast and slow pointers, the same as the fast and slow pointers, there is a ring, prove: if there is a ring, the distance between the fast and slow pointers will be reduced by 1
![fast_slow_linked_list](https://img.fuiboom.com/img/fast_slow_linked_list.png)

```go
func hasCycle(head *ListNode) bool {
    // Method: Fast and slow pointers If the fast and slow pointers are the same, there is a ring. Proof: If there is a ring, the distance between the fast and slow pointers will decrease by 1
    if head == nil {
        return false
    }
    fast := head.Next
    slow := head
    for fast != nil && fast.Next != nil {
        // Compare pointers for equality (don't use val comparison!)
        if fast == slow {
            return true
        }
        fast = fast.Next.Next
        slow = slow.Next
    }
    return false
}
```

### [linked-list-cycle-ii](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

> Given a linked list, return to the first node in the linked list that begins to enter the ring. If the linked list has no rings, returns `null`.

Method: fast and slow pointers, after a quick and slow encounter, the slow pointer returns to the head, and the fast and slow pointers move together at the same pace, and the point of encounter is the ring point
![cycled_linked_list](https://img.fuiboom.com/img/cycled_linked_list.png)

```go
func detectCycle(head *ListNode) *ListNode {
    // Method: fast and slow pointers, after a quick and slow encounter, the slow pointer returns to the head, and the fast and slow pointers move together at the same pace, and the point of encounter is the ring entry point
    if head == nil {
        return head
    }
    fast := head.Next
    slow := head

    for fast != nil && fast.Next != nil {
        if fast == slow {
            // The slow pointer moves from the beginning again, and the fast pointer moves from the next node at the first intersection
            fast = head
            slow = slow.Next // Note
            // Compare pointer objects (do not compare pointer Val values)
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return slow
        }
        fast = fast.Next.Next
        slow = slow.Next
    }
    return nil
}
```

**Pit:**

- Directly compare objects during pointer comparison, do not use value comparison, there may be duplicate values ​​in the linked list
- After the first intersection, the fast pointer needs to move at the same speed with the head pointer from the next node

Another way is fast=head, slow=head

```go
func detectCycle(head *ListNode) *ListNode {
    // Method: fast and slow pointers, after a quick and slow encounter, one of the pointers returns to the head, and the fast and slow pointers move together at the same pace, and the point of encounter is the ring entry point
    // nb+a=2nb+a
    if head == nil {
        return head
    }
    fast := head
    slow := head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            // The pointer moves again from the beginning
            fast = head
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return slow
        }
    }
    return nil
}
```

The difference between these two methods is that **usually fast=head.Next is more**, because this way you can know the previous node of the midpoint, which can be used for deletion and other operations.

- fast if initialized to head.Next The midpoint is slow.Next
- fast is initialized to head, then the midpoint is in slow

### [palindrome-linked-list](https://leetcode-cn.com/problems/palindrome-linked-list/)

> Please determine whether a linked list is a palindrome linked list.

```go
func isPalindrome(head *ListNode) bool {
    // 1 2 nil
    // 1 2 1 nil
    // 1 2 2 1 nil
    if head==nil{
        return true
    }
    slow:=head
    // fast if initialized to head.Next, the midpoint is in slow.Next
    // fast is initialized to head, then the midpoint is in slow
    fast:=head.Next
    for fast!=nil&&fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }

    tail:=reverse(slow.Next)
    // Disconnect two linked lists (need to use the node before the midpoint)
    slow.Next=nil
    for head!=nil&&tail!=nil{
        if head.Val!=tail.Val{
            return false
        }
        head=head.Next
        tail=tail.Next
    }
    return true

}

func reverse(head *ListNode)*ListNode{
    // 1->2->3
    if head==nil{
        return head
    }
    var prev *ListNode
    for head!=nil{
        t:=head.Next
        head.Next=prev
        prev=head
        head=t
    }
    return prev
}
```

### [copy-list-with-random-pointer](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)

> Given a linked list, each node contains an additional random pointer that can point to any node or empty node in the linked list.
> Request to return a deep copy of this linked list.

Method: 1. The hash table stores pointers. 2. The copy node follows the original node.

```go
func copyRandomList(head *Node) *Node {
if head == nil {
return head
}
// Copy the node, right to the back
// 1->2->3 ==> 1->1'->2->2'->3->3'
cur := head
for cur != nil {
clone := &Node{Val: cur.Val, Next: cur.Next}
temp := cur.Next
cur.Next = clone
cur = temp
}
// Handle random pointer
cur = head
for cur != nil {
if cur.Random != nil {
cur.Next.Random = cur.Random.Next
}
cur = cur.Next.Next
}
// Separate two linked lists
cur = head
cloneHead := cur.Next
for cur != nil && cur.Next != nil {
temp := cur.Next
cur.Next = cur.Next.Next
cur = temp
}
// Original linked list header: head 1->2->3
// The cloned list head: cloneHead 1'->2'->3'
return cloneHead
}
```

## to sum up

Some points that must be mastered in the linked list. Through the following practice questions, most of the linked list topics are handy.

- null/nil exception handling
- dummy node
- Fast and slow pointer
- Insert a node into the sorted linked list
- Remove a node from a linked list
- Flip linked list
- Merge two linked lists
- Find the middle node of the linked list

## Exercise

-[ ] [remove-duplicates-from-sorted-list](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)
-[ ] [remove-duplicates-from-sorted-list-ii] (https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)
-[ ] [reverse-linked-list](https://leetcode-cn.com/problems/reverse-linked-list/)
-[ ] [reverse-linked-list-ii](https://leetcode-cn.com/problems/reverse-linked-list-ii/)
-[ ] [merge-two-sorted-lists](https://leetcode-cn.com/problems/merge-two-sorted-lists/)
-[ ] [partition-list](https://leetcode-cn.com/problems/partition-list/)
-[ ] [sort-list](https://leetcode-cn.com/problems/sort-list/)
-[ ] [reorder-list](https://leetcode-cn.com/problems/reorder-list/)
-[ ] [linked-list-cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
-[ ] [linked-list-cycle-ii](https://leetcode-cn.com/problems/linked-list-cycle-ii/)
-[ ] [palindrome-linked-list](https://leetcode-cn.com/problems/palindrome-linked-list/)
-[ ] [copy-list-with-random-pointer](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)