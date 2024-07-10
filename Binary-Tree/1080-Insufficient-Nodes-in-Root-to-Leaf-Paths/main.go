package main

import "fmt"

func main() {
	// 构建一个测试二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: -5}
	root.Right = &TreeNode{Val: -3}
	root.Right.Left = &TreeNode{Val: 4}

	// 调用 sufficientSubset 函数，参数为根节点和限制值 -1
	res := sufficientSubset(root, -1)
	fmt.Println(res)
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sufficientSubset 函数返回修改后的二叉树
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	// 调用递归函数 travel，初始累加值为 0
	_, res := travel(root, 0, limit)
	return res
}

// travel 函数进行递归遍历，返回所需的最小值和修改后的节点
func travel(root *TreeNode, haved int, limit int) (int, *TreeNode) {
	if root == nil {
		return limit, nil
	}

	var needed int
	// 递归处理左子树
	leftNeeded, leftNode := travel(root.Left, haved+root.Val, limit)
	// 递归处理右子树
	rightNeeded, rightNode := travel(root.Right, haved+root.Val, limit)

	// 如果当前节点是叶子节点
	if root.Left == nil && root.Right == nil {
		needed = limit
	} else if root.Left == nil { // 如果只有右子树
		needed = rightNeeded
	} else if root.Right == nil { // 如果只有左子树
		needed = leftNeeded
	} else { // 如果左右子树都有
		needed = min(leftNeeded, rightNeeded)
	}

	// 更新当前节点的左右子树
	root.Left = leftNode
	root.Right = rightNode

	// 计算需要的值减去当前节点的值
	needed -= root.Val

	// 如果需要的值大于已经累加的值，返回 nil 表示删除该节点
	if needed > haved {
		return needed, nil
	} else {
		return needed, root
	}
}

// min 函数返回两个整数中的较小值
func min(num1 int, num2 int) int {
	if num1 <= num2 {
		return num1
	} else {
		return num2
	}
}
