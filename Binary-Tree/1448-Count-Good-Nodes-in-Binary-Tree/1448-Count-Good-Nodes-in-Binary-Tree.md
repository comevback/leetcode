# 1448. Count Good Nodes in Binary Tree
Solved
Medium
Topics
Companies
Hint
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.

Example 1:

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
Example 2:

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.

Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].

---

# Code
```go
package main

import "math"

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储结果和中间值
var res int

// goodNodes 计算二叉树中所有 "好" 节点的数量
// "好" 节点的定义是该节点在从根到该节点的路径上，节点值大于等于路径上所有节点值
func goodNodes(root *TreeNode) int {
	res = 0
	travel(root, math.MinInt)
	return res
}

// travel 递归遍历二叉树，并计算符合条件的 "好" 节点的数量
func travel(root *TreeNode, max int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果当前节点值大于等于路径上所有节点值，则为 "好" 节点
	if root.Val >= max {
		max = root.Val // 更新路径上最大值
		res += 1       // 计数器加 1
	}

	// 递归遍历左子树，传递更新后的最大值
	travel(root.Left, max)
	// 递归遍历右子树，传递更新后的最大值
	travel(root.Right, max)
}
```