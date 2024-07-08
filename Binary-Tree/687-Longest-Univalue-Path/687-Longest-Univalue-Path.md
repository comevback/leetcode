# 687. Longest Univalue Path

Solved
Medium
Topics
Companies
Given the root of a binary tree, return the length of the longest path, where each node in the path has the same value. This path may or may not pass through the root.

The length of the path between two nodes is represented by the number of edges between them.

Example 1:

Input: root = [5,4,5,1,1,null,5]
Output: 2
Explanation: The shown image shows that the longest path of the same value (i.e. 5).
Example 2:

Input: root = [1,4,5,4,4,null,5]
Output: 2
Explanation: The shown image shows that the longest path of the same value (i.e. 4).

Constraints:

The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000
The depth of the tree will not exceed 1000.

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

var res int // 全局变量，用于存储最长相同值路径的长度

// longestUnivaluePath 函数返回二叉树中相同值路径的最长长度
func longestUnivaluePath(root *TreeNode) int {
	res = 0      // 初始化结果为 0
	travel(root) // 遍历二叉树，计算最长相同值路径
	return res   // 返回最终的最长相同值路径长度
}

// travel 函数递归遍历二叉树，计算每个节点的最长相同值路径
func travel(root *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	var max int                 // 局部变量，用于存储当前节点的最长路径
	var length int              // 局部变量，用于存储当前节点的单边最长路径
	left := travel(root.Left)   // 递归计算左子树的最长路径
	right := travel(root.Right) // 递归计算右子树的最长路径

	// 检查左右子树是否与当前节点的值相同，并计算最长路径
	if root.Left != nil && root.Val == root.Left.Val && root.Right != nil && root.Val == root.Right.Val {
		max = left + right + 2 // 如果左右子树都与当前节点值相同，计算两边的总长度
		if left > right {
			length = left + 1 // 选择左子树为最长路径的一边
		} else {
			length = right + 1 // 选择右子树为最长路径的一边
		}
	} else if root.Left != nil && root.Val == root.Left.Val {
		max = left + 1 // 只有左子树与当前节点值相同
		length = left + 1
	} else if root.Right != nil && root.Val == root.Right.Val {
		max = right + 1 // 只有右子树与当前节点值相同
		length = right + 1
	}

	// 更新全局结果 res
	if max > res {
		res = max
	}

	// 如果当前节点值与左右子树的值都不相同，重置 length 为 0
	if (root.Left == nil || root.Val != root.Left.Val) && (root.Right == nil || root.Val != root.Right.Val) {
		length = 0
	}

	return length // 返回当前节点的单边最长路径长度
}
```
