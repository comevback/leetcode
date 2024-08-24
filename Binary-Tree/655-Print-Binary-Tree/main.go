package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1}           // 创建根节点
	root.Left = &TreeNode{Val: 2}       // 添加左子节点
	root.Right = &TreeNode{Val: 3}      // 添加右子节点
	root.Left.Right = &TreeNode{Val: 4} // 在左子节点下添加右子节点

	res := printTree(root) // 打印二叉树
	fmt.Println(res)       // 输出结果矩阵
}

// TreeNode 结构体定义了二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// printTree 函数生成二叉树的可视化矩阵
func printTree(root *TreeNode) [][]string {
	height := getHeight(root)                    // 计算树的高度
	matrix := make([][]string, height)           // 初始化矩阵
	col := int(math.Pow(2, float64(height)) - 1) // 计算矩阵的列数
	for i := range matrix {
		matrix[i] = make([]string, col) // 初始化每行的列
	}

	yInit := (col - 1) / 2         // 计算根节点的初始列位置
	paint(root, 0, yInit, &matrix) // 开始填充矩阵
	return matrix                  // 返回填充好的矩阵
}

// paint 函数使用递归填充矩阵，以可视化树结构
func paint(root *TreeNode, x int, y int, matrix *[][]string) {
	str := strconv.Itoa(root.Val) // 将节点值转换为字符串
	(*matrix)[x][y] = str         // 将节点值放置在矩阵中的正确位置

	dis := int(math.Pow(2, float64(len(*matrix)-x-2))) // 计算下一个节点的水平偏移量
	if root.Left != nil {
		paint(root.Left, x+1, y-dis, matrix) // 递归填充左子树
	}

	if root.Right != nil {
		paint(root.Right, x+1, y+dis, matrix) // 递归填充右子树
	}
}

// getHeight 函数递归计算树的高度
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0 // 空树高度为0
	}

	left := getHeight(root.Left)   // 计算左子树的高度
	right := getHeight(root.Right) // 计算右子树的高度

	return max(left, right) + 1 // 返回最大子树高度加1
}

// max 函数返回两个数中的最大值
func max(num1 int, num2 int) int {
	if num2 > num1 {
		return num2
	} else {
		return num1
	}
}
