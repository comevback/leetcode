package main

import (
	"fmt"
	"math"
)

func main() {
	// 创建示例二叉搜索树: [4,2,5,1,3]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	// 目标值
	target := 3.714286
	// 查找最接近目标值的节点值
	res := closestValue(root, target)
	fmt.Println(res) // 输出最接近目标值的节点值
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var closest int
var minDiff float64

// closestValue 返回二叉搜索树中最接近目标值的节点值
func closestValue(root *TreeNode, target float64) int {
	closest = 0
	minDiff = math.MaxInt // 初始化最小差值为一个很大的值
	travel(root, target)  // 遍历二叉树以找到最接近的值
	return closest
}

// travel 递归遍历二叉树以更新最接近目标值的节点值
func travel(root *TreeNode, target float64) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	diff := abs(float64(root.Val) - target) // 计算当前节点值与目标值的差异
	if diff < minDiff {
		// 如果当前差异小于最小差值，则更新最小差值和最接近的节点值
		minDiff = diff
		closest = root.Val
	} else if diff == minDiff && root.Val < closest {
		// 如果差异相等但当前节点值更小，则更新最接近的节点值
		closest = root.Val
	}

	// 根据目标值与当前节点值的比较，决定递归遍历左子树还是右子树
	if target < float64(root.Val) {
		travel(root.Left, target)
	} else {
		travel(root.Right, target)
	}
}

// abs 返回一个浮点数的绝对值
func abs(num float64) float64 {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
