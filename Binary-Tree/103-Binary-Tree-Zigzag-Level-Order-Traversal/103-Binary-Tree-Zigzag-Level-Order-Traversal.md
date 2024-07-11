# 103. Binary Tree Zigzag Level Order Traversal

Solved
Medium
Topics
Companies
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100

---

# Code

```go
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// zigzagLevelOrder 函数按 Z 字形顺序遍历二叉树
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 如果根节点为空，返回空的二维数组
	if root == nil {
		return [][]int{}
	}
	// 初始化结果数组和遍历方向标志
	res := [][]int{}
	leftToRight := true
	// 使用切片作为队列进行层序遍历
	queue := []*TreeNode{}
	queue = append(queue, root)

	// 开始层序遍历
	for len(queue) != 0 {
		// 临时数组用于存储当前层的节点值
		temp := []int{}
		// 获取当前层的节点数
		length := len(queue)

		// 根据遍历方向填充临时数组
		if leftToRight {
			for i := 0; i < length; i++ {
				temp = append(temp, queue[i].Val)
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				temp = append(temp, queue[i].Val)
			}
		}

		// 将当前层的左右子节点加入队列
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			// if leftToRight {
			// 	temp = append(temp, queue[i].Val)
			// } else {
			// 	temp = append([]int{queue[i].Val}, temp...)
			// }
		}

		// 移除已处理的当前层节点
		queue = queue[length:]
		// 将当前层的节点值加入结果数组
		res = append(res, temp)
		// 切换遍历方向
		leftToRight = !leftToRight
	}

	return res
}
```
