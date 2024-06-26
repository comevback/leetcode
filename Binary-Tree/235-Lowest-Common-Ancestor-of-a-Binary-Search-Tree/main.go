package main

func main() {
	// 示例代码，可以在此处构建二叉树并调用 lowestCommonAncestor 函数进行测试
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

	// 如果当前节点值小于 p 和 q 的值，递归查找右子树
	if root.Val < q.Val && root.Val < p.Val {
		return find(root.Right, p, q)
	} else if root.Val > q.Val && root.Val > p.Val {
		// 如果当前节点值大于 p 和 q 的值，递归查找左子树
		return find(root.Left, p, q)
	}

	// 如果当前节点值在 p 和 q 之间或者等于 p 或 q，继续查找
	res := 0
	if root.Val == q.Val || root.Val == p.Val {
		res += 1 // 如果当前节点是 p 或 q，增加 res 的计数
	}
	// 递归查找左子树和右子树，将结果累加到 res 中
	res += find(root.Right, p, q) + find(root.Left, p, q)

	// 如果在左子树和右子树中找到了 p 和 q，更新 target 为当前节点
	if res == 2 {
		target = root // 更新 target 为当前节点
		res = 0       // 重置 res，防止重复更新
	}

	return res // 返回当前节点的 res 计数
}
