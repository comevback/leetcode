# 129. Sum Root to Leaf Numbers

Solved
Medium
Topics
Companies
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.

Example 1:

Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
Example 2:

Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.

Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 9
The depth of the tree will not exceed 10.

---

# Code

```go
package main

import "fmt"

// TreeNode 定义一个二叉树节点的结构
type TreeNode struct {
	Val   int       // 节点值
	Left  *TreeNode // 左子节点
	Right *TreeNode // 右子节点
}

// res 用于存储从根到叶的所有路径之和
var res int

// main 函数是程序的入口
func main() {
	// 创建一个简单的二叉树: 1
	//                    / \
	//                   2   3
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	// 计算所有从根到叶的路径代表的数字之和
	res := sumNumbers(root)
	fmt.Println(res) // 输出结果
}

// sumNumbers 接收一个树的根节点，返回所有从根到叶的路径代表的数字之和
func sumNumbers(root *TreeNode) int {
	res = 0
	travel(root, 0)
	return res
}

// travel 是一个递归函数，用于遍历树，并计算从根到当前节点的路径代表的数字
func travel(root *TreeNode, current int) {
	// 首先计算当前节点的路径代表的数字
	current = current*10 + root.Val

	// 如果当前节点是叶节点，则将路径代表的数字加到结果中
	if root.Left == nil && root.Right == nil {
		res += current
	} else if root.Left == nil { // 否则，递归遍历左子树或右子树
		travel(root.Right, current)
	} else if root.Right == nil {
		travel(root.Left, current)
	} else {
		travel(root.Left, current)
		travel(root.Right, current)
	}
}
```
