package main

import "fmt"

func main() {
	// 构建示例二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	// 检查值为 1 和 3 的节点是否为堂兄弟节点
	res := isCousins(root, 1, 3)
	fmt.Println(res) // 输出结果
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储目标节点信息
var a, b int
var tarDepth int
var tarParent *TreeNode
var res bool

// isCousins 检查两个节点是否为堂兄弟节点
func isCousins(root *TreeNode, x int, y int) bool {
	a, b = x, y
	tarDepth = -1
	tarParent = &TreeNode{}

	// 遍历树以寻找目标节点
	travel(root, nil, 0)
	return res
}

// travel 遍历二叉树以检查节点是否为堂兄弟节点
func travel(root *TreeNode, parent *TreeNode, depth int) {
	if root == nil {
		return
	}

	// 如果找到目标节点之一
	if root.Val == a || root.Val == b {
		if tarDepth == -1 {
			// 记录第一个找到的目标节点的深度和父节点
			tarDepth = depth
			tarParent = parent
		} else {
			// 检查第二个找到的目标节点是否满足堂兄弟条件
			if tarDepth == depth && parent != tarParent {
				res = true
			} else {
				res = false
			}
		}
	}

	// 递归遍历左子树
	if root.Left != nil {
		travel(root.Left, root, depth+1)
	}

	// 递归遍历右子树
	if root.Right != nil {
		travel(root.Right, root, depth+1)
	}
}
