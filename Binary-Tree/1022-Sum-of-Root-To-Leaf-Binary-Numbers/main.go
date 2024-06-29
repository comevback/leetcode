package main

import "strconv"

func main() {
	// 可以在这里进行测试
}

// TreeNode 定义了二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个全局变量 results 用于存储所有从根到叶子的路径的二进制表示
var results []string

// sumRootToLeaf 函数计算所有从根到叶子的二进制路径的和
func sumRootToLeaf(root *TreeNode) int {
	// 初始化 results 切片为空
	results = []string{}
	// 调用 travel 函数遍历树并收集路径
	travel(root, "")
	sum := 0
	// 将所有二进制路径转换为十进制，并计算它们的和
	for _, v := range results {
		decimal, err := strconv.ParseInt(v, 2, 64)
		if err != nil {
			panic(err)
		}
		sum += int(decimal)
	}

	return sum
}

// travel 函数递归遍历二叉树，收集从根到叶子的路径
func travel(root *TreeNode, str string) {
	// 将当前节点的值转换为字符串并追加到路径中
	str += strconv.Itoa(root.Val)

	// 如果当前节点是叶子节点，将路径添加到 results 中
	if root.Left == nil && root.Right == nil {
		results = append(results, str)
		return
	}

	// 递归遍历左子树
	if root.Left != nil {
		travel(root.Left, str)
	}

	// 递归遍历右子树
	if root.Right != nil {
		travel(root.Right, str)
	}
}
