package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue1(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}
	var smallest, second int = math.MaxInt, math.MaxInt
	travel(root, &smallest, &second)
	if second != math.MaxInt {
		return second
	} else {
		return -1
	}
}

// 遍历解法
func travel(root *TreeNode, small, second *int) {
	if root == nil {
		return
	}

	if root.Val <= *small {
		*small = root.Val
	} else if root.Val < *second {
		*second = root.Val
	}

	travel(root.Left, small, second)
	travel(root.Right, small, second)
}

// solution
func findSecondMinimumValue(root *TreeNode) int {
	// 如果左右子节点都为空，说明树中只有一个节点，没有第二小的值
	if root.Left == nil && root.Right == nil {
		return -1
	}

	// 初始化左右子节点的值
	left, right := root.Left.Val, root.Right.Val

	// 如果左子节点的值等于根节点的值，递归寻找左子树中的第二小值
	if root.Val == root.Left.Val {
		left = findSecondMinimumValue(root.Left)
	}

	// 如果右子节点的值等于根节点的值，递归寻找右子树中的第二小值
	if root.Val == root.Right.Val {
		right = findSecondMinimumValue(root.Right)
	}

	// 如果左子树没有找到第二小值，返回右子树的结果
	if left == -1 {
		return right
	}

	// 如果右子树没有找到第二小值，返回左子树的结果
	if right == -1 {
		return left
	}

	// 左右子树都找到第二小值，返回其中较小的一个
	return min(left, right)
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}
