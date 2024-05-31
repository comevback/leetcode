package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printLinnkedList(head)
	res := reverseBetween(head, 1, 2)
	printLinnkedList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.对调linkNode的值Val
// 没有真正改变链表的链接结构，只重新分配了值
func reverseList1(head *ListNode) *ListNode {

	// 如果链表为nil或只有一个头元素，直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	// 两个指针，一个用于遍历链表得到每个元素的值，一个用于修改链表里每个元素的值
	var temp1 *ListNode = head
	var temp2 *ListNode = head

	// 用于存储链表的值
	var arr []int

	// 遍历链表，将链表的值存储到arr数组中
	for temp1 != nil {
		arr = append(arr, temp1.Val)
		temp1 = temp1.Next
	}

	// 第二次从头遍历链表，将arr数组的值倒序赋值给链表
	count := len(arr) - 1
	for temp2 != nil {
		temp2.Val = arr[count]
		count -= 1
		temp2 = temp2.Next
	}

	// 返回修改值后的原链表
	return head
}

// =======================================================================================================================
// 2.正确的链表反转
func reverseList(head *ListNode) *ListNode {

	// 如果链表为nil或只有一个头元素，直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	// 设一个新的数组，头元素newHead暂时为nil
	var newHead *ListNode = nil

	// 当head.Next不为nil时，继续循环
	for head != nil {
		// 保存head.Next的值
		next := head.Next
		// head.Next指向newHead
		head.Next = newHead
		// newHead前进为head
		newHead = head
		// head前进为next
		head = next
	}

	// 返回newHead
	return newHead
}

// =======================================================================================================================
// 打印链表方法
func printLinnkedList(head *ListNode) {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}

// =======================================================================================================================
// 递归实现
func reverseListRe(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := reverseListRe(head.Next)
	head.Next.Next = head // ! 这一步最重要，head之后的链表都翻转了,所以 head.Next 已经被移动到nextHead的末尾了，这一行就是把 head 移动到nextHead的末尾。
	head.Next = nil

	return nextHead
}

// =======================================================================================================================
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
