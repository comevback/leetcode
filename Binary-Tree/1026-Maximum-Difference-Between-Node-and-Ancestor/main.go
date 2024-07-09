package main

import "fmt"

func main() {
	// 创建一个测试用例的二叉树
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 19}
	root.Left.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 14}
	root.Right.Right = &TreeNode{Val: 10}
	root.Right.Right.Left = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Right.Left = &TreeNode{Val: 8}
	root.Right.Right.Right.Left.Left = &TreeNode{Val: 13}
	root.Right.Right.Right.Left.Left.Right = &TreeNode{Val: 15}

	// 计算并输出二叉树中节点与其祖先之间的最大差异
	res := maxAncestorDiff(root)
	fmt.Println(res) // 应输出 19
}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int // 全局变量，用于存储最大祖先差异

// maxAncestorDiff 函数计算二叉树中任意节点与其祖先之间的最大值差
func maxAncestorDiff(root *TreeNode) int {
	res = 0          // 初始化结果为 0
	getMaxDiff(root) // 调用递归函数计算最大差异
	return res       // 返回最终的最大差异
}

// getMaxDiff 函数递归遍历二叉树，计算每个节点与其祖先之间的最大值差，并更新全局变量 res
func getMaxDiff(root *TreeNode) (int, int) {
	var minVal, maxVal int                       // 定义当前子树的最小值和最大值
	var leftMin, leftMax, rightMin, rightMax int // 定义左子树和右子树的最小值和最大值

	// 递归计算左子树的最小值和最大值
	if root.Left != nil {
		leftMin, leftMax = getMaxDiff(root.Left)
	} else {
		leftMin, leftMax = root.Val, root.Val // 如果左子树为空，使用当前节点的值作为最小值和最大值
	}

	// 递归计算右子树的最小值和最大值
	if root.Right != nil {
		rightMin, rightMax = getMaxDiff(root.Right)
	} else {
		rightMin, rightMax = root.Val, root.Val // 如果右子树为空，使用当前节点的值作为最小值和最大值
	}

	// 更新全局变量 res，计算当前节点与左右子树的最小值和最大值之间的差异
	if abs(root.Val-leftMin) > res {
		res = abs(root.Val - leftMin)
	}
	if abs(leftMax-root.Val) > res {
		res = abs(leftMax - root.Val)
	}
	if abs(root.Val-rightMin) > res {
		res = abs(root.Val - rightMin)
	}
	if abs(rightMax-root.Val) > res {
		res = abs(rightMax - root.Val)
	}

	// 计算当前子树的最小值和最大值
	minVal = min(min(leftMin, rightMin), root.Val)
	maxVal = max(max(leftMax, rightMax), root.Val)

	return minVal, maxVal // 返回当前子树的最小值和最大值
}

// min 函数返回两个整数中的较小值
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// max 函数返回两个整数中的较大值
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// abs 函数计算一个整数的绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
