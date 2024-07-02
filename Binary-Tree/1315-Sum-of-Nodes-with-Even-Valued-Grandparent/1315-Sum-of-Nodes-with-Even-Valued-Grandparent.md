# 1315. Sum of Nodes with Even-Valued Grandparent
Solved
Medium
Topics
Companies
Hint
Given the root of a binary tree, return the sum of values of nodes with an even-valued grandparent. If there are no nodes with an even-valued grandparent, return 0.

A grandparent of a node is the parent of its parent if it exists.

Example 1:

Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 18
Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.
Example 2:

Input: root = [1]
Output: 0

Constraints:

The number of nodes in the tree is in the range [1, 104].
1 <= Node.val <= 100

---

# Code
```go
package main

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储结果和中间值
var sum int

// sumEvenGrandparent 计算二叉树中所有值为偶数的祖父节点的子节点值之和
func sumEvenGrandparent(root *TreeNode) int {
	sum = 0
	travel(root, nil, nil)
	return sum
}

// travel 递归遍历二叉树，并计算符合条件的节点值之和
func travel(root *TreeNode, parent *TreeNode, grandParent *TreeNode) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果祖父节点存在且其值为偶数，将当前节点的值加到 sum 中
	if grandParent != nil && grandParent.Val%2 == 0 {
		sum += root.Val
	}

	// 递归遍历左子树，将当前节点作为父节点，当前父节点作为祖父节点传递
	travel(root.Left, root, parent)
	// 递归遍历右子树，将当前节点作为父节点，当前父节点作为祖父节点传递
	travel(root.Right, root, parent)
}
```