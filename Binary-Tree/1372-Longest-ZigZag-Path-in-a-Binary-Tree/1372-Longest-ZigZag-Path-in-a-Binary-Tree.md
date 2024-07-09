# 1372. Longest ZigZag Path in a Binary Tree

Solved
Medium
Topics
Companies
Hint
You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:

Choose any node in the binary tree and a direction (right or left).
If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
Change the direction from right to left or from left to right.
Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.

Example 1:

Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
Example 2:

Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
Example 3:

Input: root = [1]
Output: 0

Constraints:

The number of nodes in the tree is in the range [1, 5 * 104].
1 <= Node.val <= 100

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

var maxLength int // 全局变量，用于存储最长的Z字形路径长度

// longestZigZag 函数计算二叉树中最长的Z字形路径长度
func longestZigZag(root *TreeNode) int {
	maxLength = 0        // 初始化最长路径长度为0
	zzLength(root)       // 调用辅助函数计算路径长度
	return maxLength - 1 // 返回最长路径长度减1，因为初始值是1
}

// zzLength 函数递归计算从当前节点出发的Z字形路径长度
func zzLength(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0 // 如果节点为空，返回0,0表示没有路径
	}

	var length int
	_, leftR := zzLength(root.Left)   // 递归计算左子树的路径长度
	rightL, _ := zzLength(root.Right) // 递归计算右子树的路径长度

	// 计算当前节点的最长路径长度
	if rightL > leftR {
		length = rightL + 1
	} else {
		length = leftR + 1
	}

	// 更新全局最大路径长度
	if length > maxLength {
		maxLength = length
	}

	// 返回左子树和右子树的路径长度分别加1，表示从当前节点继续向左或向右的路径长度
	return leftR + 1, rightL + 1
}
```
