# 814. Binary Tree Pruning

Solved
Medium
Topics
Companies
Given the root of a binary tree, return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

A subtree of a node node is node plus every node that is a descendant of node.

Example 1:

Input: root = [1,null,0,0,1]
Output: [1,null,0,null,1]
Explanation:
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.
Example 2:

Input: root = [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]
Example 3:

Input: root = [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]

Constraints:

The number of nodes in the tree is in the range [1, 200].
Node.val is either 0 or 1.

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

// pruneTree 函数修剪二叉树，移除所有子树中不包含值为1的节点
func pruneTree(root *TreeNode) *TreeNode {
	return getNode(root)
}

// getNode 函数递归修剪二叉树
func getNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // 如果节点为空，返回 nil
	}

	// 递归修剪左子树和右子树
	root.Left = getNode(root.Left)
	root.Right = getNode(root.Right)

	// 如果当前节点的值为 1，或者其左子树或右子树不为空，则返回当前节点
	// 否则，返回 nil，表示移除当前节点
	if root.Left != nil || root.Right != nil || root.Val == 1 {
		return root
	} else {
		return nil
	}
}
```
