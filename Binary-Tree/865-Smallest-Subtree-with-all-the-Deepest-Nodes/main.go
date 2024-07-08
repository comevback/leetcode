package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// subtreeWithAllDeepest 函数找到包含所有最深节点的最小子树
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, res := getDepth(root) // 调用 getDepth 函数计算深度并找到对应的子树
	return res               // 返回包含所有最深节点的子树
}

// getDepth 函数递归计算节点深度，并找到包含所有最深节点的最小子树
func getDepth(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil // 如果当前节点为空，返回深度 0 和 nil
	}

	// 递归计算左子树和右子树的深度，并获取对应的子树
	leftDepth, left := getDepth(root.Left)
	rightDepth, right := getDepth(root.Right)

	// 计算当前节点的深度
	depth := max(leftDepth, rightDepth) + 1
	var returnTree *TreeNode

	// 根据左右子树的深度，确定包含所有最深节点的最小子树
	if leftDepth > rightDepth {
		returnTree = left // 左子树更深，返回左子树
	} else if leftDepth < rightDepth {
		returnTree = right // 右子树更深，返回右子树
	} else {
		returnTree = root // 左右子树深度相同，返回当前节点
	}
	return depth, returnTree // 返回当前节点的深度和包含所有最深节点的子树
}

// max 函数返回两个整数中的较大值
func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	} else {
		return num1
	}
}
