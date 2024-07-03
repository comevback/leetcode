package main

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储结果
var res []int

// getLonelyNodes 返回二叉树中所有孤独节点的值
// 孤独节点定义为其父节点只有一个子节点的节点
func getLonelyNodes(root *TreeNode) []int {
	res = []int{}
	travel(root)
	return res
}

// travel 递归遍历二叉树，并记录所有孤独节点的值
func travel(root *TreeNode) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果当前节点只有一个右子节点
	if root.Left == nil && root.Right != nil {
		// 将右子节点的值添加到结果中
		res = append(res, root.Right.Val)
		// 递归遍历右子树
		travel(root.Right)
		// 如果当前节点只有一个左子节点
	} else if root.Left != nil && root.Right == nil {
		// 将左子节点的值添加到结果中
		res = append(res, root.Left.Val)
		// 递归遍历左子树
		travel(root.Left)
		// 如果当前节点有两个子节点
	} else {
		// 递归遍历左子树
		travel(root.Left)
		// 递归遍历右子树
		travel(root.Right)
	}
}
