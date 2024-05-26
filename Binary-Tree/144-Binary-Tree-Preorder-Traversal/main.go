package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	res := preorderTraversal(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := preOrder(root)
	return result
}

// 递归分解分治
func preOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}

	res = append(res, root.Val)

	res = append(res, preOrder(root.Left)...)
	res = append(res, preOrder(root.Right)...)

	return res
}

// 递归遍历
func preOrder2(root *TreeNode) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	preOrder2(root.Left)
	preOrder2(root.Right)
}

// 迭代法
func preOrder3(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}

	stack = append(stack, root)

	for len(stack) != 0 {
		poped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, poped.Val)

		if poped.Right != nil {
			stack = append(stack, poped.Right)
		}
		if poped.Left != nil {
			stack = append(stack, poped.Left)
		}
	}

	return res
}
