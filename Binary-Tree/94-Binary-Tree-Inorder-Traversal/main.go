package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	inOrder(root)
	return res
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	res = append(res, root.Val)
	inOrder(root.Right)
}
