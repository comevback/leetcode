package main

import (
	"fmt"
	"math"
)

func main() {
	var head *ListNode = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	// pairSum(head)
	res := pairSum1(head)
	println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 定义数组 O(n)-O(n)
func pairSum(head *ListNode) int {
	// 定义一个数组
	var arr []int

	// 把链表里的每个值按顺序加入到数组中
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	// 定义最大的twin sum为maxSum
	var maxSum int = math.MinInt32
	// 遍历数组，找配对的两个值加起来，如果大于maxSum，则替换为maxSum
	for i := 0; i < len(arr); i++ {
		if arr[i]+arr[len(arr)-1-i] > maxSum {
			maxSum = arr[i] + arr[len(arr)-1-i]
		}
	}

	// 返回最大值maxSum
	return maxSum
}

// ========================================================================================================================

// 2. 翻转链表（后半部分）（双指针）O(n)-O(1)
func pairSum1(head *ListNode) int {

	// 定义两个指针，一个快指针，一个慢指针
	var slow, fast *ListNode = head, head.Next

	// 快指针走两步，慢指针走一步，当快指针走到最后一个元素时，慢指针走到中间元素，此时慢指针的下一个元素就是后半部分的头
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转后半部分链表，返回翻转后的头
	reversedHead := reverseList(slow.Next) // 用两种翻转方法都可以

	// 定义最大的twin sum为maxSum
	var maxSum int = math.MinInt32

	// head前半部分正常，后半部分已经被翻转
	// reverseHead所代表的被翻转的后半部分，长度为原链表一半，所以用reverseHead != nil来判断是否遍历完
	for reversedHead != nil {
		// 如果head.Val+reversedHead.Val大于maxSum，则替换为maxSum
		if head.Val+reversedHead.Val > maxSum {
			maxSum = head.Val + reversedHead.Val
		}
		// head和reversedHead分别往后移动
		head = head.Next
		reversedHead = reversedHead.Next
	}

	// 返回最大值maxSum
	return maxSum
}

// ========================================================================================================================
// Functions

// 翻转链表1
// 详细解释见(206-Reverse-Linked-List)
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil

	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	return newHead
}

// 翻转链表2
// 详细解释见(206-Reverse-Linked-List)
func reverseList2(head *ListNode) *ListNode {

	var arr []int

	temp1 := head
	temp2 := head

	for temp1 != nil {
		arr = append(arr, temp1.Val)
		temp1 = temp1.Next
	}

	for i := 0; i < len(arr); i++ {
		temp2.Val = arr[len(arr)-1-i]
		temp2 = temp2.Next
	}

	return head
}

// 打印链表
func printLink(head *ListNode) {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	fmt.Println(arr)
}
