# 530. Minimum Absolute Difference in BST

Solved
Easy
Topics
Companies
Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

Example 1:

Input: root = [4,2,6,1,3]
Output: 1
Example 2:

Input: root = [1,0,48,null,null,12,49]
Output: 1

Constraints:

The number of nodes in the tree is in the range [2, 104].
0 <= Node.val <= 105

Note: This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/

---

# Code

```go
package main

import (
	"math"
)

func main() {
	// 你可以在这里测试 getMinimumDifference 函数
}

// TreeNode 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量，用于存储最小差值和上一个节点的值
var minDiff, last int

// getMinimumDifference 计算二叉搜索树中任意两个不同节点值之间的最小绝对差值
func getMinimumDifference(root *TreeNode) int {
	// 初始化 minDiff 为最大的整数值
	minDiff = math.MaxInt
	// 初始化 last 为最大的整数值
	last = math.MaxInt

	// 中序遍历二叉树
	inOrder(root)
	return minDiff
}

// inOrder 对二叉搜索树进行中序遍历
func inOrder(root *TreeNode) {
	// 如果节点为空，直接返回
	if root == nil {
		return
	}

	// 递归遍历左子树
	inOrder(root.Left)

	// 计算当前节点与上一个节点的差值
	diff := abs(last - root.Val)
	// 更新最小差值
	if diff < minDiff {
		minDiff = diff
	}
	// 更新上一个节点的值
	last = root.Val

	// 递归遍历右子树
	inOrder(root.Right)
}

// abs 计算一个整数的绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
```
