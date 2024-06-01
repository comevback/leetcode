package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 0}
	head.Next.Next = &ListNode{Val: 0}
	// head.Next.Next.Next = &ListNode{Val: 1}

	res := isPalindrome(head)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy // 快慢指针都指向虚拟头结点

	// 无论链表长度是奇数还是偶数，slow最终都会指向需要被反转的后半部分部分的前一个节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 获取需要被反转的后半部分的头节点，与前面断开
	newHead := slow.Next
	slow.Next = nil

	reversed := reverse(newHead)

	// 如果是奇数个节点，被断开的部分会比前半部分少一个节点，所以用reversed来作为循环条件
	for reversed != nil {
		if reversed.Val != head.Val {
			return false
		}
		reversed = reversed.Next
		head = head.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}
