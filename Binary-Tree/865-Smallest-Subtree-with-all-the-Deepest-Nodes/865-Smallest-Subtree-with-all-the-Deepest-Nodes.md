# 865. Smallest Subtree with all the Deepest Nodes

Solved
Medium
Topics
Companies
Given the root of a binary tree, the depth of each node is the shortest distance to the root.

Return the smallest subtree such that it contains all the deepest nodes in the original tree.

A node is called the deepest if it has the largest depth possible among any node in the entire tree.

The subtree of a node is a tree consisting of that node, plus the set of all descendants of that node.

Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation: We return the node with value 2, colored in yellow in the diagram.
The nodes coloured in blue are the deepest nodes of the tree.
Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2 is the smallest subtree among them, so we return it.
Example 2:

Input: root = [1]
Output: [1]
Explanation: The root is the deepest node in the tree.
Example 3:

Input: root = [0,1,3,null,2]
Output: [2]
Explanation: The deepest node in the tree is 2, the valid subtrees are the subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.

Constraints:

The number of nodes in the tree will be in the range [1, 500].
0 <= Node.val <= 500
The values of the nodes in the tree are unique.

Note: This question is the same as 1123: https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/

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

// subtreeWithAllDeepest 函数找到包含所有最深节点的最小子树
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, res := getDepth(root) // 调用 getDepth 函数计算深度并找到对应的子树
	return res               // 返回包含所有最深节点的子树
}

// getDepth 函数递归计算节点深度，并找到包含所有最深节点的最小子树
func getDepth(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil // 如果当前节点为空，返回深度 0 和 nil
	}

	// 递归计算左子树和右子树的深度，并获取对应的子树
	leftDepth, left := getDepth(root.Left)
	rightDepth, right := getDepth(root.Right)

	// 计算当前节点的深度
	depth := max(leftDepth, rightDepth) + 1
	var returnTree *TreeNode

	// 根据左右子树的深度，确定包含所有最深节点的最小子树
	if leftDepth > rightDepth {
		returnTree = left // 左子树更深，返回左子树
	} else if leftDepth < rightDepth {
		returnTree = right // 右子树更深，返回右子树
	} else {
		returnTree = root // 左右子树深度相同，返回当前节点
	}
	return depth, returnTree // 返回当前节点的深度和包含所有最深节点的子树
}

// max 函数返回两个整数中的较大值
func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	} else {
		return num1
	}
}
```
