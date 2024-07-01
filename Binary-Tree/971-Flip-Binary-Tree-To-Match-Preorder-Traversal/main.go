package main

import "fmt"

func main() {
	// 创建示例二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	// 定义遍历序列
	voyage := []int{2, 1}
	// 检查是否可以通过翻转使遍历序列匹配
	res := flipMatchVoyage(root, voyage)
	fmt.Println(res) // 输出结果
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var possible bool // 标志是否可能匹配遍历序列

// flipMatchVoyage 返回可以匹配遍历序列的翻转节点值，如果不能匹配返回 [-1]
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	possible = true
	index := 0
	reversed := []int{}

	// 前序遍历二叉树
	preOrder(root, voyage, &index, &reversed)
	if possible == false {
		return []int{-1} // 如果无法匹配，返回 [-1]
	} else {
		return reversed // 返回翻转的节点值
	}
}

// preOrder 前序遍历二叉树，同时检查并记录需要翻转的节点
func preOrder(root *TreeNode, voyage []int, index *int, reversed *[]int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	if root.Val != voyage[*index] {
		possible = false // 当前节点值不匹配遍历序列，设置 possible 为 false
		return
	}

	// 检查是否需要翻转，翻转条件：左子节点存在且值不等于遍历序列的下一个值
	if root.Left != nil && root.Left.Val != voyage[*index+1] {
		reverse(root)                           // 翻转左右子节点
		*reversed = append(*reversed, root.Val) // 记录翻转的节点值
	}

	// 检查左右子节点是否都无法匹配遍历序列
	if root.Left != nil && root.Left.Val != voyage[*index+1] && root.Right != nil && root.Right.Val != voyage[*index+1] {
		possible = false // 无法匹配，设置 possible 为 false
	}

	*index += 1 // 增加遍历序列的索引

	// 递归前序遍历左子树和右子树
	preOrder(root.Left, voyage, index, reversed)
	preOrder(root.Right, voyage, index, reversed)
}

// reverse 翻转二叉树的左右子节点
func reverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
}
