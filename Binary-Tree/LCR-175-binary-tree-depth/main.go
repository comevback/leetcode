package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDepth int

func calculateDepth(root *TreeNode) int {
	//遍历+外部变量
	if root == nil {
		return 0
	}
	iteraDepth(root, 0)
	return maxDepth

	// 分治
	// res := getDepth(root)
	// return res
}

// 1.分治思路
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)

	return 1 + max(left, right)
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// 2.遍历思路
func iteraDepth(root *TreeNode, depth int) {
	depth += 1
	if depth > maxDepth {
		maxDepth = depth
	}
	if root.Left != nil {
		iteraDepth(root.Left, depth)
	}
	if root.Right != nil {
		iteraDepth(root.Right, depth)
	}
}
