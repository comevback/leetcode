# 83. Remove Duplicates from Sorted List
Solved
Easy
Topics
Companies
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.


Example 1:
![alt text](https://assets.leetcode.com/uploads/2021/01/04/list1.jpg)
> Input: head = [1,1,2]
Output: [1,2]

Example 2:
![alt text](https://assets.leetcode.com/uploads/2021/01/04/list2.jpg)
> Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.

---

# Code
```go
package main

func main() {
	var head *ListNode
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}

	res := deleteDuplicates1(head)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 单指针
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		if current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

// 2. 双指针
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next

	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val
		} else {
			fast = fast.Next
		}
	}

	slow.Next = nil
	return head
}
```