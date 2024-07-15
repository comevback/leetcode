# 671. Second Minimum Node In a Binary Tree

Solved
Easy
Topics
Companies
Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes. More formally, the property root.val = min(root.left.val, root.right.val) always holds.

Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.

If no such second minimum value exists, output -1 instead.

Example 1:

Input: root = [2,2,5,null,null,5,7]
Output: 5
Explanation: The smallest value is 2, the second smallest value is 5.
Example 2:

Input: root = [2,2,2]
Output: -1
Explanation: The smallest value is 2, but there isn't any second smallest value.

Constraints:

The number of nodes in the tree is in the range [1, 25].
1 <= Node.val <= 231 - 1
root.val == min(root.left.val, root.right.val) for each internal node of the tree.

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

func findSecondMinimumValue1(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}
	var smallest, second int = math.MaxInt, math.MaxInt
	travel(root, &smallest, &second)
	if second != math.MaxInt {
		return second
	} else {
		return -1
	}
}

// 遍历解法
func travel(root *TreeNode, small, second *int) {
	if root == nil {
		return
	}

	if root.Val <= *small {
		*small = root.Val
	} else if root.Val < *second {
		*second = root.Val
	}

	travel(root.Left, small, second)
	travel(root.Right, small, second)
}

// solution
func findSecondMinimumValue(root *TreeNode) int {
	// 如果左右子节点都为空，说明树中只有一个节点，没有第二小的值
	if root.Left == nil && root.Right == nil {
		return -1
	}

	// 初始化左右子节点的值
	left, right := root.Left.Val, root.Right.Val

	// 如果左子节点的值等于根节点的值，递归寻找左子树中的第二小值
	if root.Val == root.Left.Val {
		left = findSecondMinimumValue(root.Left)
	}

	// 如果右子节点的值等于根节点的值，递归寻找右子树中的第二小值
	if root.Val == root.Right.Val {
		right = findSecondMinimumValue(root.Right)
	}

	// 如果左子树没有找到第二小值，返回右子树的结果
	if left == -1 {
		return right
	}

	// 如果右子树没有找到第二小值，返回左子树的结果
	if right == -1 {
		return left
	}

	// 左右子树都找到第二小值，返回其中较小的一个
	return min(left, right)
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}
```
