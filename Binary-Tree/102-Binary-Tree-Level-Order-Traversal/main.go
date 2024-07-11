package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		// 如果树为空，返回空的二维数组
		return [][]int{}
	}

	// 初始化结果集
	res := [][]int{}
	// 初始化队列，将根节点加入队列
	queue := []*TreeNode{root}

	// 层序遍历
	for len(queue) != 0 {
		// 用于存储当前层的节点值
		temp := []int{}
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
			// 将当前节点的值加入当前层的结果集
			temp = append(temp, queue[i].Val)
		}

		// 移除当前层的节点
		queue = queue[length:]
		// 将当前层的结果加入总结果集
		res = append(res, temp)
	}

	return res
}
