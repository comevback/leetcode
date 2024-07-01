package main

import (
	"fmt"
)

func main() {
	// 创建示例二叉树: [4,2,6,3,1,5]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	// 在深度为 2 的位置添加值为 1 的新行
	res := addOneRow(root, 1, 2)
	fmt.Println(res) // 打印修改后的二叉树根节点
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// addOneRow 在给定深度处添加一行值为 val 的新节点
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	// 如果深度为 1，则创建一个新节点作为根节点，并将原根节点作为其左子节点
	if depth == 1 {
		newNode := &TreeNode{
			Val: val,
		}
		newNode.Left = root
		root = newNode
	}

	// 使用广度优先搜索在指定深度处添加新行
	bfs(root, depth, val)
	return root
}

// bfs 使用广度优先搜索在指定深度处添加新行
func bfs(root *TreeNode, depth int, val int) {
	currentDepth := 1
	queue := []*TreeNode{}    // 当前层节点队列
	newQueue := []*TreeNode{} // 下一层节点队列
	queue = append(queue, root)

	// 遍历到指定深度的前一层
	for currentDepth < depth-1 {
		for len(queue) != 0 {
			if queue[0].Left != nil {
				newQueue = append(newQueue, queue[0].Left)
			}
			if queue[0].Right != nil {
				newQueue = append(newQueue, queue[0].Right)
			}
			queue = queue[1:]
		}
		// 更新队列为下一层节点
		queue = newQueue
		newQueue = []*TreeNode{}
		currentDepth += 1
	}

	// 在指定深度处添加新行
	for len(queue) != 0 {
		// 添加左子节点
		if queue[0].Left != nil {
			newNode := &TreeNode{
				Val: val,
			}
			newNode.Left = queue[0].Left
			queue[0].Left = newNode
		} else {
			queue[0].Left = &TreeNode{
				Val: val,
			}
		}

		// 添加右子节点
		if queue[0].Right != nil {
			newNode := &TreeNode{
				Val: val,
			}
			newNode.Right = queue[0].Right
			queue[0].Right = newNode
		} else {
			queue[0].Right = &TreeNode{
				Val: val,
			}
		}

		queue = queue[1:]
	}
}
