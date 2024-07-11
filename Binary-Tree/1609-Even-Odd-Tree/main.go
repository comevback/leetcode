package main

import "math"

func main() {
	// 示例用例
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 12}

	result := isEvenOddTree(root)
	println(result) // 输出 true 或 false
}

// TreeNode 结构体表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isEvenOddTree 判断二叉树是否为偶数奇数树
func isEvenOddTree(root *TreeNode) bool {
	// 如果根节点为空，返回 false
	if root == nil {
		return false
	}
	// 使用切片作为队列进行层序遍历
	queue := []*TreeNode{}
	// 将根节点加入队列
	queue = append(queue, root)
	// 标记当前层是否为偶数层
	even := true
	// 定义前一个节点值和当前节点值
	var pre, cur int

	// 开始层序遍历
	for len(queue) != 0 {
		// 根据当前层是否为偶数层初始化 pre 和 cur
		if even {
			pre, cur = math.MinInt, math.MinInt
		} else {
			pre, cur = math.MaxInt, math.MaxInt
		}
		// 获取当前层的节点数
		length := len(queue)
		// 遍历当前层的每个节点
		for i := 0; i < length; i++ {
			// 将当前节点的左子节点加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 将当前节点的右子节点加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 更新前一个节点值
			pre = cur
			// 更新当前节点值
			cur = queue[i].Val
			// 判断当前层的节点值是否满足条件
			if even {
				// 偶数层：节点值必须是奇数且严格递增
				if queue[i].Val%2 == 0 || pre >= cur {
					return false
				}
			} else {
				// 奇数层：节点值必须是偶数且严格递减
				if queue[i].Val%2 != 0 || pre <= cur {
					return false
				}
			}
		}
		// 切换到下一层
		even = !even
		// 移除当前层的节点
		queue = queue[length:]
	}
	// 如果所有层的节点值都满足条件，返回 true
	return true
}
