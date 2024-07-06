package main

import "fmt"

func main() {
	// 构建初始的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}

	// 要删除的节点值
	to_delete := []int{2, 1}
	res := delNodes(root, to_delete)
	// 打印结果森林的根节点值
	for _, v := range res {
		fmt.Println(v.Val)
	}
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nodeMap map[int]int // 用于存储要删除的节点值
var res []*TreeNode     // 存储结果森林的根节点

// delNodes 删除指定节点值的节点，并返回森林的根节点
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res = []*TreeNode{}           // 初始化结果列表
	nodeMap = make(map[int]int)   // 初始化节点值映射
	for _, v := range to_delete { // 将要删除的节点值存储到映射中
		nodeMap[v] += 1
	}
	del(root) // 调用删除函数
	// 如果根节点不在删除列表中，添加到结果列表
	if nodeMap[root.Val] == 0 {
		res = append(res, root)
	}
	return res // 返回结果列表
}

// del 递归删除指定节点值的节点
func del(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归删除左子树和右子树的指定节点
	del(root.Left)
	del(root.Right)

	// 如果左子树节点在删除列表中，删除左子树节点
	if root.Left != nil && nodeMap[root.Left.Val] != 0 {
		root.Left = nil
	}
	// 如果右子树节点在删除列表中，删除右子树节点
	if root.Right != nil && nodeMap[root.Right.Val] != 0 {
		root.Right = nil
	}

	// 如果当前节点在删除列表中，将其子节点作为新的根节点添加到结果列表中
	if nodeMap[root.Val] != 0 {
		if root.Left != nil {
			res = append(res, root.Left)
		}
		if root.Right != nil {
			res = append(res, root.Right)
		}
		// 将当前节点从父节点中删除
		root.Left = nil
		root.Right = nil
	}
}
