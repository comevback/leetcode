package main

func main() {
	var head *ListNode
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}

	res := deleteDuplicates1(head)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 单指针
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		if current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

// 2. 双指针
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next

	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val
		} else {
			fast = fast.Next
		}
	}

	slow.Next = nil
	return head
}
