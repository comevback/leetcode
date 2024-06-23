package main

import "fmt"

func main() {
	// 构建二叉搜索树
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}

	// 调用 kthSmallest1 函数找到第 k 小元素
	res := kthSmallest1(root, 2)
	// 打印结果
	fmt.Println(res)
}

// 定义二叉树节点结构
type TreeNode struct {
	Val      int
	Left     *TreeNode
	Right    *TreeNode
	nodeNums int // 用于存储以当前节点为根的子树节点总数
}

var res int // 定义全局变量 res 用于存储第 k 小元素

// 1. 使用中序遍历的方法查找第 k 小元素
func kthSmallest(root *TreeNode, k int) int {
	res = 0
	count := 1
	inOrder(root, &count, k)
	return res
}

// inOrder 函数用于中序遍历二叉搜索树
func inOrder(root *TreeNode, count *int, k int) {
	if root == nil {
		return
	}

	// 递归遍历左子树
	inOrder(root.Left, count, k)

	// 如果计数器等于 k，说明找到第 k 小元素
	if *count == k {
		res = root.Val
	}
	// 即使找到第 k 小元素，也继续递归遍历，以保持简单的递归结构
	// 这有助于在某些情况下提高缓存命中率和减少上下文切换开销

	*count += 1 // 计数器加 1

	// 递归遍历右子树
	inOrder(root.Right, count, k)
}

// ------------------------------------------------------------------------------------------------------------------
// 2. 使用预先计算子树节点数量的方法查找第 k 小元素
func kthSmallest1(root *TreeNode, k int) int {
	res = 0
	countNodes(root)
	findK(root, k)
	return res
}

// countNodes 函数用于计算并存储每个节点的子节点数量
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftNums := countNodes(root.Left)
	rightNums := countNodes(root.Right)

	root.nodeNums = leftNums + rightNums + 1 // 当前节点的子树节点总数
	return root.nodeNums
}

// findK 函数用于查找第 k 小的元素
func findK(root *TreeNode, k int) {
	if root == nil {
		return
	}

	// 计算左子树的节点数
	left := 0
	if root.Left != nil {
		left = root.Left.nodeNums
	}
	m := left + 1

	// 判断当前节点是否为第 k 小元素
	if m == k {
		res = root.Val
	} else if m > k {
		findK(root.Left, k)
	} else if m < k {
		findK(root.Right, k-m)
	}
}
