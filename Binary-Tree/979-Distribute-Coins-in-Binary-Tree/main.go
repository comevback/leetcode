package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// totalMoves 记录总的移动次数
var totalMoves int

// distributeCoins 函数计算使所有节点的硬币分布均匀所需的最小移动次数
func distributeCoins(root *TreeNode) int {
	totalMoves = 0
	travel(root)
	return totalMoves
}

// travel 函数进行后序遍历，返回当前子树中多余或缺少的硬币数量
func travel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 递归遍历左子树，获取左子树中多余或缺少的硬币数量
	left := travel(root.Left)
	// 递归遍历右子树，获取右子树中多余或缺少的硬币数量
	right := travel(root.Right)

	// 计算当前节点的多余或缺少的硬币数量
	// root.Val 表示当前节点的硬币数量
	// left 表示左子树的多余或缺少的硬币数量
	// right 表示右子树的多余或缺少的硬币数量
	// 每个节点需要1个硬币，所以减去1
	neededMoves := (root.Val + left + right) - 1

	// 累加当前节点的移动次数
	totalMoves += abs(neededMoves)

	// 返回当前子树的多余或缺少的硬币数量
	return neededMoves
}

// abs 函数返回绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
