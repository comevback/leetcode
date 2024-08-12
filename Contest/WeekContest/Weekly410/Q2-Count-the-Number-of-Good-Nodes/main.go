package main

import (
	"fmt"
)

func main() {
	// 输入的边数组，表示一棵树 [6,0],[1,0],[5,1],[2,5],[3,1],[4,3]
	edges := [][]int{{6, 0}, {1, 0}, {5, 1}, {2, 5}, {3, 1}, {4, 3}}
	res := countGoodNodes(edges) // 计算树中的“好节点”数量
	fmt.Println(res)             // 输出结果
}

// countGoodNodes 函数计算树中“好节点”的数量
func countGoodNodes(edges [][]int) int {
	// 创建邻接表，表示每个节点的相邻节点
	// adjs 数组的长度为节点数量（等于边的数量 + 1）
	adjs := make([][]int, len(edges)+1)
	for _, edge := range edges {
		adjs[edge[0]] = append(adjs[edge[0]], edge[1]) // 添加边的两端节点到各自的邻接列表中
		adjs[edge[1]] = append(adjs[edge[1]], edge[0])
	}

	var goodCount int                        // 用于统计“好节点”的数量
	subTreeSize := make([]int, len(edges)+1) // 用于存储每个节点的子树大小

	// 定义 DFS 函数，递归计算每个节点的子树大小，并判断是否为“好节点”
	var dfs func(int, int) int
	dfs = func(node, parent int) int {
		// 如果子树大小已经计算过，则直接返回，避免重复计算
		if subTreeSize[node] != 0 {
			return subTreeSize[node]
		}

		size := 1            // 初始化当前节点的子树大小为 1（包括自身）
		isGood := true       // 用于判断当前节点是否为“好节点”
		var childSizes []int // 用于存储所有子节点的子树大小

		// 遍历当前节点的所有子节点
		for _, child := range adjs[node] {
			if child == parent { // 如果子节点是父节点，则跳过（避免回到父节点导致死循环）
				continue
			}
			childSize := dfs(child, node)              // 递归计算子节点的子树大小
			childSizes = append(childSizes, childSize) // 将子节点的子树大小添加到数组中
			size += childSize                          // 更新当前节点的子树大小
		}

		// 检查所有子节点的子树大小是否相同
		if len(childSizes) > 0 {
			firstChildSize := childSizes[0] // 第一个子节点的子树大小
			for _, v := range childSizes {
				if v != firstChildSize { // 如果有不同的子树大小，则当前节点不是“好节点”
					isGood = false
					break
				}
			}
		}

		// 如果当前节点是“好节点”，则计数加 1
		if isGood {
			goodCount++
		}

		subTreeSize[node] = size // 保存当前节点的子树大小，以便后续使用
		return size              // 返回当前节点的子树大小
	}

	dfs(0, -1)       // 从根节点（假设为 0）开始 DFS 遍历，父节点设为 -1 表示没有父节点
	return goodCount // 返回“好节点”的总数量
}
