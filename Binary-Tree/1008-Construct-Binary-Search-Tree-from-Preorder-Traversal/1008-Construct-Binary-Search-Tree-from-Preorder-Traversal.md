# 1008. Construct Binary Search Tree from Preorder Traversal

Solved
Medium
Topics
Companies
Given an array of integers preorder, which represents the preorder traversal of a BST (i.e., binary search tree), construct the tree and return its root.

It is guaranteed that there is always possible to find a binary search tree with the given requirements for the given test cases.

A binary search tree is a binary tree where for every node, any descendant of Node.left has a value strictly less than Node.val, and any descendant of Node.right has a value strictly greater than Node.val.

A preorder traversal of a binary tree displays the value of the node first, then traverses Node.left, then traverses Node.right.

Example 1:

Input: preorder = [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
Example 2:

Input: preorder = [1,3]
Output: [1,null,3]

Constraints:

1 <= preorder.length <= 100
1 <= preorder[i] <= 1000
All the values of preorder are unique.

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

func bstFromPreorder(preorder []int) *TreeNode {
	// 如果数组为空，返回 nil
	if len(preorder) == 0 {
		return nil
	}

	// 创建当前子树的根节点
	newNode := &TreeNode{Val: preorder[0]}

	// 找到第一个大于根节点值的位置
	index := 1
	for index < len(preorder) && preorder[index] <= preorder[0] {
		index++
	}

	// 递归构建左子树
	// preorder[1:index] 包含所有小于根节点值的元素
	newNode.Left = bstFromPreorder(preorder[1:index])

	// 递归构建右子树
	// preorder[index:] 包含所有大于根节点值的元素
	newNode.Right = bstFromPreorder(preorder[index:])

	// 返回当前子树的根节点
	return newNode
}
```
