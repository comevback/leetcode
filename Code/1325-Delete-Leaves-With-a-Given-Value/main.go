package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	printTree(tree)
	removeLeafNodes(tree, 2)
	printTree(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =====================================================================================================================
// 1. 双重指针
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	dfs(&root, target) // 传入指针的地址，这样可以修改指针的值
	return root
}

// dfs函数，接收一个指向指针的指针，以及目标值
func dfs(root **TreeNode, target int) {
	if (*root).Left != nil {
		dfs(&(*root).Left, target)
	}

	if (*root).Right != nil {
		dfs(&(*root).Right, target)
	}
	// 如果当前节点是叶子节点，且值等于目标值，将当前节点置为nil
	if (*root).Val == target && (*root).Left == nil && (*root).Right == nil {
		*root = nil
	}
}

// =====================================================================================================================
// 2. 直接赋值，from leetcode
func removeLeafNodes1(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}

	// 因为removeLeafNodes函数返回的是 *TreeNode，所以可以在此直接把指针赋值为nil
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}

	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Val == target && root.Left == nil && root.Right == nil {
		root = nil
		return root
	}
	return root
}

// 打印树方法
func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
