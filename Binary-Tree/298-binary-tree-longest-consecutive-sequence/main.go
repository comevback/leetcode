package main

func main() {
	// 测试用例
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	// 输出最长连续序列长度
	println(longestConsecutive(root)) // 输出应为 3
}

// TreeNode 定义了二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个全局变量 maxLength 用于存储最长连续序列的长度
var maxLength int

// longestConsecutive 函数接受一个二叉树的根节点并返回该二叉树的最长连续序列的长度
func longestConsecutive(root *TreeNode) int {
	// 初始化 maxLength 为 0
	maxLength = 0
	// 从根节点开始进行前序遍历，初始深度为 1
	preOrder(root, 1)
	// 返回最长连续序列的长度
	return maxLength
}

// preOrder 函数进行前序遍历以查找最长连续序列
// root: 当前节点
// depth: 当前连续序列的长度
func preOrder(root *TreeNode, depth int) {
	// 如果当前节点为空，直接返回
	if root == nil {
		return
	}

	// 如果当前深度大于 maxLength，更新 maxLength
	if depth > maxLength {
		maxLength = depth
	}

	// 递归遍历左子树
	if root.Left != nil {
		// 如果左子节点值与当前节点值连续，则增加深度
		if root.Left.Val == root.Val+1 {
			preOrder(root.Left, depth+1)
			// 否则重置深度为 1
		} else {
			preOrder(root.Left, 1)
		}
	}

	// 递归遍历右子树
	if root.Right != nil {
		// 如果右子节点值与当前节点值连续，则增加深度
		if root.Right.Val == root.Val+1 {
			preOrder(root.Right, depth+1)
			// 否则重置深度为 1
		} else {
			preOrder(root.Right, 1)
		}
	}
}
