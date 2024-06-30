package main

func main() {
	// 示例代码可在此处添加，以测试 sumOfLeftLeaves 函数
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

// sumOfLeftLeaves 计算所有左叶子节点的和
func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	travel(root, "") // 从根节点开始遍历，根节点没有父节点方向
	return sum
}

// travel 递归遍历二叉树，累加左叶子节点的值
func travel(root *TreeNode, parent string) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果当前节点是左叶子节点，则累加其值到 sum 中
	if root.Left == nil && root.Right == nil && parent == "left" {
		sum += root.Val
	}

	// 递归遍历左子树，标记当前节点为 "left"
	if root.Left != nil {
		travel(root.Left, "left")
	}

	// 递归遍历右子树，标记当前节点为 "right"
	if root.Right != nil {
		travel(root.Right, "right")
	}
}
