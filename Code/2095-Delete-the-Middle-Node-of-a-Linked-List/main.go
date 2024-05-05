package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	printLink(node1)
	res := deleteMiddle(node1)
	printLink(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func deleteMiddle(head *ListNode) *ListNode {

	// 如果只有一个节点，直接返回nil
	if head.Next == nil {
		return nil
	}

	// 定义快慢指针
	var slow, fast *ListNode = head, head

	// 每次循环，快慢指针都向后移动一步，如果快指针的下一个节点为空，说明慢指针指向的是中间节点
	// 如果快指针的下一个节点不为空，说明慢指针指向的是中间节点的前一个节点，快指针继续向后移动一步
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	// 当快指针无法再向后移动时，这时候慢指针指向的就是中间节点
	// 当中间节点的下一个节点不为空时，将中间节点替换为下一个节点，相当于删除了中间节点
	if slow.Next != nil {
		slow.Val = slow.Next.Val
		temp := slow.Next.Next
		slow.Next.Next = nil
		slow.Next = temp
	} else {
		// 当中间节点的下一个节点为空时，说明中间节点是最后一个节点，直接将中间节点置为nil
		head.Next = nil
	}

	// 返回链表的头节点
	return head
}

// 打印链表方法
func printLink(head *ListNode) {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	fmt.Println(arr)
}
