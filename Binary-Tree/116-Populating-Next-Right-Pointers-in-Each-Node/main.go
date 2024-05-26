package main

import "fmt"

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	res := connect(root)
	fmt.Println(res)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 1.遍历思路 层序遍历
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) != 0 {
		temp := []*Node{}
		for i := 0; i < len(queue); i++ {
			if i == len(queue)-1 {
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		// queue = make([]*Node, len(temp))
		// copy(queue, temp)
		queue = temp
	}
	return root
}

// 2.DFS思路
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	link(root.Left, root.Right)
	return root
}

func link(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2

	link(node1.Left, node1.Right)
	link(node2.Left, node2.Right)

	link(node1.Right, node2.Left)
}

// 3. from leetcode fastest
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}

	connect(root.Left)
	connect(root.Right)

	return root
}
