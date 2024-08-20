# 111. Minimum Depth of Binary Tree

Solved
Easy
Topics
Companies
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 2
Example 2:

Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5

Constraints:

The number of nodes in the tree is in the range [0, 105].
-1000 <= Node.val <= 1000

---

# Code

```go
package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDepth int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth = math.MaxInt
	travel(root, 1)
	return maxDepth
}

func travel(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if depth < maxDepth {
			maxDepth = depth
		}
	} else {
		travel(root.Left, depth+1)
		travel(root.Right, depth+1)
	}
}
```
