# 257. Binary Tree Paths

Easy

Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.


Example 1:
> Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Example 2:
> Input: root = [1]
Output: ["1"]
 

Constraints:

The number of nodes in the tree is in the range [1, 100].
-100 <= Node.val <= 100

---

# Code
```go
package main

import "strconv"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	// 遍历思路
	// paths = []string{}
	// if root == nil {
	// 	return []string{}
	// }

	// collectPath(root, "")
	// return paths

	// 分解思路
	res := getPath(root)
	return res
}

// 1. 遍历思路
func collectPath(root *TreeNode, path string) {
	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	}

	path += "->"

	if root.Left != nil {
		collectPath(root.Left, path)
	}

	if root.Right != nil {
		collectPath(root.Right, path)
	}
}

// 2. 分解思路
func getPath(root *TreeNode) []string {
	paths := []string{}

	if root.Left == nil && root.Right == nil {
		str := strconv.Itoa(root.Val)
		paths = append(paths, str)
		return paths
	}

	paths = append(paths, getPath(root.Left)...)
	paths = append(paths, getPath(root.Right)...)

	for i := range paths {
		paths[i] = strconv.Itoa(root.Val) + "->" + paths[i]
	}

	return paths
}
```