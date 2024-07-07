package main

import "fmt"

func main() {
	// 构建二叉树
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}

	// 计算路径和
	res := pathSum(root, 22)
	fmt.Println(res) // 输出所有满足条件的路径
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int // 全局变量用于存储结果路径

// pathSum 返回所有从根节点到叶子节点的路径，使得路径上所有节点值之和等于目标和
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res = [][]int{} // 初始化结果
	arr := []int{}  // 存储当前路径
	// travel(root, []int{}, 0, targetSum) // 使用 travel 方法
	dfs(root, &arr, 0, targetSum) // 使用 dfs 方法
	return res
}

// travel 递归检查路径并存储满足条件的路径（分解方法）
func travel(root *TreeNode, arr []int, sum int, targetSum int) {
	sum += root.Val             // 更新当前路径和
	arr = append(arr, root.Val) // 将当前节点值添加到路径中

	// 如果当前节点是叶子节点，并且路径和等于目标和，存储路径
	if root.Left == nil && root.Right == nil && sum == targetSum {
		temp := make([]int, len(arr)) // 创建路径的副本
		copy(temp, arr)               // 拷贝路径
		res = append(res, temp)       // 将路径添加到结果中
	}

	// 递归检查左子树
	if root.Left != nil {
		travel(root.Left, arr, sum, targetSum)
	}
	// 递归检查右子树
	if root.Right != nil {
		travel(root.Right, arr, sum, targetSum)
	}
}

// dfs 递归检查路径并存储满足条件的路径（遍历方法）
func dfs(root *TreeNode, arr *[]int, sum int, targetSum int) {
	sum += root.Val                   // 更新当前路径和
	(*arr) = append((*arr), root.Val) // 将当前节点值添加到路径中

	// 如果当前节点是叶子节点，并且路径和等于目标和，存储路径
	if root.Left == nil && root.Right == nil && sum == targetSum {
		temp := make([]int, len((*arr))) // 创建路径的副本
		copy(temp, (*arr))               // 拷贝路径
		res = append(res, temp)          // 将路径添加到结果中
	}

	// 递归检查左子树
	if root.Left != nil {
		dfs(root.Left, arr, sum, targetSum)
	}

	// 递归检查右子树
	if root.Right != nil {
		dfs(root.Right, arr, sum, targetSum)
	}

	// 回溯到上一个节点
	(*arr) = (*arr)[:len(*arr)-1]
	sum -= root.Val
}
