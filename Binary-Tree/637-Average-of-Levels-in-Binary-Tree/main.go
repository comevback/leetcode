package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// averageOfLevels 函数返回二叉树每一层的平均值
func averageOfLevels(root *TreeNode) []float64 {
	// 用于存储每层平均值的结果集
	res := []float64{}

	// 初始化队列，将根节点加入队列
	queue := []*TreeNode{}
	queue = append(queue, root)

	// 层序遍历
	for len(queue) != 0 {
		length := len(queue) // 当前层的节点数
		sum := 0             // 当前层节点值的总和

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
			// 累加当前节点的值
			sum += queue[i].Val
		}

		// 移除当前层的节点
		queue = queue[length:]

		// 计算当前层的平均值
		avr := float64(sum) / float64(length)
		// 将平均值加入结果集
		res = append(res, avr)
	}

	return res
}
