package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	// 如果数组为空，返回 nil
	if len(preorder) == 0 {
		return nil
	}

	// 创建当前子树的根节点
	newNode := &TreeNode{Val: preorder[0]}

	// 找到第一个大于根节点值的位置
	index := 1
	for index < len(preorder) && preorder[index] <= preorder[0] {
		index++
	}

	// 递归构建左子树
	// preorder[1:index] 包含所有小于根节点值的元素
	newNode.Left = bstFromPreorder(preorder[1:index])

	// 递归构建右子树
	// preorder[index:] 包含所有大于根节点值的元素
	newNode.Right = bstFromPreorder(preorder[index:])

	// 返回当前子树的根节点
	return newNode
}
