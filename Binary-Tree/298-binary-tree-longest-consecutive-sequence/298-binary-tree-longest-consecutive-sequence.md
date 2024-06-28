# 298. 二叉树最长连续序列
给你一棵指定的二叉树的根节点 root ，请你计算其中 最长连续序列路径 的长度。

最长连续序列路径 是依次递增 1 的路径。该路径，可以是从某个初始节点到树中任意节点，通过「父 - 子」关系连接而产生的任意路径。且必须从父节点到子节点，反过来是不可以的。
 
示例 1：

输入：root = [1,null,3,2,4,null,null,null,5]
输出：3
解释：当中，最长连续序列是 3-4-5 ，所以返回结果为 3 。
解释
示例 2：

输入：root = [2,null,3,2,null,1]
输出：2
解释：当中，最长连续序列是 2-3 。注意，不是 3-2-1，所以返回 2 。
解释
提示：

树中节点的数目在范围 [1, 3 * 104] 内
-3 * 104 <= Node.val <= 3 * 104

---

# Code
```go
package main

func main() {
	// 测试用例
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	// 输出最长连续序列长度
	println(longestConsecutive(root)) // 输出应为 3
}

// TreeNode 定义了二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个全局变量 maxLength 用于存储最长连续序列的长度
var maxLength int

// longestConsecutive 函数接受一个二叉树的根节点并返回该二叉树的最长连续序列的长度
func longestConsecutive(root *TreeNode) int {
	// 初始化 maxLength 为 0
	maxLength = 0
	// 从根节点开始进行前序遍历，初始深度为 1
	preOrder(root, 1)
	// 返回最长连续序列的长度
	return maxLength
}

// preOrder 函数进行前序遍历以查找最长连续序列
// root: 当前节点
// depth: 当前连续序列的长度
func preOrder(root *TreeNode, depth int) {
	// 如果当前节点为空，直接返回
	if root == nil {
		return
	}

	// 如果当前深度大于 maxLength，更新 maxLength
	if depth > maxLength {
		maxLength = depth
	}

	// 递归遍历左子树
	if root.Left != nil {
		// 如果左子节点值与当前节点值连续，则增加深度
		if root.Left.Val == root.Val+1 {
			preOrder(root.Left, depth+1)
			// 否则重置深度为 1
		} else {
			preOrder(root.Left, 1)
		}
	}

	// 递归遍历右子树
	if root.Right != nil {
		// 如果右子节点值与当前节点值连续，则增加深度
		if root.Right.Val == root.Val+1 {
			preOrder(root.Right, depth+1)
			// 否则重置深度为 1
		} else {
			preOrder(root.Right, 1)
		}
	}
}
```