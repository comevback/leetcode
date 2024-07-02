# 993. Cousins in Binary Tree
Solved
Easy
Topics
Companies
Given the root of a binary tree with unique values and the values of two different nodes of the tree x and y, return true if the nodes corresponding to the values x and y in the tree are cousins, or false otherwise.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Note that in a binary tree, the root node is at the depth 0, and children of each depth k node are at the depth k + 1.

Example 1:

Input: root = [1,2,3,4], x = 4, y = 3
Output: false
Example 2:

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true
Example 3:

Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false

Constraints:

The number of nodes in the tree is in the range [2, 100].
1 <= Node.val <= 100
Each node has a unique value.
x != y
x and y are exist in the tree.

---

# Code
```go
package main

import "fmt"

func main() {
	// 构建示例二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	// 检查值为 1 和 3 的节点是否为堂兄弟节点
	res := isCousins(root, 1, 3)
	fmt.Println(res) // 输出结果
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储目标节点信息
var a, b int
var tarDepth int
var tarParent *TreeNode
var res bool

// isCousins 检查两个节点是否为堂兄弟节点
func isCousins(root *TreeNode, x int, y int) bool {
	a, b = x, y
	tarDepth = -1
	tarParent = &TreeNode{}

	// 遍历树以寻找目标节点
	travel(root, nil, 0)
	return res
}

// travel 遍历二叉树以检查节点是否为堂兄弟节点
func travel(root *TreeNode, parent *TreeNode, depth int) {
	if root == nil {
		return
	}

	// 如果找到目标节点之一
	if root.Val == a || root.Val == b {
		if tarDepth == -1 {
			// 记录第一个找到的目标节点的深度和父节点
			tarDepth = depth
			tarParent = parent
		} else {
			// 检查第二个找到的目标节点是否满足堂兄弟条件
			if tarDepth == depth && parent != tarParent {
				res = true
			} else {
				res = false
			}
		}
	}

	// 递归遍历左子树
	if root.Left != nil {
		travel(root.Left, root, depth+1)
	}

	// 递归遍历右子树
	if root.Right != nil {
		travel(root.Right, root, depth+1)
	}
}
```