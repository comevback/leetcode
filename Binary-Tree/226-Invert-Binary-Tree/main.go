package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	invertTree(root)
	fmt.Println(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre []int

// 1. 遍历思路
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	reversePreOrder(root)
	return root
}

func reversePreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left

	reversePreOrder(root.Right)
	reversePreOrder(root.Left)
}

// 2. 分解思路
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}
