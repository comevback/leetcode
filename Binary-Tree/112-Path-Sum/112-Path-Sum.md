# 112. Path Sum
Solved
Easy
Topics
Companies
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

Example 1:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.
Example 2:

Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.
Example 3:

Input: root = [], targetSum = 0
Output: false
Explanation: Since the tree is empty, there are no root-to-leaf paths.

Constraints:

The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

---
# Code
```go
package main

func main() {
	// 示例主函数，可以在此构建树并测试 hasPathSum 函数
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool // 全局变量用于存储结果

// hasPathSum 检查是否存在从根节点到叶子节点的路径，使得路径上所有节点值之和等于目标和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false // 如果根节点为空，返回 false
	}

	res = false             // 初始化全局结果变量
	travel(root, targetSum) // 调用递归函数检查路径
	return res              // 返回结果
	// return getSum(root, targetSum) // 也可以使用另一种递归方法
}

// travel 递归检查是否存在路径使得节点值之和等于目标和
func travel(root *TreeNode, targetSum int) {
	targetSum -= root.Val // 减去当前节点的值

	// 如果当前节点是叶子节点，检查路径和是否等于目标和
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			res = true // 找到路径，设置结果为 true
		}
	}

	// 递归检查左子树
	if root.Left != nil {
		travel(root.Left, targetSum)
	}

	// 递归检查右子树
	if root.Right != nil {
		travel(root.Right, targetSum)
	}
}

// getSum 递归检查是否存在路径使得节点值之和等于目标和（另一种实现方法）
func getSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false // 如果根节点为空，返回 false
	}

	// 如果当前节点是叶子节点，检查路径和是否等于目标和
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true // 找到路径，返回 true
	}

	// 递归检查左子树和右子树
	return getSum(root.Left, targetSum-root.Val) || getSum(root.Right, targetSum-root.Val)
}
```