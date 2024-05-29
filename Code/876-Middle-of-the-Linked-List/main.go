package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node2 // node1 指向 node2
	node2.Next = node3 // node2 指向 node3

	fmt.Println(middleNode(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 我的解决方案
func middleNode(head *ListNode) *ListNode {
	// 定义计数变量count
	var count int = 0
	// 定义一个切片，用来把每个节点装进去，最后知道链表总节点数量之和，用中间数-1作为下标找出中间节点
	var arr []*ListNode

	// 当链表节点不为nil时循环，每次循环把当前节点存入arr切片，count数加一，然后下一个节点变成当前节点
	for head != nil {
		arr = append(arr, head)
		count += 1
		head = head.Next
	}

	// 得到总节点数后，根据题目要求，有两个中间节点时取后一个，所以是count/2 + 1
	var middleNum int = count/2 + 1

	// 通过中间数得到下标为中间数-1，从arr切片中找出中间节点
	return arr[middleNum-1]
}

// 别人的解决方案
func middleNode2(head *ListNode) *ListNode {
	//  1 2 3 4 5 6
	//        |
	//                |
	if head == nil {
		return nil
	}

	// 定义两个节点，一开始快节点比慢节点快一步，此后每次循环，快节点走两步，慢节点走一步，最后当快节点走到尾时，慢节点就是中间节点
	slow := head
	fast := head.Next

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}

	return slow
}
