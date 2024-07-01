package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用数组来保存结果，每个索引对应一个 x 坐标
var resArr [2000][]int

// verticalTraversal 返回二叉树的垂直遍历结果
func verticalTraversal(root *TreeNode) [][]int {
	// 初始化结果数组
	resArr = [2000][]int{}
	height := 0
	x := 0

	// 进行广度优先搜索
	travelBFS(root, height, x)

	// 将结果数组中的非空部分转为切片返回
	res := [][]int{}
	for _, v := range resArr {
		if len(v) != 0 {
			res = append(res, v)
		}
	}

	return res
}

// Xnode 用于保存树节点及其对应的 x 坐标
type Xnode struct {
	node *TreeNode
	x    int
}

// travelBFS 进行二叉树的广度优先搜索，并记录节点值到结果数组
func travelBFS(root *TreeNode, height int, x int) {
	queue := []*Xnode{}    // 当前层节点队列
	newQueue := []*Xnode{} // 下一层节点队列
	currentXmap := make(map[int]int)

	// 将根节点加入队列
	queue = append(queue, &Xnode{
		root,
		x,
	})

	// 开始广度优先搜索
	for {
		if len(queue) == 0 && len(newQueue) == 0 {
			break // 如果当前层和下一层都为空，结束搜索
		}

		for i := 0; i < len(queue); i++ {
			if queue[i].node.Left != nil {
				// 将左子节点加入下一层队列，x 坐标减 1
				newXnode := &Xnode{
					queue[i].node.Left,
					queue[i].x - 1,
				}
				newQueue = append(newQueue, newXnode)
			}

			if queue[i].node.Right != nil {
				// 将右子节点加入下一层队列，x 坐标加 1
				newXnode := &Xnode{
					queue[i].node.Right,
					queue[i].x + 1,
				}
				newQueue = append(newQueue, newXnode)
			}

			// 将节点值加入结果数组相应位置
			resArr[queue[i].x+1000] = append(resArr[queue[i].x+1000], queue[i].node.Val)

			// 对于当前 x 坐标，按照节点值升序排列
			for j := len(resArr[queue[i].x+1000]) - 1; j > 0 && j > len(resArr[queue[i].x+1000])-1-currentXmap[queue[i].x]; j-- {
				if resArr[queue[i].x+1000][j] < resArr[queue[i].x+1000][j-1] {
					resArr[queue[i].x+1000][j], resArr[queue[i].x+1000][j-1] = resArr[queue[i].x+1000][j-1], resArr[queue[i].x+1000][j]
				}
			}
			currentXmap[queue[i].x] += 1
		}
		// 更新队列为下一层
		queue = newQueue
		newQueue = []*Xnode{}
		currentXmap = make(map[int]int)
		height += 1
	}
}
