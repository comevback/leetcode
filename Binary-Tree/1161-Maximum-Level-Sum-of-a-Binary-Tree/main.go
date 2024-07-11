package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	// 初始化最大和为最小整数，最大层号为 1
	max := math.MinInt
	maxLevel := 1
	// 当前层号为 1
	currentLevel := 1
	// 使用切片作为队列进行层序遍历
	queue := []*TreeNode{}
	// 将根节点加入队列
	queue = append(queue, root)

	// 层序遍历
	for len(queue) != 0 {
		// 当前层的节点和
		tempSum := 0
		// 当前层的节点数
		length := len(queue)
		// 遍历当前层的每个节点
		for i := 0; i < length; i++ {
			// 将左子节点加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 将右子节点加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			// 累加当前层的节点值
			tempSum += queue[i].Val
		}

		// 移除当前层的节点
		queue = queue[length:]
		// 如果当前层的节点和大于最大和，更新最大和和最大层号
		if tempSum > max {
			max = tempSum
			maxLevel = currentLevel
		}

		// 进入下一层
		currentLevel += 1
	}

	// 返回具有最大层和的层号
	return maxLevel
}
