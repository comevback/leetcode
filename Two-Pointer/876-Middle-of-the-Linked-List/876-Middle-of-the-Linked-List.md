# 876. Middle of the Linked List

Easy

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.


Example 1:
![lc-midlist1](https://assets.leetcode.com/uploads/2021/07/23/lc-midlist1.jpg)
> Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
![lc-midlist2](https://assets.leetcode.com/uploads/2021/07/23/lc-midlist2.jpg)
> Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
 

Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100

---

# Code
```go
package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	node1.Next = node2 // node1 指向 node2
	node2.Next = node3 // node2 指向 node3
	node3.Next = node4 // node3 指向 node4
	node4.Next = node5 // node4 指向 node5
	node5.Next = node6 // node5 指向 node6

	fmt.Println(middleNode(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 直接查找
func middleNode1(head *ListNode) *ListNode {
	// 定义计数变量count
	var count int = 0
	// 定义一个切片，用来把每个节点装进去，最后知道链表总节点数量之和，用中间数-1作为下标找出中间节点
	var arr []*ListNode

	// 当链表节点不为nil时循环，每次循环把当前节点存入arr切片，count数加一，然后下一个节点变成当前节点
	for head != nil {
		arr = append(arr, head)
		count += 1
		head = head.Next
	}

	// 得到总节点数后，根据题目要求，有两个中间节点时取后一个，所以是count/2 + 1
	var middleNum int = count/2 + 1

	// 通过中间数得到下标为中间数-1，从arr切片中找出中间节点
	return arr[middleNum-1]
}

// =====================================================================================================================
// 2. 快慢指针
// 设定一个虚拟节点dummy，dummy.Next = head，然后定义两个指针slow和fast，都指向dummy
// fast每次走两步，slow每次走一步，当fast走到尾部nil时，slow走到中间节点
func middleNode(head *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy

	for fast != nil {
		// fast先走一步，然后判断fast是否为nil，不为nil再走一步
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		// slow每次走一步
		slow = slow.Next
	}

	// 返回slow，即为中间节点
	return slow
}
```