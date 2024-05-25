package main

import "fmt"

func main() {

	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   5,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   3,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// }

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	res := diameterOfBinaryTree(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLength int

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength = 0
	getLeftRightLen(root)
	return maxLength
}

func getLeftRightLen(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var left, right int

	if root.Left != nil {
		left = getLeftRightLen(root.Left) + 1
	}
	if root.Right != nil {
		right = getLeftRightLen(root.Right) + 1
	}

	if right+left > maxLength {
		maxLength = right + left
	}

	length := max(left, right)
	return length
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
