# 116. Populating Next Right Pointers in Each Node

<https://labuladong.online/algo/data-structure/binary-tree-part1/#%E7%AC%AC%E4%BA%8C%E9%A2%98%E3%80%81%E5%A1%AB%E5%85%85%E8%8A%82%E7%82%B9%E7%9A%84%E5%8F%B3%E4%BE%A7%E6%8C%87%E9%92%88>

Medium

You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

 

Example 1:
![sample](https://assets.leetcode.com/uploads/2019/02/14/116_sample.png)
> Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Example 2:
> Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 212 - 1].
-1000 <= Node.val <= 1000
 

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
```