# 951. Flip Equivalent Binary Trees
Solved
Medium
Topics
Companies
For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.

A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.

Given the roots of two binary trees root1 and root2, return true if the two trees are flip equivalent or false otherwise.

Example 1:

Flipped Trees Diagram
Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
Output: true
Explanation: We flipped at nodes with values 1, 3, and 5.
Example 2:

Input: root1 = [], root2 = []
Output: true
Example 3:

Input: root1 = [], root2 = [1]
Output: false

Constraints:

The number of nodes in each tree is in the range [0, 100].
Each tree will have unique node values in the range [0, 99].

---
# Code
```go
package main

import "fmt"

func main() {
	// 构建第一个二叉树
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 8}
	root1.Right.Left = &TreeNode{Val: 6}

	// 构建第二个二叉树
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 6}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 5}
	root2.Right.Right.Left = &TreeNode{Val: 8}
	root2.Right.Right.Right = &TreeNode{Val: 7}

	// 检查两个二叉树是否翻转等价
	res := flipEquiv(root1, root2)
	fmt.Println(res) // 输出 true 或 false
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// flipEquiv 检查两个二叉树是否翻转等价
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return check(root1, root2)
}

// check 递归检查两个二叉树是否翻转等价
func check(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true // 两棵树都为空，返回 true
	} else if root1 == nil || root2 == nil {
		return false // 一棵树为空，另一棵树不为空，返回 false
	}

	if root1.Val != root2.Val {
		return false // 两棵树的节点值不同，返回 false
	}

	// 检查左子树和右子树是否需要翻转
	if (root1.Left != nil && root2.Left != nil && root1.Left.Val != root2.Left.Val) || (root1.Left == nil && root2.Left != nil) || (root1.Left != nil && root2.Left == nil) {
		flip(root2) // 如果需要翻转，翻转 root2 的子树
	}

	// 递归检查左子树和右子树是否翻转等价
	return check(root1.Left, root2.Left) && check(root1.Right, root2.Right)
}

// flip 翻转二叉树的左右子树
func flip(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
}
```