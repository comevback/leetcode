package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLength int // 全局变量，用于存储最长的Z字形路径长度

// longestZigZag 函数计算二叉树中最长的Z字形路径长度
func longestZigZag(root *TreeNode) int {
	maxLength = 0        // 初始化最长路径长度为0
	zzLength(root)       // 调用辅助函数计算路径长度
	return maxLength - 1 // 返回最长路径长度减1，因为初始值是1
}

// zzLength 函数递归计算从当前节点出发的Z字形路径长度
func zzLength(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0 // 如果节点为空，返回0,0表示没有路径
	}

	var length int
	_, leftR := zzLength(root.Left)   // 递归计算左子树的路径长度
	rightL, _ := zzLength(root.Right) // 递归计算右子树的路径长度

	// 计算当前节点的最长路径长度
	if rightL > leftR {
		length = rightL + 1
	} else {
		length = leftR + 1
	}

	// 更新全局最大路径长度
	if length > maxLength {
		maxLength = length
	}

	// 返回左子树和右子树的路径长度分别加1，表示从当前节点继续向左或向右的路径长度
	return leftR + 1, rightL + 1
}
