package main

import "math"

func main() {

}

// TreeNode 结构体定义了二叉树节点
type TreeNode struct {
	Val   int       // 节点存储的整数值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

var res []int // 存储结果的切片，记录出现次数最多的值
var max int   // 记录最大的出现次数
var last int  // 上一个访问的节点值
var times int // 当前值出现的次数

// findMode 函数查找给定二叉树中的众数（出现次数最多的元素）
func findMode(root *TreeNode) []int {
	res = []int{}
	last = math.MaxInt // 将 last 初始化为最大整数，避免与树中的任何数相等
	times = 0
	max = 0
	travel(root) // 调用 travel 函数开始中序遍历
	return res   // 返回结果
}

// travel 函数进行中序遍历
func travel(root *TreeNode) {
	if root == nil {
		return // 如果节点为空，返回
	}

	travel(root.Left) // 递归访问左子树

	// 处理当前节点
	if root.Val == last {
		times += 1 // 如果当前节点值等于上一个节点值，增加计数
	} else {
		times = 1 // 如果不等，重置计数并更新 last
		last = root.Val
	}

	// 更新结果
	if times > max {
		max = times
		res = []int{root.Val} // 如果当前次数大于已知最大次数，更新结果为当前节点值
	} else if times == max {
		res = append(res, root.Val) // 如果当前次数等于最大次数，将当前值添加到结果中
	}

	travel(root.Right) // 递归访问右子树
}
