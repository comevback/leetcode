package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// pruneTree 函数修剪二叉树，移除所有子树中不包含值为1的节点
func pruneTree(root *TreeNode) *TreeNode {
	return getNode(root)
}

// getNode 函数递归修剪二叉树
func getNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // 如果节点为空，返回 nil
	}

	// 递归修剪左子树和右子树
	root.Left = getNode(root.Left)
	root.Right = getNode(root.Right)

	// 如果当前节点的值为 1，或者其左子树或右子树不为空，则返回当前节点
	// 否则，返回 nil，表示移除当前节点
	if root.Left != nil || root.Right != nil || root.Val == 1 {
		return root
	} else {
		return nil
	}
}
