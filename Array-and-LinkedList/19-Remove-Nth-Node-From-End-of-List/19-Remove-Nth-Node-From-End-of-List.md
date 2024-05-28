# 19. Remove Nth Node From End of List

Medium

Given the head of a linked list, remove the nth node from the end of the list and return its head.


Example 1:
![remove_ex](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)
> Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
> Input: head = [1], n = 1
Output: []

Example 3:
> Input: head = [1,2], n = 1
Output: [1]
 

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 

Follow up: Could you do this in one pass?

---

# Code
```go
package main

func main() {
	var head *ListNode = &ListNode{1, nil}
	// head.Next = &ListNode{2, nil}
	// head.Next.Next = &ListNode{3, nil}
	// head.Next.Next.Next = &ListNode{4, nil}
	// head.Next.Next.Next.Next = &ListNode{5, nil}
	res := removeNthFromEnd(head, 1)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 一次遍历得到结果，前后两个指针，前指针走了n+1次之后，后指针才动
// 这样当前指针到达尾部的nil时，后指针到达了从后往前的第K个元素的前一个元素，可以删除第K个元素
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	var lead, delete *ListNode = dummy, dummy
	count := n
	for lead != nil {
		if count < 0 {
			delete = delete.Next
		}
		lead = lead.Next
		count -= 1
	}

	delete.Next = delete.Next.Next
	return dummy.Next
}
```