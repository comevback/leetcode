package main

import "fmt"

func main() {
	// head := [1,2,3,4,5,6,7,8]
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	printLinnkedList(head)
	res := oddEvenList(head)
	printLinnkedList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func oddEvenList(head *ListNode) *ListNode {

	// 如果链表为空或者只有一个节点，直接返回nil或者head
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	// odd指针指向奇数节点，even指针指向偶数节点
	odd, even := head, head.Next
	// evenHead指向偶数节点的头节点
	evenHead := even

	// 遍历链表，将奇数节点和偶数节点分开循环
	// 当even.Next != nil时，说明还有至少odd可以指向下一个节点，还可以继续操作
	for even.Next != nil {
		// 如果even.Next.Next != nil，则even也可以指向下一个节点，继续操作
		if even.Next.Next != nil {
			// 把odd和even都指向下一个节点
			odd.Next = even.Next
			even.Next = even.Next.Next
			odd = odd.Next
			even = even.Next
		} else {
			// 如果even.Next.Next == nil，则说明even已经到了最后一个节点，此时odd.Next指向even.Next，odd指向下一个节点
			odd.Next = even.Next
			odd = odd.Next
			even.Next = nil
		}
	}

	// 当odd链表和even链表都遍历完之后，将odd链表的尾节点指向even链表的头节点，链接成一个链表
	odd.Next = evenHead

	// 返回处理好的链表
	return head
}

// 打印链表方法
func printLinnkedList(head *ListNode) {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}
