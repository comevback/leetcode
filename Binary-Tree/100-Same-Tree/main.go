package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree 判断两棵二叉树是否相同
// 如果两棵二叉树结构相同且节点值相同，则返回 true，否则返回 false
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 如果两棵树都为空，返回 true
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil { // 如果其中一棵树为空，另一棵树不为空，返回 false
		return false
	}

	// 如果当前节点值不相同，返回 false
	if p.Val != q.Val {
		return false
	}

	// 递归比较左子树和右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
