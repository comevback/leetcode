# 501. Find Mode in Binary Search Tree

Solved
Easy
Topics
Companies
Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in any order.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or equal to the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:

Input: root = [1,null,2,2]
Output: [2]
Example 2:

Input: root = [0]
Output: [0]

Constraints:

The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105

Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

---

# Code

```go
package main

import "math"

func main() {

}

// TreeNode 结构体定义了二叉树节点
type TreeNode struct {
	Val   int       // 节点存储的整数值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

var res []int // 存储结果的切片，记录出现次数最多的值
var max int   // 记录最大的出现次数
var last int  // 上一个访问的节点值
var times int // 当前值出现的次数

// findMode 函数查找给定二叉树中的众数（出现次数最多的元素）
func findMode(root *TreeNode) []int {
	res = []int{}
	last = math.MaxInt // 将 last 初始化为最大整数，避免与树中的任何数相等
	times = 0
	max = 0
	travel(root) // 调用 travel 函数开始中序遍历
	return res   // 返回结果
}

// travel 函数进行中序遍历
func travel(root *TreeNode) {
	if root == nil {
		return // 如果节点为空，返回
	}

	travel(root.Left) // 递归访问左子树

	// 处理当前节点
	if root.Val == last {
		times += 1 // 如果当前节点值等于上一个节点值，增加计数
	} else {
		times = 1 // 如果不等，重置计数并更新 last
		last = root.Val
	}

	// 更新结果
	if times > max {
		max = times
		res = []int{root.Val} // 如果当前次数大于已知最大次数，更新结果为当前节点值
	} else if times == max {
		res = append(res, root.Val) // 如果当前次数等于最大次数，将当前值添加到结果中
	}

	travel(root.Right) // 递归访问右子树
}
```
