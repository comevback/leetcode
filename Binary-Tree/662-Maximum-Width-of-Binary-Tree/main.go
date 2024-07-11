package main

import "fmt"

func main() {
	// 构建二叉树
	// tree = [1,3,2,5,null,null,9,6,null,7]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Left.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}
	root.Right.Right.Left = &TreeNode{Val: 7}

	// 计算二叉树的最大宽度
	res := widthOfBinaryTree(root)
	fmt.Println(res) // 输出结果
}

// TreeNode 结构体表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量 max 用于记录最大宽度
var max int

// widthOfBinaryTree 函数计算二叉树的最大宽度
func widthOfBinaryTree(root *TreeNode) int {
	max = 0
	travel(root)
	return max
}

// Node 结构体用于在队列中存储树节点及其对应的编号
type Node struct {
	treenode *TreeNode
	id       int
}

// travel 函数进行层序遍历，计算每层的宽度并更新最大宽度
func travel(root *TreeNode) {
	// 定义两个队列，一个用于当前层，一个用于下一层
	queue := []*Node{}
	newQueue := []*Node{}

	// 将根节点及其编号（1）加入队列
	queue = append(queue, &Node{treenode: root, id: 1})

	for {
		// 如果两个队列都为空，则退出循环
		if len(queue) == 0 && len(newQueue) == 0 {
			break
		}

		var left, right int // 记录当前层的最左和最右编号
		for i := 0; i < len(queue); i++ {
			// 记录当前层的最左节点编号
			if i == 0 {
				left = queue[i].id
			}

			// 记录当前层的最右节点编号
			if i == len(queue)-1 {
				right = queue[i].id
			}

			// 将左子节点加入下一层队列，并设置其编号
			if queue[i].treenode.Left != nil {
				newQueue = append(newQueue, &Node{treenode: queue[i].treenode.Left, id: 2 * queue[i].id})
			}

			// 将右子节点加入下一层队列，并设置其编号
			if queue[i].treenode.Right != nil {
				newQueue = append(newQueue, &Node{treenode: queue[i].treenode.Right, id: 2*queue[i].id + 1})
			}
		}

		// 计算当前层的宽度
		width := right - left + 1
		// 更新最大宽度
		if width > max {
			max = width
		}

		// 将下一层队列赋值给当前层队列，并清空下一层队列
		queue = newQueue
		newQueue = []*Node{}
	}
}
