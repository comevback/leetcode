package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDepth int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth = math.MaxInt
	travel(root, 1)
	return maxDepth
}

func travel(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if depth < maxDepth {
			maxDepth = depth
		}
	} else {
		travel(root.Left, depth+1)
		travel(root.Right, depth+1)
	}
}
