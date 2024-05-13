# 700. Search in a Binary Search Tree

Easy

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

Example 1:
![tree1](https://assets.leetcode.com/uploads/2021/01/12/tree1.jpg)
> Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]

Example 2:
![tree2](https://assets.leetcode.com/uploads/2021/01/12/tree2.jpg)
> Input: root = [4,2,7,1,3], val = 5
Output: []
 

Constraints:
> The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 107
root is a binary search tree.
1 <= val <= 107

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

// 1. My solution
func searchBST(root *TreeNode, val int) *TreeNode {
	current := root

	for current != nil {
		if val < current.Val {
			current = current.Left
		} else if val > current.Val {
			current = current.Right
		} else {
			return current
		}
	}

	return nil
}

// 2. Solution from leetcode 递归
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)

}
```