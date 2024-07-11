package main

import "fmt"

func main() {
	// 构建二叉树并进行完全性检查
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}

	res := isCompleteTree(root)
	fmt.Println(res) // 输出二叉树是否为完全二叉树
}

// TreeNode 结构体表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node 结构体表示节点及其编号
type Node struct {
	node *TreeNode
	id   int
}

// isCompleteTree 函数检查二叉树是否为完全二叉树
func isCompleteTree(root *TreeNode) bool {
	// 初始化队列，将根节点加入队列
	queue := []*Node{}
	queue = append(queue, &Node{node: root, id: 1})
	// pre 和 cur 用于记录前一个节点和当前节点的编号
	var pre, cur int

	// 层序遍历
	for len(queue) != 0 {
		// 获取当前层的节点数
		length := len(queue)

		// 遍历当前层的每个节点
		for i := 0; i < length; i++ {
			// 将当前节点的左子节点加入队列，并编号为 2 * id
			if queue[i].node.Left != nil {
				queue = append(queue, &Node{node: queue[i].node.Left, id: 2 * queue[i].id})
			}
			// 将当前节点的右子节点加入队列，并编号为 2 * id + 1
			if queue[i].node.Right != nil {
				queue = append(queue, &Node{node: queue[i].node.Right, id: 2*queue[i].id + 1})
			}
			// 更新前一个节点和当前节点的编号
			pre = cur
			cur = queue[i].id
			// 如果前一个节点的编号不等于当前节点编号减一，返回 false
			if pre != cur-1 {
				return false
			}
		}
		// 移除当前层的节点
		queue = queue[length:]
	}
	// 如果遍历完所有节点没有发现不符合完全二叉树性质的情况，返回 true
	return true
}
