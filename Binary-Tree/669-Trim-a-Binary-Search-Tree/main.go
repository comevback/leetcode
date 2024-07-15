package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	return trim(root, low, high)
}

// trim 是一个递归函数，用于修剪二叉搜索树
func trim(root *TreeNode, low int, high int) *TreeNode {
	// 基本情况：如果节点为空，返回 nil
	if root == nil {
		return nil
	}

	// 如果当前节点的值大于高值，递归修剪左子树
	if root.Val > high {
		return trim(root.Left, low, high)
	}

	// 如果当前节点的值小于低值，递归修剪右子树
	if root.Val < low {
		return trim(root.Right, low, high)
	}

	// 当前节点值在 [low, high] 范围内，递归修剪左右子树
	root.Left = trim(root.Left, low, high)
	root.Right = trim(root.Right, low, high)

	// 返回修剪后的当前节点
	return root
}
