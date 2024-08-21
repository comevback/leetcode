# 107. Binary Tree Level Order Traversal II

Solved
Medium
Topics
Companies
Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

---

# Code

```go
package main

import "fmt"

func main() {
	// 创建一棵二叉树
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// 调用levelOrderBottom函数，获取自底向上的层次遍历结果
	res := levelOrderBottom(root)
	// 输出结果
	for _, v := range res {
		for _, val := range v {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

// TreeNode 代表二叉树的节点
type TreeNode struct {
	Val   int       // 节点的值
	Left  *TreeNode // 左子节点的指针
	Right *TreeNode // 右子节点的指针
}

// levelOrderBottom 进行二叉树的层次遍历并返回自底向上的结果
func levelOrderBottom(root *TreeNode) [][]int {
	// 如果根节点为空，直接返回空的二维数组
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}            // 存储最终的层次遍历结果
	queue := []*TreeNode{}      // 队列用于广度优先遍历
	queue = append(queue, root) // 将根节点加入队列

	// 当队列不为空时，继续遍历
	for len(queue) != 0 {
		length := len(queue) // 当前层级节点的数量
		temp := []int{}      // 存储当前层级的节点值
		// 遍历当前层的所有节点
		for i := 0; i < length; i++ {
			// 将左子节点加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 将右子节点加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 将当前节点的值加入临时数组
			temp = append(temp, queue[i].Val)
		}
		// 将当前层的节点值数组加入最终结果中
		res = append(res, temp)
		// 移除当前层的节点
		queue = queue[length:]
	}

	// 反转最终的层次遍历结果，使其自底向上
	reverse(res)
	return res
}

// reverse 反转二维数组的顺序
func reverse(arr [][]int) {
	// 如果数组长度小于2，直接返回
	if len(arr) < 2 {
		return
	}

	head, tail := 0, len(arr)-1

	// 交换头尾元素，直到遍历完数组
	for head < tail {
		arr[head], arr[tail] = arr[tail], arr[head]
		head += 1
		tail -= 1
	}
}
```
