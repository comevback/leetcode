# 404. Sum of Left Leaves
Solved
Easy
Topics
Companies
Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
Example 2:

Input: root = [1]
Output: 0

Constraints:

The number of nodes in the tree is in the range [1, 1000].
-1000 <= Node.val <= 1000

---

# Code
```go
package main

func main() {
	// 示例代码可在此处添加，以测试 sumOfLeftLeaves 函数
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

// sumOfLeftLeaves 计算所有左叶子节点的和
func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	travel(root, "") // 从根节点开始遍历，根节点没有父节点方向
	return sum
}

// travel 递归遍历二叉树，累加左叶子节点的值
func travel(root *TreeNode, parent string) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果当前节点是左叶子节点，则累加其值到 sum 中
	if root.Left == nil && root.Right == nil && parent == "left" {
		sum += root.Val
	}

	// 递归遍历左子树，标记当前节点为 "left"
	if root.Left != nil {
		travel(root.Left, "left")
	}

	// 递归遍历右子树，标记当前节点为 "right"
	if root.Right != nil {
		travel(root.Right, "right")
	}
}
```