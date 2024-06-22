package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	res := bstToGst(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bstToGst 将二叉搜索树 (BST) 转换为较大和树 (GST)
func bstToGst(root *TreeNode) *TreeNode {
	var sum int = 0
	reverseOrder(root, &sum)
	// printRoot(root)
	return root
}

// reverseOrder 以逆中序遍历的方式遍历树
func reverseOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 先遍历右子树
	reverseOrder(root.Right, sum)
	// 临时保存当前节点的值
	temp := root.Val
	// 将当前节点的值加上累积和
	root.Val += *sum
	// 更新累积和
	*sum += temp
	// 再遍历左子树
	reverseOrder(root.Left, sum)
}

// func inOrder(root *TreeNode, vals *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	inOrder(root.Left, vals)
// 	*vals = append(*vals, root.Val)
// 	inOrder(root.Right, vals)
// }

// func printRoot(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	printRoot(root.Left)
// 	fmt.Println(root.Val)
// 	printRoot(root.Right)
// }
