# 92. Reverse Linked List II
Solved
Medium
Topics
Companies
Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

 

Example 1:
![rev2ex2](https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg)
> Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Example 2:
> Input: head = [5], left = 1, right = 1
Output: [5]
 

Constraints:

The number of nodes in the list is n.
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
 

Follow up: Could you do it in one pass?

---

# Code
```go
package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printLinnkedList(head)
	res := reverseN(head, 4)
	printLinnkedList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// *****************************************************  迭代实现  ******************************************************
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// 如果链表为空或只有一个节点，直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	var index int = 0
	var current *ListNode

	// 创建一个哑节点，并将其Next指向head
	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	// 保存哑节点的引用，以便后续返回链表头部
	saveHead := dummy

	// 移动dummy节点到left-1的位置，以找到需要反转部分的前一个节点
	for index < left-1 {
		dummy = dummy.Next
		index += 1
	}
	// 当前节点指向需要反转部分的第一个节点
	current = dummy.Next

	// 创建一个新的头节点newHead，用于构建反转后的链表
	var newHead *ListNode = nil

	// 反转链表从left到right部分
	for current != nil {
		// 保存当前节点的下一个节点
		next := current.Next
		// 将当前节点的Next指向newHead（即反转链表的前一个节点）
		current.Next = newHead
		// 更新newHead为当前节点
		newHead = current
		// 更新当前节点为下一个节点
		current = next
		index += 1
		// 当到达right位置时，停止反转，并连接反转后的链表与剩余链表
		if index == right {
			cur := newHead
			// 找到反转后链表的最后一个节点
			for cur.Next != nil {
				cur = cur.Next
			}
			// 将反转后的链表的最后一个节点连接到原链表剩余部分
			cur.Next = next
			break
		}
	}

	// 将反转部分的头节点连接到原链表的前半部分
	dummy.Next = newHead

	// 返回新的链表头节点
	return saveHead.Next
}

// *****************************************************  递归实现  ******************************************************

// ** 反转链表中从left到right的部分（递归实现）
// ** 用到翻转链表的前N个节点的递归实现
func reverseBetweenRe(head *ListNode, left int, right int) *ListNode {
	// 如果left等于1，直接反转前right个节点
	if left == 1 {
		return reverseListReUntil(head, right)
	}

	// 递归调用，缩小反转区间
	rest := reverseBetweenRe(head.Next, left-1, right-1)
	head.Next = rest
	return head
}

// ** 反转链表的前N个节点 （递归实现） 思路和用递归翻转整个链表类似
var last *ListNode // ** 外部变量，用于保存翻转后的链表的尾节点
func reverseListReUntil(head *ListNode, N int) *ListNode {
	// 当N等于1时，保存下一个节点并返回当前节点作为新头节点
	if N == 1 {
		last = head.Next
		return head
	}
	nextHead := reverseListReUntil(head.Next, N-1)
	head.Next.Next = head
	head.Next = last

	return nextHead
}

// ** 反转链表的前N个节点 （递归实现） 优化
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 { //base case
		return head
	}

	reverseHead := reverseN(head.Next, n-1)
	// tmp 保存了后面不需要反转的部分，假设n=2，也就是递归到倒数第二个要翻转的元素时，这里的head.Next就是最后一个要被翻转的元素
	// 这里用tmp保存了head.Next.Next，也就是后面不需要反转的部分，然后把这次循环到的head放到head.Next的后面，最后再把tmp（不需要反转的部分）放到head的后面
	tmp := head.Next.Next
	head.Next.Next = head
	head.Next = tmp

	return reverseHead
}

// ** 反转整个链表（递归实现）
func reverseListRe(head *ListNode) *ListNode {
	// 基本情况，如果链表为空或只有一个节点，直接返回head
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := reverseListRe(head.Next)
	head.Next.Next = head // * 这一步最重要，head之后的链表都翻转了，所以 head.Next 已经被移动到nextHead的末尾了，这一行就是把 head 移动到nextHead的末尾。
	head.Next = nil
	return nextHead
}

// *****************************************************  辅助函数  ******************************************************
// 打印链表的方法
func printLinnkedList(head *ListNode) {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}
```