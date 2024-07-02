package main

import "math"

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储结果和中间值
var res int

// goodNodes 计算二叉树中所有 "好" 节点的数量
// "好" 节点的定义是该节点在从根到该节点的路径上，节点值大于等于路径上所有节点值
func goodNodes(root *TreeNode) int {
	res = 0
	travel(root, math.MinInt)
	return res
}

// travel 递归遍历二叉树，并计算符合条件的 "好" 节点的数量
func travel(root *TreeNode, max int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果当前节点值大于等于路径上所有节点值，则为 "好" 节点
	if root.Val >= max {
		max = root.Val // 更新路径上最大值
		res += 1       // 计数器加 1
	}

	// 递归遍历左子树，传递更新后的最大值
	travel(root.Left, max)
	// 递归遍历右子树，传递更新后的最大值
	travel(root.Right, max)
}
