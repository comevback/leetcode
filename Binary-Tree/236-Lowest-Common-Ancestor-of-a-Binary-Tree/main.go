package main

import "fmt"

func main() {
	// 构造二叉树: [3,5,1,6,2,0,8,null,null,7,4]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	// 设置节点 p 和 q
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 4}

	// 查找最低公共祖先
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val) // 输出公共祖先的值
}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 变量 target 用于存储最低公共祖先节点
var target *TreeNode

// 查找最低公共祖先的主函数
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	target = &TreeNode{} // 初始化 target
	find(root, p, q)     // 调用辅助函数 find 查找公共祖先
	return target        // 返回找到的公共祖先
}

// 辅助函数 find 在二叉树中查找节点 p 和 q，并更新 target 为最低公共祖先
func find(root, p, q *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	var res int
	// 递归查找左子树和右子树
	res += find(root.Left, p, q) + find(root.Right, p, q)

	// 如果当前节点是 p 或 q，增加 res 的计数
	if root.Val == p.Val || root.Val == q.Val {
		res += 1
	}

	// 如果在左子树和右子树中找到了 p 和 q，则当前节点为公共祖先
	if res == 2 {
		target = root // 更新 target 为当前节点
		res = 0       // 重置 res，防止重复更新
	}
	return res // 返回当前节点的 res 计数
}
