package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	travel(root)
	return res
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}

	travel(root.Left)
	travel(root.Right)
	res = append(res, root.Val)
}
