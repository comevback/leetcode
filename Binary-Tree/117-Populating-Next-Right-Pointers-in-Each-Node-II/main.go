package main

import "fmt"

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
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

// connect 函数连接每一层的节点，并返回根节点
func connect(root *Node) *Node {
	travel(root)
	return root
}

// travel 函数使用广度优先遍历连接每一层的节点
func travel(root *Node) {
	// 定义两个队列，queue 保存当前层的节点，newQueue 保存下一层的节点
	queue := []*Node{}
	newQueue := []*Node{}

	// 将根节点加入队列
	queue = append(queue, root)

	// 当新队列和当前队列都为空时，停止循环
	for {
		if len(newQueue) == 0 && len(queue) == 0 {
			break
		}

		// 遍历当前队列中的所有节点
		for i := 0; i < len(queue); i++ {
			// 如果当前节点有左子节点，将左子节点加入新队列
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}

			// 如果当前节点有右子节点，将右子节点加入新队列
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}

		// 遍历新队列中的所有节点，将每个节点的 Next 指向下一个节点
		for i := 0; i < len(newQueue)-1; i++ {
			newQueue[i].Next = newQueue[i+1]
		}

		// 将新队列赋值给当前队列，并清空新队列
		queue = newQueue
		newQueue = []*Node{}
	}
}

// ==== 恒定空间复杂度 ====
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 初始化当前节点为根节点
	current := root

	for current != nil {
		// 使用一个临时的虚拟节点来连接下一层的节点
		dummy := &Node{}
		tail := dummy

		// 遍历当前层的所有节点
		for current != nil {
			if current.Left != nil {
				tail.Next = current.Left
				tail = tail.Next
			}
			if current.Right != nil {
				tail.Next = current.Right
				tail = tail.Next
			}
			// 移动到当前层的下一个节点
			current = current.Next
		}

		// 移动到下一层的最左边的节点
		current = dummy.Next
	}

	return root
}
