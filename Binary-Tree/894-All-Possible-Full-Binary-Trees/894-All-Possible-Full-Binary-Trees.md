# 894. All Possible Full Binary Trees
Solved
Medium
Topics
Companies
Given an integer n, return a list of all possible full binary trees with n nodes. Each node of each tree in the answer must have Node.val == 0.

Each element of the answer is the root node of one possible tree. You may return the final list of trees in any order.

A full binary tree is a binary tree where each node has exactly 0 or 2 children.

Example 1:

Input: n = 7
Output: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
Example 2:

Input: n = 3
Output: [[0,0,0]]

Constraints:

1 <= n <= 20

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

var memo map[int][]*TreeNode // 定义全局变量 memo 用于缓存结果

func allPossibleFBT(n int) []*TreeNode {
	memo = make(map[int][]*TreeNode) // 初始化 memo
	if n%2 == 0 {
		return []*TreeNode{} // 满二叉树的节点数必须是奇数
	}
	return build(n) // 调用 build 函数构建所有可能的满二叉树
}

// build 递归地构建所有可能的满二叉树
func build(n int) []*TreeNode {
	res := []*TreeNode{} // 存储当前 n 的所有可能的满二叉树

	if n == 1 {
		// 只有一个节点的树
		res = append(res, &TreeNode{Val: 0})
		return res
	}

	// 检查 memo 中是否已存在 n 的结果，如果存在直接返回
	if _, exist := memo[n]; exist {
		return memo[n]
	}

	// 遍历可能的左子树和右子树节点数
	for i := 1; i < n; i += 2 {
		j := n - i - 1 // 计算右子树的节点数

		leftSubTrees := build(i)  // 递归构建左子树
		rightSubTrees := build(j) // 递归构建右子树

		// 组合左子树和右子树，并创建新的树节点
		for _, left := range leftSubTrees {
			for _, right := range rightSubTrees {
				newNode := &TreeNode{
					Val: 0,
				}

				newNode.Left = left
				newNode.Right = right
				res = append(res, newNode)
			}
		}
	}

	memo[n] = res // 缓存当前 n 的结果
	return res
}
```