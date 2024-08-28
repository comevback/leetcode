package main

func main() {

}

// TreeNode 定义了二叉树节点的结构，用于逻辑运算树
type TreeNode struct {
	Val   int       // 节点的值，可以是 0（假），1（真），2（逻辑或），3（逻辑与）
	Left  *TreeNode // 指向左子节点
	Right *TreeNode // 指向右子节点
}

// evaluateTree 对给定的逻辑运算树进行求值。
// 该函数递归地对树进行遍历，根据节点的类型（操作数或运算符）计算最终的布尔值。
func evaluateTree(root *TreeNode) bool {
	// 如果节点是叶子节点，则根据其值直接返回 true 或 false
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false // 节点值为0表示假
		} else if root.Val == 1 {
			return true // 节点值为1表示真
		}
	}

	// 递归地计算左子树和右子树的值
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	// 根据当前节点的运算符，对子树的值进行逻辑运算
	if root.Val == 2 {
		return left || right // 节点值为2表示逻辑或
	} else {
		return left && right // 节点值为3表示逻辑与
	}
}
