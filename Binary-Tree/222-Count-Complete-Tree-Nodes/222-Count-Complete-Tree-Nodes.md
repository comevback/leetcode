# 222. Count Complete Tree Nodes
Solved
Easy
Topics
Companies
Given the root of a complete binary tree, return the number of the nodes in the tree.

According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Design an algorithm that runs in less than O(n) time complexity.

Example 1:

Input: root = [1,2,3,4,5,6]
Output: 6
Example 2:

Input: root = []
Output: 0
Example 3:

Input: root = [1]
Output: 1

Constraints:

The number of nodes in the tree is in the range [0, 5 * 104].
0 <= Node.val <= 5 * 104
The tree is guaranteed to be complete.

---

# Code
```go
package main

import "math"

func main() {
	// 可以在这里构建二叉树并调用 countNodes1、countNodes2 或 countNodes 函数进行测试
}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归计算节点数量的函数1
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回0
	}

	// 初始值为1，因为根节点不为空
	res := 1
	// 递归计算左子树和右子树的节点数并累加
	res += countNodes1(root.Left) + countNodes1(root.Right)

	return res // 返回总节点数
}

// 全局变量用于记录节点数量
var count int

// 计算节点数量的函数2，使用前序遍历
func countNodes2(root *TreeNode) int {
	count = 0      // 重置全局计数器
	preOrder(root) // 调用前序遍历函数
	return count   // 返回计数结果
}

// 前序遍历函数
func preOrder(root *TreeNode) {
	if root == nil {
		return // 如果当前节点为空，返回
	}

	count += 1           // 节点不为空时，计数加1
	preOrder(root.Left)  // 递归遍历左子树
	preOrder(root.Right) // 递归遍历右子树
}

// ************************************************************************************************************

// 计算完全二叉树的节点数量
func countNodes(root *TreeNode) int {
	// 找到树的高度
	height := findHeight(root)
	// 如果左子树和右子树的高度相等，说明是满二叉树
	if height == findHeightRight(root) {
		// 满二叉树的节点数量公式: 2^height - 1
		return int(math.Pow(float64(2), float64(height))) - 1
	}

	// 否则，递归计算左右子树的节点数量
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// 找到树的高度（沿左子树）
func findHeight(root *TreeNode) int {
	current := root
	height := 0

	// 沿左子树向下遍历，计算高度
	for current != nil {
		height += 1
		current = current.Left
	}

	return height // 返回高度
}

// 找到树的高度（沿右子树）
func findHeightRight(root *TreeNode) int {
	current := root
	height := 0

	// 沿右子树向下遍历，计算高度
	for current != nil {
		height += 1
		current = current.Right
	}

	return height // 返回高度
}
```