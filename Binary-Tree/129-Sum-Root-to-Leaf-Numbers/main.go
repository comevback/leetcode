package main

import "fmt"

// TreeNode 定义一个二叉树节点的结构
type TreeNode struct {
	Val   int       // 节点值
	Left  *TreeNode // 左子节点
	Right *TreeNode // 右子节点
}

// res 用于存储从根到叶的所有路径之和
var res int

// main 函数是程序的入口
func main() {
	// 创建一个简单的二叉树: 1
	//                    / \
	//                   2   3
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	// 计算所有从根到叶的路径代表的数字之和
	res := sumNumbers(root)
	fmt.Println(res) // 输出结果
}

// sumNumbers 接收一个树的根节点，返回所有从根到叶的路径代表的数字之和
func sumNumbers(root *TreeNode) int {
	res = 0
	travel(root, 0)
	return res
}

// travel 是一个递归函数，用于遍历树，并计算从根到当前节点的路径代表的数字
func travel(root *TreeNode, current int) {
	// 首先计算当前节点的路径代表的数字
	current = current*10 + root.Val

	// 如果当前节点是叶节点，则将路径代表的数字加到结果中
	if root.Left == nil && root.Right == nil {
		res += current
	} else if root.Left == nil { // 否则，递归遍历左子树或右子树
		travel(root.Right, current)
	} else if root.Right == nil {
		travel(root.Left, current)
	} else {
		travel(root.Left, current)
		travel(root.Right, current)
	}
}
