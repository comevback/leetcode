package main

import (
	"math"
)

func main() {
	// 你可以在这里测试 getMinimumDifference 函数
}

// TreeNode 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量，用于存储最小差值和上一个节点的值
var minDiff, last int

// getMinimumDifference 计算二叉搜索树中任意两个不同节点值之间的最小绝对差值
func getMinimumDifference(root *TreeNode) int {
	// 初始化 minDiff 为最大的整数值
	minDiff = math.MaxInt
	// 初始化 last 为最大的整数值
	last = math.MaxInt

	// 中序遍历二叉树
	inOrder(root)
	return minDiff
}

// inOrder 对二叉搜索树进行中序遍历
func inOrder(root *TreeNode) {
	// 如果节点为空，直接返回
	if root == nil {
		return
	}

	// 递归遍历左子树
	inOrder(root.Left)

	// 计算当前节点与上一个节点的差值
	diff := abs(last - root.Val)
	// 更新最小差值
	if diff < minDiff {
		minDiff = diff
	}
	// 更新上一个节点的值
	last = root.Val

	// 递归遍历右子树
	inOrder(root.Right)
}

// abs 计算一个整数的绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
