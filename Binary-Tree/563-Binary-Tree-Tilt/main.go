package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int // 全局变量，用于存储树的所有节点的坡度之和

// findTilt 函数计算二叉树的坡度之和
func findTilt(root *TreeNode) int {
	res = 0      // 初始化结果为 0
	travel(root) // 遍历二叉树，计算每个节点的坡度并累加到 res
	return res   // 返回最终的坡度之和
}

// travel 函数递归遍历二叉树，计算每个节点的坡度并累加到 res
func travel(root *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	// 递归计算左子树和右子树的和
	left := travel(root.Left)
	right := travel(root.Right)

	// 计算当前子树的和
	sum := left + right + root.Val
	// 计算当前节点的坡度
	diff := abs(left - right)
	// 将当前节点的坡度累加到 res
	res += diff

	// 将当前节点的值设为当前节点的坡度（这步实际在本题中非必要，只是保留原始逻辑）
	root.Val = diff
	return sum // 返回当前子树的和
}

// abs 函数计算一个整数的绝对值
func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
