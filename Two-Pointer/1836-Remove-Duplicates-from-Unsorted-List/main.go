package main

import "fmt"

func main() {
	// 创建链表 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
	head := &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}

	fmt.Println("Original list:")
	printList(head)

	// 删除重复节点
	head = removeDuplicateNodes(head)

	fmt.Println("List after removing duplicates:")
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两次遍历，第一次记录每个元素出现次数，第二次找到出现次数等于1的连接
func removeDuplicateNodes(head *ListNode) *ListNode {
	var Repeated map[int]int = make(map[int]int)
	current := head

	for current != nil {
		Repeated[current.Val] += 1
		current = current.Next
	}

	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	deleted, current := dummy, head

	for current != nil {
		if Repeated[current.Val] == 1 {
			deleted.Next = current
			deleted = deleted.Next
		}
		current = current.Next
	}

	deleted.Next = nil
	return dummy.Next
}
