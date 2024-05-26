# 114. Flatten Binary Tree to Linked List

<https://labuladong.online/algo/data-structure/binary-tree-part1/#%E7%AC%AC%E4%B8%89%E9%A2%98%E3%80%81%E5%B0%86%E4%BA%8C%E5%8F%89%E6%A0%91%E5%B1%95%E5%BC%80%E4%B8%BA%E9%93%BE%E8%A1%A8>

Medium

Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.
 

Example 1:
![flaten](https://assets.leetcode.com/uploads/2021/01/14/flaten.jpg)
> Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:
> Input: root = []
Output: []

Example 3:
> Input: root = [0]
Output: [0]
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
 

Follow up: Can you flatten the tree in-place (with O(1) extra space)?

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

// 1.遍历，使用了额外空间
var pre []*TreeNode

func flatten1(root *TreeNode) {
	pre = []*TreeNode{}
	if root == nil {
		return
	}

	preOrder(root)
	assign(root)
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	pre = append(pre, root)

	preOrder(root.Left)
	preOrder(root.Right)
}

func assign(root *TreeNode) {
	if len(pre) == 0 {
		return
	}
	root = pre[0]
	root.Left = nil
	if len(pre) == 1 {
		root.Right = nil
	} else {
		root.Right = pre[1]
	}

	pre = pre[1:]
	assign(root.Right)
}

// *******************************************************************************************************************
// 2. 原地，不使用额外空间
func flatten(root *TreeNode) {
	reform(root)
}

func reform(root *TreeNode) {
	if root == nil {
		return
	}
	reform(root.Left)
	reform(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}

	temp.Right = right
}
```