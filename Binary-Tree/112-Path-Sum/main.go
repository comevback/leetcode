package main

func main() {
	// 示例主函数，可以在此构建树并测试 hasPathSum 函数
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool // 全局变量用于存储结果

// hasPathSum 检查是否存在从根节点到叶子节点的路径，使得路径上所有节点值之和等于目标和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false // 如果根节点为空，返回 false
	}

	res = false             // 初始化全局结果变量
	travel(root, targetSum) // 调用递归函数检查路径
	return res              // 返回结果
	// return getSum(root, targetSum) // 也可以使用另一种递归方法
}

// travel 递归检查是否存在路径使得节点值之和等于目标和
func travel(root *TreeNode, targetSum int) {
	targetSum -= root.Val // 减去当前节点的值

	// 如果当前节点是叶子节点，检查路径和是否等于目标和
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			res = true // 找到路径，设置结果为 true
		}
	}

	// 递归检查左子树
	if root.Left != nil {
		travel(root.Left, targetSum)
	}

	// 递归检查右子树
	if root.Right != nil {
		travel(root.Right, targetSum)
	}
}

// getSum 递归检查是否存在路径使得节点值之和等于目标和（另一种实现方法）
func getSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false // 如果根节点为空，返回 false
	}

	// 如果当前节点是叶子节点，检查路径和是否等于目标和
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true // 找到路径，返回 true
	}

	// 递归检查左子树和右子树
	return getSum(root.Left, targetSum-root.Val) || getSum(root.Right, targetSum-root.Val)
}
