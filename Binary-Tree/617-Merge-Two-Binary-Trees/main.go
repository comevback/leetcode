package main

import "fmt"

func main() {
	// 构建第一个二叉树
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 5}

	// 构建第二个二叉树
	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 7}

	// 合并两个二叉树
	res := mergeTrees(root1, root2)
	fmt.Println(res)
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mergeTrees 合并两个二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil // 如果两个树都为空，返回 nil
	} else if root1 == nil {
		root1 = &TreeNode{Val: 0} // 如果 root1 为空，创建一个值为 0 的节点
	} else if root2 == nil {
		root2 = &TreeNode{Val: 0} // 如果 root2 为空，创建一个值为 0 的节点
	}
	build(root1, root2) // 使用 build 函数合并树
	return root1
	// return get(root1, root2) // 或者使用 get 函数合并树
}

// build 函数递归合并两个树（遍历方法）
func build(root1 *TreeNode, root2 *TreeNode) {
	if root1 == nil {
		return // 如果 root1 为空，直接返回
	}
	root1.Val += root2.Val // 将 root2 的值加到 root1 上

	// 处理左子节点
	if root1.Left == nil && root2.Left != nil {
		root1.Left = &TreeNode{Val: 0} // 如果 root1 左子节点为空，root2 左子节点不为空，创建一个值为 0 的节点
	} else if root2.Left == nil && root1.Left != nil {
		root2.Left = &TreeNode{Val: 0} // 如果 root2 左子节点为空，root1 左子节点不为空，创建一个值为 0 的节点
	}

	// 处理右子节点
	if root1.Right == nil && root2.Right != nil {
		root1.Right = &TreeNode{Val: 0} // 如果 root1 右子节点为空，root2 右子节点不为空，创建一个值为 0 的节点
	} else if root2.Right == nil && root1.Right != nil {
		root2.Right = &TreeNode{Val: 0} // 如果 root2 右子节点为空，root1 右子节点不为空，创建一个值为 0 的节点
	}

	// 递归处理左子树和右子树
	build(root1.Left, root2.Left)
	build(root1.Right, root2.Right)
}

// get 函数递归合并两个树（分解方法）
func get(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2 // 如果 root1 为空，返回 root2
	}
	if root2 == nil {
		return root1 // 如果 root2 为空，返回 root1
	}

	root1.Val += root2.Val // 将 root2 的值加到 root1 上

	// 递归处理左子树和右子树
	root1.Left = get(root1.Left, root2.Left)
	root1.Right = get(root1.Right, root2.Right)

	return root1 // 返回合并后的树
}
