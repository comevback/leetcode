package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1.遍历，使用了额外空间
var pre []*TreeNode

func flatten1(root *TreeNode) {
	pre = []*TreeNode{}
	if root == nil {
		return
	}

	preOrder(root)
	assign(root)
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	pre = append(pre, root)

	preOrder(root.Left)
	preOrder(root.Right)
}

func assign(root *TreeNode) {
	if len(pre) == 0 {
		return
	}
	root = pre[0]
	root.Left = nil
	if len(pre) == 1 {
		root.Right = nil
	} else {
		root.Right = pre[1]
	}

	pre = pre[1:]
	assign(root.Right)
}

// *******************************************************************************************************************
// 2. 原地，不使用额外空间
func flatten(root *TreeNode) {
	reform(root)
}

func reform(root *TreeNode) {
	if root == nil {
		return
	}
	reform(root.Left)
	reform(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}

	temp.Right = right
}
