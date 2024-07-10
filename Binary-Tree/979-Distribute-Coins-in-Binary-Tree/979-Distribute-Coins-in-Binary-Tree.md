# 979. Distribute Coins in Binary Tree

Solved
Medium
Topics
Companies
You are given the root of a binary tree with n nodes where each node in the tree has node.val coins. There are n coins in total throughout the whole tree.

In one move, we may choose two adjacent nodes and move one coin from one node to another. A move may be from parent to child, or from child to parent.

Return the minimum number of moves required to make every node have exactly one coin.

Example 1:

Input: root = [3,0,0]
Output: 2
Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.
Example 2:

Input: root = [0,3,0]
Output: 3
Explanation: From the left child of the root, we move two coins to the root [taking two moves]. Then, we move one coin from the root of the tree to the right child.

Constraints:

The number of nodes in the tree is n.
1 <= n <= 100
0 <= Node.val <= n
The sum of all Node.val is n.

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

// totalMoves 记录总的移动次数
var totalMoves int

// distributeCoins 函数计算使所有节点的硬币分布均匀所需的最小移动次数
func distributeCoins(root *TreeNode) int {
	totalMoves = 0
	travel(root)
	return totalMoves
}

// travel 函数进行后序遍历，返回当前子树中多余或缺少的硬币数量
func travel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 递归遍历左子树，获取左子树中多余或缺少的硬币数量
	left := travel(root.Left)
	// 递归遍历右子树，获取右子树中多余或缺少的硬币数量
	right := travel(root.Right)

	// 计算当前节点的多余或缺少的硬币数量
	// root.Val 表示当前节点的硬币数量
	// left 表示左子树的多余或缺少的硬币数量
	// right 表示右子树的多余或缺少的硬币数量
	// 每个节点需要1个硬币，所以减去1
	neededMoves := (root.Val + left + right) - 1

	// 累加当前节点的移动次数
	totalMoves += abs(neededMoves)

	// 返回当前子树的多余或缺少的硬币数量
	return neededMoves
}

// abs 函数返回绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
```
