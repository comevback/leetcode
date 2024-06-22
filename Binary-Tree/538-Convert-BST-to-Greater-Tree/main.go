package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// convertBST 将二叉搜索树 (BST) 转换为较大和树 (GST)
func convertBST(root *TreeNode) *TreeNode {
	var sum int = 0
	reverseOrder(root, &sum)
	// printRoot(root)
	return root
}

// reverseOrder 以逆中序遍历的方式遍历树
func reverseOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 先遍历右子树
	reverseOrder(root.Right, sum)
	// 临时保存当前节点的值
	temp := root.Val
	// 将当前节点的值加上累积和
	root.Val += *sum
	// 更新累积和
	*sum += temp
	// 再遍历左子树
	reverseOrder(root.Left, sum)
}
