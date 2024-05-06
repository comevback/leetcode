package main

import (
	"fmt"
)

func main() {
	// 创建一个链表
	newLink := NewLinkedList([]int{5, 2, 13, 3, 8})
	// 对链表进行操作
	res := removeNodes1(newLink.head)
	// 打印操作后的链表
	printList(res)

}

// ======================================== 1.双指针 ========================================

// 1.双指针 从前到后遍历
// 核心思想是快指针寻找是否有比当前最大值大的节点，慢指针判断从哪里开始删除
func removeNodes(head *ListNode) *ListNode {

	// 如果链表为空，直接返回
	if head == nil {
		return head
	}

	// 初始化最大值、慢指针、快指针、新链表头
	var maxVal int = head.Val
	var slow, fast *ListNode = head, head
	var newHead *ListNode = head

	// 当快指针不为空且快指针的下一个节点不为空时
	for fast != nil && fast.Next != nil {
		// 如果快指针的下一个节点的值大于快指针的值，就触发检查
		if fast.Next.Val > fast.Val {
			// 如果快指针的下一个节点的值大于当前最大值，就更新最大值和慢指针，新链表头也更新为这个最大的节点
			if fast.Next.Val > maxVal {
				newHead = fast.Next
				maxVal = fast.Next.Val
				slow = fast
			} else if fast.Next.Val > slow.Next.Val && slow != fast {
				// 如果快指针的下一个节点的值大于慢指针的值，但不大于最大值，慢指针回到新链表头，
				// 然后慢指针一直往后找到比快指针的下一个节点小的节点，然后删除中间所有节点，把慢指针直接指向快指针的下一个节点
				// 然后慢指针更新为快指针
				slow = newHead
				for slow.Next.Val >= fast.Next.Val && slow != fast {
					slow = slow.Next
				}
				slow.Next = fast.Next
				slow = fast
			} else {
				// 如果快指针的下一个节点的值小于慢指针的值，慢指针一直往后找到比快指针的下一个节点小的节点，
				// 然后删除中间所有节点，把慢指针直接指向快指针的下一个节点
				// 然后慢指针更新为快指针
				for slow.Next.Val > fast.Next.Val && slow != fast {
					slow = slow.Next
				}
				slow.Next = fast.Next
				slow = fast
			}
		}
		// 快指针一直往后走
		fast = fast.Next
	}

	// 返回新链表头
	return newHead
}

// ======================================== 2. 翻转 ========================================
// 1. 翻转链表
// 首先翻转链表，遍历翻转后的链表，要复合条件，这个节点的值要大于等于前一个节点的值，也就是不断递增
// 如果不符合条件，就删除跳过这个节点
func removeNodes1(head *ListNode) *ListNode {

	// 如果链表为空，直接返回
	head = reverseLink(head)
	printList(head)
	original := head

	for head.Next != nil {
		if head.Next.Val < head.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	newHead := reverseLink(original)
	printList(newHead)
	return newHead
}

// ======================================== 1. 定义和方法 ========================================

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表定义
type LinkedList struct {
	head *ListNode
	tail *ListNode
}

// 链表的构造函数
func NewLinkedList(nums []int) *LinkedList {

	// 如果数组为空，直接返回空链表
	if len(nums) == 0 {
		return nil
	}

	// 定义链表头和当前节点，头节点值为数组第一个元素
	var head *ListNode = &ListNode{Val: nums[0]}
	var current *ListNode = &ListNode{}
	// 初始时头节点的下一个节点为当前节点
	head.Next = current

	// 遍历数组，构造链表，把每个节点的值赋值为数组当前元素的值
	for i := 1; i < len(nums); i++ {
		// 如果是最后一个元素，就直接赋值给当前节点的值，不要新增节点
		if i == len(nums)-1 {
			current.Val = nums[i]
			break
		}
		current.Val = nums[i]
		newNode := &ListNode{}
		current.Next = newNode
		current = current.Next
	}

	// 链表中，头节点就是链表的头，当前节点就是链表的尾
	var linkList *LinkedList = &LinkedList{
		head: head,
		tail: current,
	}

	// 返回链表
	return linkList
}

// 打印链表为数组
func printList(head *ListNode) {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	fmt.Println(arr)
}

// 翻转链表
func reverseLink(head *ListNode) *ListNode {
	// 如果链表为空，直接返回
	if head == nil {
		return head
	}
	// 定义新链表头
	var newHead *ListNode

	// 遍历链表，把每个节点的Next指向新链表头，然后更新新链表头为当前节点
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	// 返回新链表头
	return newHead
}
