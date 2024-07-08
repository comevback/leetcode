package main

func main() {

}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int // 全局变量，用于累加符合条件的节点值

// rangeSumBST 函数返回给定范围 [low, high] 内的所有节点值的和
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum = 0                 // 重置 sum 为 0
	travel(root, low, high) // 调用递归函数 travel 进行遍历和求和
	return sum              // 返回最终的和
	// 也可以使用 getSum 函数来替代 travel 函数
	// return getSum(root, low, high)
}

// travel 函数递归遍历二叉树，并累加符合条件的节点值
func travel(root *TreeNode, low int, high int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	if root.Val <= high && root.Val >= low {
		sum += root.Val // 如果当前节点的值在范围内，累加到 sum
	}

	travel(root.Left, low, high)  // 递归遍历左子树
	travel(root.Right, low, high) // 递归遍历右子树
}

// getSum 函数递归遍历二叉树，并返回范围内的节点值的和
func getSum(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	res := 0 // 局部变量，用于存储当前子树的和

	if root.Val <= high && root.Val >= low {
		res += root.Val // 如果当前节点的值在范围内，累加到 res
	}

	res += getSum(root.Left, low, high) + getSum(root.Right, low, high) // 递归计算左右子树的和，并累加到 res
	return res                                                          // 返回当前子树的和
}
