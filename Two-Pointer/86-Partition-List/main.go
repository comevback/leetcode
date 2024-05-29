package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	x := 3
	partition(head, x)
	fmt.Println(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var dummy1 *ListNode = &ListNode{}
	var dummy2 *ListNode = &ListNode{}

	var smaller = dummy1
	var bigger = dummy2
	var current = head

	for current != nil {
		if current.Val < x {
			newNode := &ListNode{Val: current.Val} // 这里不能直接赋值，因为会形成环，所以要重新创建一个节点
			smaller.Next = newNode
			smaller = smaller.Next
		} else {
			newNode := &ListNode{Val: current.Val}
			bigger.Next = newNode
			bigger = bigger.Next
		}
		current = current.Next

		// 如果不每次都用新节点，那么每次都要把 current.Next 置为nil，断开
		// temp := current.Next
		// current.Next = nil
		// current = temp
	}

	smaller.Next = dummy2.Next
	return dummy1.Next
}
