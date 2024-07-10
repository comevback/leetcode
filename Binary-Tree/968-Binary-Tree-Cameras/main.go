package main

import "fmt"

func main() {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Left.Left = &TreeNode{Val: 0}

	res := minCameraCover(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minCameraCover 函数计算覆盖所有节点所需的最小摄像头数量
func minCameraCover(root *TreeNode) int {
	// 如果根节点是唯一的节点，返回1个摄像头
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 调用递归函数 travel
	totalNums, _, headCover := travel(root)
	// 如果根节点没有被覆盖，增加一个摄像头
	if !headCover {
		totalNums += 1
	}
	return totalNums
}

// travel 函数返回值为 (摄像头数量, 当前节点是否安装摄像头, 当前节点是否被覆盖)
func travel(root *TreeNode) (int, bool, bool) {
	// 基本情况，如果当前节点为空，返回0个摄像头，且节点未被覆盖
	if root == nil {
		return 0, false, false
	}

	// 如果当前节点是叶子节点，返回0个摄像头，且节点未被覆盖
	if root.Left == nil && root.Right == nil {
		return 0, false, false
	}

	// 初始化变量
	var camera bool = false                                  // 当前节点是否安装摄像头
	var cameraNum int = 0                                    // 需要的摄像头数量
	var cover bool = false                                   // 当前节点是否被覆盖
	leftNums, leftCamera, leftCover := travel(root.Left)     // 递归左子树
	rightNums, rightCamera, rightCover := travel(root.Right) // 递归右子树

	// 如果左子树或右子树安装了摄像头，当前节点被覆盖
	if leftCamera || rightCamera {
		cover = true
	}

	// 如果左子树或右子树的子节点没有被覆盖，当前节点需要安装摄像头
	if (root.Left != nil && !leftCover) || (root.Right != nil && !rightCover) {
		camera = true
		cover = true
		cameraNum = 1
	}

	return leftNums + rightNums + cameraNum, camera, cover
}
