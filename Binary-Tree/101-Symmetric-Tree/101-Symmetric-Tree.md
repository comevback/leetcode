# 101. Symmetric Tree
Solved
Easy
Topics
Companies
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:

Input: root = [1,2,2,3,4,4,3]
Output: true
Example 2:

Input: root = [1,2,2,null,3,null,3]
Output: false

Constraints:

The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100 

Follow up: Could you solve it both recursively and iteratively?

---
# Code
```go
package main

import "fmt"

func main() {
	// 构建初始的二叉树
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: -64}
	root.Right = &TreeNode{Val: -64}
	root.Left.Left = &TreeNode{Val: -6}
	root.Left.Right = &TreeNode{Val: -90}
	root.Right.Left = &TreeNode{Val: -90}
	root.Right.Right = &TreeNode{Val: -6}
	root.Left.Left.Left = &TreeNode{Val: 88}
	root.Left.Left.Right = &TreeNode{Val: -44}
	root.Left.Right.Left = &TreeNode{Val: 68}
	root.Left.Right.Right = &TreeNode{Val: -65}
	root.Right.Left.Left = &TreeNode{Val: -76}
	root.Right.Left.Right = &TreeNode{Val: 68}
	root.Right.Right.Left = &TreeNode{Val: -44}
	root.Right.Right.Right = &TreeNode{Val: 88}

	// 检查树是否对称
	res := isSymmetric(root)
	fmt.Println(res) // 输出 true 或 false
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 判断二叉树是否对称
func isSymmetric(root *TreeNode) bool {
	// 使用递归方法检查对称性
	// return mirror(root.Left, root.Right)
	// 使用迭代方法检查对称性
	return checkBFS(root)
}

// ****************************************************  递归  ******************************************************

// mirror 递归地检查两棵树是否互为镜像
func mirror(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true // 两棵树都为空，返回 true
	} else if tree1 == nil || tree2 == nil {
		return false // 一棵树为空，另一棵树不为空，返回 false
	}

	if tree1.Val != tree2.Val {
		return false // 两棵树的节点值不同，返回 false
	}

	// 递归检查左子树和右子树是否互为镜像
	return mirror(tree1.Left, tree2.Right) && mirror(tree1.Right, tree2.Left)
}

// ****************************************************  迭代  ******************************************************

// checkBFS 迭代地检查二叉树是否对称
func checkBFS(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true // 根节点为空或只有一个节点，返回 true
	}

	// 初始化两个队列，用于广度优先搜索
	queue1 := []*TreeNode{root.Left}
	queue2 := []*TreeNode{root.Right}

	for len(queue1) > 0 && len(queue2) > 0 {
		if len(queue1) != len(queue2) {
			return false // 两个队列的长度不同，返回 false
		}

		left, right := 0, len(queue2)-1
		for left < len(queue1) && right >= 0 {
			// 检查对应节点是否相同
			if (queue1[left] == nil && queue2[right] != nil) || (queue1[left] != nil && queue2[right] == nil) || (queue1[left] != nil && queue2[right] != nil && queue1[left].Val != queue2[right].Val) {
				return false // 节点不相同，返回 false
			}
			left++
			right--
		}

		newQueue1 := []*TreeNode{}
		newQueue2 := []*TreeNode{}

		// 将左子树和右子树的子节点添加到新队列中
		for _, node := range queue1 {
			if node != nil {
				newQueue1 = append(newQueue1, node.Left, node.Right)
			}
		}
		for _, node := range queue2 {
			if node != nil {
				newQueue2 = append(newQueue2, node.Right, node.Left)
			}
		}

		queue1 = newQueue1
		queue2 = newQueue2
	}

	return true
}

// ********************************************  leetcode solution  ************************************************

// isSymmetric1 使用 LeetCode 的解法迭代地检查二叉树是否对称
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true // 根节点为空，返回 true
	}

	// 初始化队列，用于广度优先搜索
	q := []*TreeNode{root.Left, root.Right}

	for len(q) > 0 {
		L, R := q[0], q[1]
		q = q[2:]

		if L == nil && R == nil {
			continue // 两个节点都为空，继续检查下一对节点
		}

		if (L != nil && R == nil) || (L == nil && R != nil) || (L != nil && R != nil && L.Val != R.Val) {
			return false // 节点不相同，返回 false
		}

		// 将子节点按顺序添加到队列中
		q = append(q, L.Left, R.Right, L.Right, R.Left)
	}

	return true
}
```