# 938. Range Sum of BST

Solved
Easy
Topics
Companies
Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].

Example 1:

Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.

Constraints:

The number of nodes in the tree is in the range [1, 2 * 104].
1 <= Node.val <= 105
1 <= low <= high <= 105
All Node.val are unique.

---

# Code

```go
package main

func main() {

}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int // 全局变量，用于累加符合条件的节点值

// rangeSumBST 函数返回给定范围 [low, high] 内的所有节点值的和
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum = 0                 // 重置 sum 为 0
	travel(root, low, high) // 调用递归函数 travel 进行遍历和求和
	return sum              // 返回最终的和
	// 也可以使用 getSum 函数来替代 travel 函数
	// return getSum(root, low, high)
}

// travel 函数递归遍历二叉树，并累加符合条件的节点值
func travel(root *TreeNode, low int, high int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	if root.Val <= high && root.Val >= low {
		sum += root.Val // 如果当前节点的值在范围内，累加到 sum
	}

	travel(root.Left, low, high)  // 递归遍历左子树
	travel(root.Right, low, high) // 递归遍历右子树
}

// getSum 函数递归遍历二叉树，并返回范围内的节点值的和
func getSum(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	res := 0 // 局部变量，用于存储当前子树的和

	if root.Val <= high && root.Val >= low {
		res += root.Val // 如果当前节点的值在范围内，累加到 res
	}

	res += getSum(root.Left, low, high) + getSum(root.Right, low, high) // 递归计算左右子树的和，并累加到 res
	return res                                                          // 返回当前子树的和
}
```
