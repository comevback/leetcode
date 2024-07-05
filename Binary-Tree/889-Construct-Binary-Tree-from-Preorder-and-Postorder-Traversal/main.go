package main

import "fmt"

func main() {
	root := constructFromPrePost_75([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(preorder) == 0 {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	lefthead := preorder[1]

	var divide int
	for i := 1; i < len(postorder); i++ {
		if postorder[i] == lefthead {
			divide = i
			break
		}
	}

	leftPostOrder := postorder[:divide+1]
	rightPostOrder := postorder[divide+1 : len(postorder)-1]

	leftPreOrder := preorder[1 : 1+len(leftPostOrder)]
	rightPreOrder := preorder[1+len(leftPostOrder):]

	root.Left = constructFromPrePost(leftPreOrder, leftPostOrder)
	root.Right = constructFromPrePost(rightPreOrder, rightPostOrder)

	return root
}

// **************************************************************************************************************
// review 7.5
func constructFromPrePost_75(preorder []int, postorder []int) *TreeNode {
	return build_75(preorder, postorder)
}

// build_75 递归地构建二叉树
func build_75(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// 创建新的树节点
	newNode := &TreeNode{
		Val: preorder[0],
	}

	// 如果只有一个节点，直接返回
	if len(postorder) == 1 {
		return newNode
	}

	leftPreOrder, leftPostOrder, rightPreOrder, rightPostOrder := []int{}, []int{}, []int{}, []int{}

	// 查找左子树的前序和后序遍历数组
	for i, v := range preorder {
		if v == postorder[len(postorder)-2] {
			leftPreOrder = preorder[1:i]
			rightPreOrder = preorder[i:]
			break
		}
	}

	// 分割右子树的后序遍历数组
	rightPostOrder = postorder[len(postorder)-1-len(rightPreOrder) : len(postorder)-1]
	// 分割左子树的后序遍历数组
	leftPostOrder = postorder[:len(postorder)-1-len(rightPreOrder)]

	// 递归构建左子树
	newNode.Left = build_75(leftPreOrder, leftPostOrder)
	// 递归构建右子树
	newNode.Right = build_75(rightPreOrder, rightPostOrder)

	return newNode
}
