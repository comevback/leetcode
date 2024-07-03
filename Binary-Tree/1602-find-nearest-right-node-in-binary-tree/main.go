package main

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findNearestRightNode 查找目标节点 u 的最近右侧邻居节点
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	return BFS(root, u)
}

// BFS 使用广度优先搜索（层次遍历）来查找目标节点的最近右侧邻居节点
func BFS(root *TreeNode, target *TreeNode) *TreeNode {
	queue := []*TreeNode{}      // 当前层节点队列
	newQueue := []*TreeNode{}   // 下一层节点队列
	queue = append(queue, root) // 将根节点添加到当前层队列

	for {
		if len(queue) == 0 && len(newQueue) == 0 {
			break // 如果当前层和下一层队列都为空，结束搜索
		}

		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left) // 将左子节点添加到下一层队列
			}

			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right) // 将右子节点添加到下一层队列
			}

			if queue[i] == target {
				if i+1 < len(queue) {
					return queue[i+1] // 返回目标节点的右侧邻居节点
				}
			}
		}

		queue = newQueue         // 更新当前层队列为下一层队列
		newQueue = []*TreeNode{} // 重置下一层队列
	}

	return nil // 如果没有找到右侧邻居节点，返回 nil
}
