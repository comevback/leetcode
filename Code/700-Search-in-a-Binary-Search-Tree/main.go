package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. My solution
func searchBST(root *TreeNode, val int) *TreeNode {
	current := root

	for current != nil {
		if val < current.Val {
			current = current.Left
		} else if val > current.Val {
			current = current.Right
		} else {
			return current
		}
	}

	return nil
}

// 2. Solution from leetcode 递归
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)

}
