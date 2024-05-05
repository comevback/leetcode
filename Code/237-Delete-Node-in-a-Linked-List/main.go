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
