package main

func main() {
	var head *ListNode = &ListNode{1, nil}
	// head.Next = &ListNode{2, nil}
	// head.Next.Next = &ListNode{3, nil}
	// head.Next.Next.Next = &ListNode{4, nil}
	// head.Next.Next.Next.Next = &ListNode{5, nil}
	res := removeNthFromEnd(head, 1)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 一次遍历得到结果，前后两个指针，前指针走了n+1次之后，后指针才动
// 这样当前指针到达尾部的nil时，后指针到达了从后往前的第K个元素的前一个元素，可以删除第K个元素
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	var lead, delete *ListNode = dummy, dummy
	count := n
	for lead != nil {
		if count < 0 {
			delete = delete.Next
		}
		lead = lead.Next
		count -= 1
	}

	delete.Next = delete.Next.Next
	return dummy.Next
}
