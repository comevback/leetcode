# 2331. Evaluate Boolean Binary Tree

Solved
Easy
Topics
Companies
Hint
You are given the root of a full binary tree with the following properties:

Leaf nodes have either the value 0 or 1, where 0 represents False and 1 represents True.
Non-leaf nodes have either the value 2 or 3, where 2 represents the boolean OR and 3 represents the boolean AND.
The evaluation of a node is as follows:

If the node is a leaf node, the evaluation is the value of the node, i.e. True or False.
Otherwise, evaluate the node's two children and apply the boolean operation of its value with the children's evaluations.
Return the boolean result of evaluating the root node.

A full binary tree is a binary tree where each node has either 0 or 2 children.

A leaf node is a node that has zero children.

Example 1:

Input: root = [2,1,3,null,null,0,1]
Output: true
Explanation: The above diagram illustrates the evaluation process.
The AND node evaluates to False AND True = False.
The OR node evaluates to True OR False = True.
The root node evaluates to True, so we return true.
Example 2:

Input: root = [0]
Output: false
Explanation: The root node is a leaf node and it evaluates to false, so we return false.

Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 3
Every node has either 0 or 2 children.
Leaf nodes have a value of 0 or 1.
Non-leaf nodes have a value of 2 or 3.

---

# Code

```go
package main

func main() {

}

// TreeNode 定义了二叉树节点的结构，用于逻辑运算树
type TreeNode struct {
	Val   int       // 节点的值，可以是 0（假），1（真），2（逻辑或），3（逻辑与）
	Left  *TreeNode // 指向左子节点
	Right *TreeNode // 指向右子节点
}

// evaluateTree 对给定的逻辑运算树进行求值。
// 该函数递归地对树进行遍历，根据节点的类型（操作数或运算符）计算最终的布尔值。
func evaluateTree(root *TreeNode) bool {
	// 如果节点是叶子节点，则根据其值直接返回 true 或 false
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false // 节点值为0表示假
		} else if root.Val == 1 {
			return true // 节点值为1表示真
		}
	}

	// 递归地计算左子树和右子树的值
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	// 根据当前节点的运算符，对子树的值进行逻辑运算
	if root.Val == 2 {
		return left || right // 节点值为2表示逻辑或
	} else {
		return left && right // 节点值为3表示逻辑与
	}
}
```
