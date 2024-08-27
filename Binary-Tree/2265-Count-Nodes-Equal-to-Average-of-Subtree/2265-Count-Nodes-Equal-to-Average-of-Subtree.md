# 2265. Count Nodes Equal to Average of Subtree

Solved
Medium
Topics
Companies
Hint
Given the root of a binary tree, return the number of nodes where the value of the node is equal to the average of the values in its subtree.

Note:

The average of n elements is the sum of the n elements divided by n and rounded down to the nearest integer.
A subtree of root is a tree consisting of root and all of its descendants.

Example 1:

Input: root = [4,8,5,0,1,null,6]
Output: 5
Explanation:
For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
For the node with value 0: The average of its subtree is 0 / 1 = 0.
For the node with value 1: The average of its subtree is 1 / 1 = 1.
For the node with value 6: The average of its subtree is 6 / 1 = 6.
Example 2:

Input: root = [1]
Output: 1
Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.

Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 1000

---

# Code

```go
package main

func main() {

}

// TreeNode 定义了二叉树的节点结构
type TreeNode struct {
	Val   int       // 当前节点的值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

// averageOfSubtree 计算并返回二叉树中满足条件的子树数量。
// 条件是子树所有节点的值的平均数等于根节点的值。
func averageOfSubtree(root *TreeNode) int {
	_, _, res := getNodes(root) // 调用 getNodes 函数处理整个树
	return res
}

// getNodes 递归计算一个节点及其所有子节点的总和、节点数量和满足条件的子树数量。
// 它返回三个值：节点值的总和，节点数量和满足平均值条件的子树数量。
func getNodes(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0 // 如果节点为空，返回0
	}

	// 递归计算左子树的总和、节点数和满足条件的子树数
	leftSum, leftNums, leftRes := getNodes(root.Left)
	// 递归计算右子树的总和、节点数和满足条件的子树数
	rightSum, rightNums, rightRes := getNodes(root.Right)

	// 计算包含当前节点的总和和节点数
	tempSum := leftSum + rightSum + root.Val
	tempNums := leftNums + rightNums + 1

	// 初始满足条件的子树数为左右子树满足条件的总和
	tempRes := rightRes + leftRes

	// 检查当前子树的平均值是否等于根节点的值
	if tempSum/tempNums == root.Val {
		tempRes += 1 // 如果满足条件，增加结果计数
	}

	// 返回当前子树的总和、节点数和满足条件的子树数
	return tempSum, tempNums, tempRes
}
```
