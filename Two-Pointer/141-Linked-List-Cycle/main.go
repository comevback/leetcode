package main

func main() {
	var head *ListNode
	head = &ListNode{Val: 3}

	res := hasCycle(head)
	println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
