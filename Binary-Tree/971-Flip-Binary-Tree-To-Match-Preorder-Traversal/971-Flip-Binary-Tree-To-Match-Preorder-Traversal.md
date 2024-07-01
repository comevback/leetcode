# 971. Flip Binary Tree To Match Preorder Traversal
Solved
Medium
Topics
Companies
You are given the root of a binary tree with n nodes, where each node is uniquely assigned a value from 1 to n. You are also given a sequence of n values voyage, which is the desired pre-order traversal of the binary tree.

Any node in the binary tree can be flipped by swapping its left and right subtrees. For example, flipping node 1 will have the following effect:

Flip the smallest number of nodes so that the pre-order traversal of the tree matches voyage.

Return a list of the values of all flipped nodes. You may return the answer in any order. If it is impossible to flip the nodes in the tree to make the pre-order traversal match voyage, return the list [-1].

Example 1:

Input: root = [1,2], voyage = [2,1]
Output: [-1]
Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.
Example 2:

Input: root = [1,2,3], voyage = [1,3,2]
Output: [1]
Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.
Example 3:

Input: root = [1,2,3], voyage = [1,2,3]
Output: []
Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.

Constraints:

The number of nodes in the tree is n.
n == voyage.length
1 <= n <= 100
1 <= Node.val, voyage[i] <= n
All the values in the tree are unique.
All the values in voyage are unique.

---

# Code
```go
package main

import "fmt"

func main() {
	// 创建示例二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	// 定义遍历序列
	voyage := []int{2, 1}
	// 检查是否可以通过翻转使遍历序列匹配
	res := flipMatchVoyage(root, voyage)
	fmt.Println(res) // 输出结果
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var possible bool // 标志是否可能匹配遍历序列

// flipMatchVoyage 返回可以匹配遍历序列的翻转节点值，如果不能匹配返回 [-1]
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	possible = true
	index := 0
	reversed := []int{}

	// 前序遍历二叉树
	preOrder(root, voyage, &index, &reversed)
	if possible == false {
		return []int{-1} // 如果无法匹配，返回 [-1]
	} else {
		return reversed // 返回翻转的节点值
	}
}

// preOrder 前序遍历二叉树，同时检查并记录需要翻转的节点
func preOrder(root *TreeNode, voyage []int, index *int, reversed *[]int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	if root.Val != voyage[*index] {
		possible = false // 当前节点值不匹配遍历序列，设置 possible 为 false
		return
	}

	// 检查是否需要翻转，翻转条件：左子节点存在且值不等于遍历序列的下一个值
	if root.Left != nil && root.Left.Val != voyage[*index+1] {
		reverse(root)                           // 翻转左右子节点
		*reversed = append(*reversed, root.Val) // 记录翻转的节点值
	}

	// 检查左右子节点是否都无法匹配遍历序列
	if root.Left != nil && root.Left.Val != voyage[*index+1] && root.Right != nil && root.Right.Val != voyage[*index+1] {
		possible = false // 无法匹配，设置 possible 为 false
	}

	*index += 1 // 增加遍历序列的索引

	// 递归前序遍历左子树和右子树
	preOrder(root.Left, voyage, index, reversed)
	preOrder(root.Right, voyage, index, reversed)
}

// reverse 翻转二叉树的左右子节点
func reverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
}
```