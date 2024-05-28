# 86. Partition List

Medium

Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.


Example 1:
![partition](https://assets.leetcode.com/uploads/2021/01/04/partition.jpg)
> Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]

Example 2:
> Input: head = [2,1], x = 2
Output: [1,2]
 

Constraints:

The number of nodes in the list is in the range [0, 200].
-100 <= Node.val <= 100
-200 <= x <= 200

---

# Code
```go
package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	x := 3
	partition(head, x)
	fmt.Println(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var dummy1 *ListNode = &ListNode{}
	var dummy2 *ListNode = &ListNode{}

	var smaller = dummy1
	var bigger = dummy2
	var current = head

	for current != nil {
		if current.Val < x {
			newNode := &ListNode{Val: current.Val} // 这里不能直接赋值，因为会形成环，所以要重新创建一个节点
			smaller.Next = newNode
			smaller = smaller.Next
		} else {
			newNode := &ListNode{Val: current.Val}
			bigger.Next = newNode
			bigger = bigger.Next
		}
		current = current.Next

		// 如果不每次都用新节点，那么每次都要把 current.Next 置为nil，断开
		// temp := current.Next
		// current.Next = nil
		// current = temp
	}

	smaller.Next = dummy2.Next
	return dummy1.Next
}
```