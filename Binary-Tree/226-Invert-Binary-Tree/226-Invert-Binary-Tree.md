# 226. Invert Binary Tree

Easy

Given the root of a binary tree, invert the tree, and return its root.

 

Example 1:
> Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
> Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
> Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

---

# Code
```go
package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	invertTree(root)
	fmt.Println(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre []int

// 1. 遍历思路
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	reversePreOrder(root)
	return root
}

func reversePreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left

	reversePreOrder(root.Right)
	reversePreOrder(root.Left)
}

// 2. 分解思路
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}
```