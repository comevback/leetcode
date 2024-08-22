# 145. Binary Tree Postorder Traversal

Solved
Easy
Topics
Companies
Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]
Output: [3,2,1]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]

Constraints:

The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?

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

var res []int

func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	travel(root)
	return res
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}

	travel(root.Left)
	travel(root.Right)
	res = append(res, root.Val)
}
```
