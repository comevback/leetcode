# 110. Balanced Binary Tree

Solved
Easy
Topics
Companies
Given a binary tree, determine if it is height-balanced.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: true
Example 2:

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true

Constraints:

The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104

---

# Code

```go
package main

func main() {

}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool // 全局变量，用于存储树是否平衡的结果

// isBalanced 函数判断二叉树是否平衡
func isBalanced(root *TreeNode) bool {
	res = true     // 初始化结果为 true
	getDepth(root) // 调用 getDepth 函数计算深度，并判断平衡性
	return res     // 返回结果
}

// getDepth 函数递归计算节点深度，并判断平衡性
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回深度 0
	}

	// 递归计算左子树和右子树的深度
	left := getDepth(root.Left)
	right := getDepth(root.Right)

	// 计算左右子树的深度差
	diff := left - right
	// 如果深度差大于 1，说明树不平衡
	if diff < -1 || diff > 1 {
		res = false
	}

	// 计算当前节点的深度
	var depth int
	if left < right {
		depth = right
	} else {
		depth = left
	}

	// 当前节点的深度为左右子树深度的最大值加 1
	depth += 1
	return depth
}
```
