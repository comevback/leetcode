104. Maximum Depth of Binary Tree

Easy

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
> Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
> Input: root = [1,null,2]
Output: 2
 

Constraints:
> The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100

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

// 1.我的想法，每个节点的最大深度就是其左右子树的最大深度里更大的那个，所以用分治法，传递一个深度，每次遍历都加一，同时返回当前节点得到的最大深度。
func maxDepth(root *TreeNode) int {
	// 如果节点为空，直接返回0
	if root == nil {
		return 0
	}

	depth := 0 // 设定初始深度为0

	res := inTravel(root, depth)
	return res
}

func inTravel(root *TreeNode, depth int) int {
	depth += 1
	var leftMax, rightMax int = depth, depth
	if root.Left != nil {
		leftMax = inTravel(root.Left, depth)
	}
	if root.Right != nil {
		rightMax = inTravel(root.Right, depth)
	}

	max := max(leftMax, rightMax)
	return max
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// ************************************************************************************
// from leetcode
func maxDepth1(root *TreeNode) int {
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := 1 + dfs(node.Left)
		right := 1 + dfs(node.Right)
		return max(left, right)
	}
	return dfs(root)
}
```