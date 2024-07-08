# 563. Binary Tree Tilt

Solved
Easy
Topics
Companies
Hint
Given the root of a binary tree, return the sum of every tree node's tilt.

The tilt of a tree node is the absolute difference between the sum of all left subtree node values and all right subtree node values. If a node does not have a left child, then the sum of the left subtree node values is treated as 0. The rule is similar if the node does not have a right child.

Example 1:

Input: root = [1,2,3]
Output: 1
Explanation:
Tilt of node 2 : |0-0| = 0 (no children)
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
Sum of every tilt : 0 + 0 + 1 = 1
Example 2:

Input: root = [4,2,9,3,5,null,7]
Output: 15
Explanation:
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 5 : |0-0| = 0 (no children)
Tilt of node 7 : |0-0| = 0 (no children)
Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15
Example 3:

Input: root = [21,7,14,1,1,2,2,3,3]
Output: 9

Constraints:

The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000

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

var res int // 全局变量，用于存储树的所有节点的坡度之和

// findTilt 函数计算二叉树的坡度之和
func findTilt(root *TreeNode) int {
	res = 0      // 初始化结果为 0
	travel(root) // 遍历二叉树，计算每个节点的坡度并累加到 res
	return res   // 返回最终的坡度之和
}

// travel 函数递归遍历二叉树，计算每个节点的坡度并累加到 res
func travel(root *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	// 递归计算左子树和右子树的和
	left := travel(root.Left)
	right := travel(root.Right)

	// 计算当前子树的和
	sum := left + right + root.Val
	// 计算当前节点的坡度
	diff := abs(left - right)
	// 将当前节点的坡度累加到 res
	res += diff

	// 将当前节点的值设为当前节点的坡度（这步实际在本题中非必要，只是保留原始逻辑）
	root.Val = diff
	return sum // 返回当前子树的和
}

// abs 函数计算一个整数的绝对值
func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
```
