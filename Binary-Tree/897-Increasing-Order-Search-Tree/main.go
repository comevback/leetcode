package main

import (
	"fmt"
)

func main() {
	// 构建示例树： [5,3,6,2,4,null,8,1,null,null,null,7,9]
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 9}

	// 调用 increasingBST 函数
	res := increasingBST(root)

	// 打印重新排列后的树
	for res != nil {
		fmt.Println(res.Val)
		res = res.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cur *TreeNode

// increasingBST 函数接收一个树的根节点，返回一个有序链表的根节点
func increasingBST(root *TreeNode) *TreeNode {
	// 创建一个哑结点
	res := &TreeNode{}
	// cur 指向哑结点
	cur = res
	// 进行中序遍历
	inOrder(root)
	// 返回新树的实际根节点（哑结点的右子节点）
	return res.Right
}

// inOrder 函数进行中序遍历，并重新构建树
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	// 递归遍历左子树
	inOrder(root.Left)

	// 处理当前节点
	// 将当前节点的左子节点设为 nil
	root.Left = nil
	// 将当前节点链接到新树的末尾
	cur.Right = root
	// 更新 cur 为当前节点
	cur = root

	// 递归遍历右子树
	inOrder(root.Right)
}

// getNode 函数接收一个树的根节点，返回一个有序链表的根节点
func getNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 递归处理左子树
	left := getNode(root.Left)
	// 将当前节点的左子节点设为 nil
	root.Left = nil
	// 递归处理右子树
	right := getNode(root.Right)
	// 将右子树链接到当前节点
	root.Right = right

	// 如果左子树为空，直接返回当前节点
	if left == nil {
		return root
	}

	// 如果左子树非空，需要将当前节点和右子树接到左子树的末尾
	p := left
	// 遍历左子树，找到最右端的节点
	for p != nil && p.Right != nil {
		p = p.Right
	}
	// 将当前节点链接到左子树的末尾
	p.Right = root

	// 返回新树的根节点
	return left
}
