# 538. Convert BST to Greater Tree
Solved
Medium
Topics
Companies
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

The number of nodes in the tree is in the range [0, 104].
-104 <= Node.val <= 104
All the values in the tree are unique.
root is guaranteed to be a valid binary search tree.
 

Note: This question is the same as 1038: https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

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

// convertBST 将二叉搜索树 (BST) 转换为较大和树 (GST)
func convertBST(root *TreeNode) *TreeNode {
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
```