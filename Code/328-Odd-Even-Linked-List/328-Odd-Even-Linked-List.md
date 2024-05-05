# 328. Odd Even Linked List（链表）

Medium

Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

 

Example 1:
![oddeven-linked-list.jpg](https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg)
> Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]

Example 2:
![oddeven2-linked-list.jpg](https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg)
> Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
 

Constraints:
> The number of nodes in the linked list is in the range [0, 104].
-106 <= Node.val <= 106

---

## 思路
双指针，分别指向奇数和偶数节点
每次循环，奇数指针指向下一个奇数节点，偶数指针指向下一个偶数节点
最后将奇数节点的尾部指向偶数节点的头部

## 代码
```go
package main

import "fmt"

func main() {
	// head := [1,2,3,4,5,6,7,8]
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	printLinnkedList(head)
	res := oddEvenList(head)
	printLinnkedList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func oddEvenList(head *ListNode) *ListNode {

	// 如果链表为空或者只有一个节点，直接返回nil或者head
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	// odd指针指向奇数节点，even指针指向偶数节点
	odd, even := head, head.Next
	// evenHead指向偶数节点的头节点
	evenHead := even

	// 遍历链表，将奇数节点和偶数节点分开循环
	// 当even.Next != nil时，说明还有至少odd可以指向下一个节点，还可以继续操作
	for even.Next != nil {
		// 如果even.Next.Next != nil，则even也可以指向下一个节点，继续操作
		if even.Next.Next != nil {
			// 把odd和even都指向下一个节点
			odd.Next = even.Next
			even.Next = even.Next.Next
			odd = odd.Next
			even = even.Next
		} else {
			// 如果even.Next.Next == nil，则说明even已经到了最后一个节点，此时odd.Next指向even.Next，odd指向下一个节点
			odd.Next = even.Next
			odd = odd.Next
			even.Next = nil
		}
	}

	// 当odd链表和even链表都遍历完之后，将odd链表的尾节点指向even链表的头节点，链接成一个链表
	odd.Next = evenHead

	// 返回处理好的链表
	return head
}

// 打印链表方法
func printLinnkedList(head *ListNode) {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}
```