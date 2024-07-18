# 109. Convert Sorted List to Binary Search Tree

Medium
Topics
Companies
Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height-balanced binary search tree.

Example 1:

Input: head = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
Example 2:

Input: head = []
Output: []

Constraints:

The number of nodes in head is in the range [0, 2 * 104].
-105 <= Node.val <= 105

---

# Code

```go
package main

import (
	"fmt"
	"math"
)

// ListNode 定义单链表中的一个节点。
type ListNode struct {
	Val  int       // 节点的值
	Next *ListNode // 指向下一个节点的指针
}

// TreeNode 定义二叉树中的一个节点。
type TreeNode struct {
	Val   int       // 节点的值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

func main() {
	// 创建链表 -10 -> -3 -> 0 -> 5 -> 9
	head := &ListNode{Val: -10}
	head.Next = &ListNode{Val: -3}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 9}

	// 将排序链表转换为高度平衡的二叉搜索树
	res := sortedListToBST(head)
	// 打印结果
	fmt.Println(res)
}

// sortedListToBST 接收一个排序的链表并返回一个高度平衡的二叉搜索树。
func sortedListToBST(head *ListNode) *TreeNode {
	length := calLength(head)        // 计算链表长度
	res := build(&head, 0, length-1) // 构建BST
	return res
}

// build 递归函数用于构建BST。
func build(head **ListNode, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	var newNode *TreeNode = &TreeNode{}    // 创建新的树节点
	mid := left + (right-left)/2           // 计算中点
	leftNode := build(head, left, mid-1)   // 递归构建左子树
	newNode.Val = (*head).Val              // 设置当前节点的值
	(*head) = (*head).Next                 // 移动到下一个链表节点
	rightNode := build(head, mid+1, right) // 递归构建右子树
	newNode.Left = leftNode                // 设置左子树
	newNode.Right = rightNode              // 设置右子树
	return newNode
}

// calLength 计算链表的长度。
func calLength(head *ListNode) int {
	current := head
	length := 0
	for current != nil {
		length += 1
		current = current.Next
	}
	return length
}

// calDepth 计算树的深度。
func calDepth(num int) int {
	level := 1
	sum := 1
	index := 1
	for sum < num {
		index *= 2
		sum += index
		level += 1
	}
	return level
}

// calDepth1 另一种计算树深度的方法，使用对数函数。
func calDepth1(num int) int {
	if num <= 0 {
		return 0
	}
	return int(math.Floor(math.Log2(float64(num))) + 1)
}
```
