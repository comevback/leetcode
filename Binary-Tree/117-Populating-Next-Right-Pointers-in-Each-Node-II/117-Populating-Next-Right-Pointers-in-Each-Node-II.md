# 117. Populating Next Right Pointers in Each Node II

Solved
Medium
Topics
Companies
Given a binary tree

struct Node {
int val;
Node *left;
Node *right;
Node \*next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Example 1:

Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
Example 2:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 6000].
-100 <= Node.val <= 100

Follow-up:

You may only use constant extra space.
The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.

---

# Code

```go
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
```
