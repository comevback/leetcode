package main

import (
	"fmt"
	"math"
)

func main() {
	// 构建二叉树
	// 树的结构为：5,4,8,11,null,13,4,7,2,null,null,null,1
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	// 计算最大路径和
	res := maxPathSum(root)
	fmt.Println(res) // 输出最大路径和
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int // 全局变量用于存储最大路径和

// maxPathSum 计算二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	max = math.MinInt // 初始化最大路径和为最小整数
	getSum(root)      // 递归计算路径和
	return max        // 返回最大路径和
}

// getSum 递归计算以当前节点为根的子树的最大路径和
func getSum(root *TreeNode) int {
	if root == nil {
		return 0 // 如果节点为空，返回0
	}

	// 递归计算左子树和右子树的路径和
	left := getSum(root.Left)
	right := getSum(root.Right)

	// 计算包含当前节点的路径和，忽略负路径
	var tempMax int
	if left > 0 {
		tempMax += left
	}
	if right > 0 {
		tempMax += right
	}
	tempMax += root.Val

	// 更新全局最大路径和
	if tempMax > max {
		max = tempMax
	}

	// 计算从当前节点出发的最大路径和
	if left > right {
		tempMax = root.Val + left
	} else {
		tempMax = root.Val + right
	}

	// 如果路径和为负，则返回0，否则返回路径和
	if tempMax > 0 {
		return tempMax
	} else {
		return 0
	}
}
