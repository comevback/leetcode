# 102. Binary Tree Level Order Traversal

Solved
Medium
Topics
Companies
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
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

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		// 如果树为空，返回空的二维数组
		return [][]int{}
	}

	// 初始化结果集
	res := [][]int{}
	// 初始化队列，将根节点加入队列
	queue := []*TreeNode{root}

	// 层序遍历
	for len(queue) != 0 {
		// 用于存储当前层的节点值
		temp := []int{}
		// 获取当前层的节点数
		length := len(queue)

		// 遍历当前层的每个节点
		for i := 0; i < length; i++ {
			// 将当前节点的左子节点加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 将当前节点的右子节点加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 将当前节点的值加入当前层的结果集
			temp = append(temp, queue[i].Val)
		}

		// 移除当前层的节点
		queue = queue[length:]
		// 将当前层的结果加入总结果集
		res = append(res, temp)
	}

	return res
}
```
