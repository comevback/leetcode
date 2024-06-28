package main

func main() {
	// 可以在这里进行测试
}

// TreeNode 定义了二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个全局变量 res 用于存储右视图结果
var res []int

// rightSideView 函数接受一个二叉树的根节点并返回该二叉树的右视图
func rightSideView(root *TreeNode) []int {
	// 初始化 res 切片为空
	res = []int{}
	// 调用 postOrder 函数从根节点开始进行遍历，初始深度为 0
	depth := 0
	postOrder(root, depth)
	// 返回右视图结果
	return res
}

// postOrder 函数进行右子树优先的深度优先搜索
func postOrder(root *TreeNode, depth int) {
	// 如果当前节点为空，直接返回
	if root == nil {
		return
	}

	// 如果当前深度大于或等于 res 的长度，说明该深度的第一个节点还没有被记录
	if depth >= len(res) {
		// 将当前节点的值添加到 res 中
		res = append(res, root.Val)
	}

	// 先递归遍历右子树，深度加 1
	postOrder(root.Right, depth+1)
	// 再递归遍历左子树，深度加 1
	postOrder(root.Left, depth+1)
}
