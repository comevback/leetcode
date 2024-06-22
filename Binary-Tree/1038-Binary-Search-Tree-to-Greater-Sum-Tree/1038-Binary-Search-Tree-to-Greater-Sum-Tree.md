# 1038. Binary Search Tree to Greater Sum Tree
Solved
Medium
Topics
Companies
Hint
Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.

As a reminder, a binary search tree is a tree that satisfies these constraints:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:

Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
Example 2:

Input: root = [0,null,1]
Output: [1,null,1]
 

Constraints:

The number of nodes in the tree is in the range [1, 100].
0 <= Node.val <= 100
All the values in the tree are unique.
 

Note: This question is the same as 538: https://leetcode.com/problems/convert-bst-to-greater-tree/

---

# Code
```go
package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	res := bstToGst(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bstToGst 将二叉搜索树 (BST) 转换为较大和树 (GST)
func bstToGst(root *TreeNode) *TreeNode {
	var sum int = 0
	reverseOrder(root, &sum)
	// printRoot(root)
	return root
}

// reverseOrder 以逆中序遍历的方式遍历树
func reverseOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 先遍历右子树
	reverseOrder(root.Right, sum)
	// 临时保存当前节点的值
	temp := root.Val
	// 将当前节点的值加上累积和
	root.Val += *sum
	// 更新累积和
	*sum += temp
	// 再遍历左子树
	reverseOrder(root.Left, sum)
}

// func inOrder(root *TreeNode, vals *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	inOrder(root.Left, vals)
// 	*vals = append(*vals, root.Val)
// 	inOrder(root.Right, vals)
// }

// func printRoot(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	printRoot(root.Left)
// 	fmt.Println(root.Val)
// 	printRoot(root.Right)
// }
```