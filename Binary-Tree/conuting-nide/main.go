package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNode(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return 1 + countNode(root.Left) + countNode(root.Right)
}
