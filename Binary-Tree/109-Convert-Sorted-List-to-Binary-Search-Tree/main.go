package main

import (
	"fmt"
	"math"
)

func main() {
	head := &ListNode{Val: -10}
	head.Next = &ListNode{Val: -3}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 9}

	res := sortedListToBST(head)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	length := calLength(head)
	res := build(&head, 0, length-1)
	return res
}

func build(head **ListNode, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	var newNode *TreeNode = &TreeNode{}
	mid := left + (right-left)/2
	leftNode := build(head, left, mid-1)
	newNode.Val = (*head).Val
	(*head) = (*head).Next
	rightNode := build(head, mid+1, right)
	newNode.Left = leftNode
	newNode.Right = rightNode
	return newNode
}

func calLength(head *ListNode) int {
	current := head
	length := 0
	for current != nil {
		length += 1
		current = current.Next
	}

	return length
}

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

func calDepth1(num int) int {
	if num <= 0 {
		return 0
	}
	return int(math.Floor(math.Log2(float64(num))) + 1)
}
