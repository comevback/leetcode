# 237. Delete Node in a Linked List（删除链表元素）

Medium

There is a singly-linked list head and we want to delete a node node in it.

You are given the node to be deleted node. You will not be given access to the first node of head.

All the values of the linked list are unique, and it is guaranteed that the given node node is not the last node in the linked list.

Delete the given node. Note that by deleting the node, we do not mean removing it from memory. We mean:

The value of the given node should not exist in the linked list.
The number of nodes in the linked list should decrease by one.
All the values before node should be in the same order.
All the values after node should be in the same order.
Custom testing:

For the input, you should provide the entire linked list head and the node to be given node. node should not be the last node of the list and should be an actual node in the list.
We will build the linked list and pass the node to your function.
The output will be the entire list after calling your function.
 

Example 1:
> Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.

Example 2:
> Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.
 

Constraints:
> The number of the nodes in the given list is in the range [2, 1000].
-1000 <= Node.val <= 1000
The value of each node in the list is unique.
The node to be deleted is in the list and is not a tail node.

---

## 思路
把当前节点整体替换为下一个节点，包括：
1. 当前节点的值替换为下一个节点的值
2. 当前节点的下一个节点替换为下下个节点
3. 删除下一个节点（下一个节点的值指向nil）

```go
package main

import (
	"fmt"
)

func main() {
	//var head *ListNode = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: nil}}}}
	var head *ListNode = &ListNode{Val: 4, Next: nil}
	var follow1 *ListNode = &ListNode{Val: 5, Next: nil}
	head.Next = follow1
	var follow2 *ListNode = &ListNode{Val: 1, Next: nil}
	follow1.Next = follow2
	var follow3 *ListNode = &ListNode{Val: 9, Next: nil}
	follow2.Next = follow3
	printNode(head)
	deleteNode(follow1)
	printNode(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 我的解法
func deleteNode(node *ListNode) {
	// 把当前节点的值替换为下一个节点的值
	node.Val = node.Next.Val
	// 然后把下一个节点的next存进temp
	temp := node.Next.Next
	// 下一个节点的next指向nil
	node.Next.Next = nil
	// 当前节点的next指向temp
	node.Next = temp
}

// 打印整个链表
func printNode(node *ListNode) {
	// 用数组存储链表的值
	var arr []int

	// 遍历链表，把每一个值存入数组
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}

	// 打印数组
	fmt.Println(arr)
}
```