# 25. Reverse Nodes in k-Group

<https://labuladong.online/algo/data-structure/reverse-nodes-in-k-group/>

Solved
Hard
Topics
Companies
Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

 

Example 1:
![alt text](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)
> Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Example 2:
![alt text](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)
> Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
 

Constraints:

The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000
 

Follow-up: Can you solve the problem in O(1) extra memory space?

---

# Code
```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Test case 1
	// Input: 1->2->3->4->5, k = 2
	// Output: 2->1->4->3->5
	// Explanation: Reverse the first 2 nodes and the last 2 nodes.
	// 1->2->3->4->5 => 2->1->3->4->5 => 2->1->4->3->5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	res := reverseKGroupRe(head, 2)
	printList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ******************************************************  迭代法  **********************************************************
// ** 利用一个index变量来记录已经反转的节点数，当index等于k时，说明这一轮已经反转了k个节点，reversed是用来保存这些已经翻转好的部分
// ** 每次反转k个节点后，将reversed迭代到末尾，然后把新翻转的newHead接到reversed后面
// ** 如果翻转到最后一轮，但是剩余的节点不足k个，那么直接将已经翻转的节点再翻转回来，然后接到reversed后面
// ** 因为最初用saveHead保存reversed的头部，最后返回saveHead.Next
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}

	index := 0
	var newHead *ListNode = nil
	var reversed *ListNode = &ListNode{}
	saveHead := reversed

	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
		index += 1

		// 当反转k个节点后
		if index == k {
			index = 0
			// 将reversed迭代到末尾
			for reversed.Next != nil {
				reversed = reversed.Next
			}
			reversed.Next = newHead // 将新翻转的newHead接到reversed后面
			newHead = nil           // 重置newHead
		}
	}

	// 反转最后不足k个节点的部分
	for reversed.Next != nil {
		reversed = reversed.Next
	}

	lastList := reverse(newHead)
	reversed.Next = lastList

	return saveHead.Next
}

// 反转整个链表（辅助函数）
func reverse(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

// *********************************************************  递归法  *******************************************************
// ** 递归法的思路是先检查链表是否至少有k个节点，如果没有则不需要反转
// ** 递归反转后续的k个节点，然后反转这一轮的前k个节点
// ** 最后返回新的头部
func reverseKGroupRe(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	current := head
	// 检查链表是否至少有k个节点，如果没有则不需要反转
	for i := 1; i < k; i++ {
		current = current.Next
		if current == nil {
			return head
		}
	}

	// 递归反转后续的k个节点
	var postreversed *ListNode
	if current.Next == nil {
		postreversed = nil
	} else {
		postreversed = reverseKGroupRe(current.Next, k)
	}
	current.Next = postreversed

	// 反转这一轮的前k个节点
	reversed, err := reverseN(head, k)
	if err != nil {
		return head
	}
	return reversed
}

// 翻转前N个元素
func reverseN(head *ListNode, N int) (*ListNode, error) {
	if N == 1 {
		return head, nil
	}

	if head.Next == nil {
		return nil, errors.New("n overflow")
	}
	newHead, err := reverseN(head.Next, N-1)
	if err != nil {
		return nil, errors.New("n overflow")
	}
	noReverse := head.Next.Next
	head.Next.Next = head
	head.Next = noReverse

	return newHead, nil
}

// 打印链表的方法
func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
```