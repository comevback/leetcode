package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "14(2(3)(1))(6(5))"
	root := str2tree(s)
	printTree(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将字符串转换为二叉树
func str2tree(s string) *TreeNode {
	return dfsBuild(s)
}

// 深度优先构建二叉树
func dfsBuild(s string) *TreeNode {
	if len(s) == 0 {
		return nil // 如果字符串为空，返回nil
	}

	// 找到第一个'('的位置，截取节点值
	numIndex := 0
	for numIndex < len(s) && s[numIndex] != '(' {
		numIndex++
	}
	intByte := s[0:numIndex] // 节点值部分的字符串

	// 将字符串转换为整数
	val, err := strconv.Atoi(intByte)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil // 如果转换出错，返回nil
	}

	// 创建新的树节点
	newNode := &TreeNode{Val: val}

	// 如果只有节点值，没有子树，直接返回节点
	if len(s) == 1 {
		return newNode
	}

	// 处理左子树
	bNum := 0
	index := numIndex
	if index < len(s) && s[index] == '(' {
		bNum++ // 遇到左括号，计数加1
	} else {
		return newNode // 没有子树，直接返回节点
	}

	// 找到左子树字符串的结束位置
	for index < len(s) && bNum != 0 {
		index++
		if index < len(s) && s[index] == '(' {
			bNum++ // 遇到左括号，计数加1
		} else if index < len(s) && s[index] == ')' {
			bNum-- // 遇到右括号，计数减1
		}
	}

	// 提取左子树字符串，并递归构建左子树
	leftStr := s[numIndex+1 : index]
	newNode.Left = dfsBuild(leftStr)

	// 处理右子树
	if index != len(s)-1 {
		// 提取右子树字符串，并递归构建右子树
		rightStr := s[index+2 : len(s)-1]
		newNode.Right = dfsBuild(rightStr)
	}

	return newNode // 返回构建好的树节点
}

// 打印二叉树（前序遍历）
func printTree(root *TreeNode) {
	if root == nil {
		return // 如果节点为空，直接返回
	}
	fmt.Println(root.Val) // 打印当前节点值
	printTree(root.Left)  // 打印左子树
	printTree(root.Right) // 打印右子树
}
