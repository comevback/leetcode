# 3249. Count the Number of Good Nodes

Solved
Medium
Companies
Hint
There is an undirected tree with n nodes labeled from 0 to n - 1, and rooted at node 0. You are given a 2D integer array edges of length n - 1, where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

A node is good if all the
subtrees
rooted at its children have the same size.

Return the number of good nodes in the given tree.

A subtree of treeName is a tree consisting of a node in treeName and all of its descendants.

Example 1:

Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]

Output: 7

Explanation:

All of the nodes of the given tree are good.

Example 2:

Input: edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]

Output: 6

Explanation:

There are 6 good nodes in the given tree. They are colored in the image above.

Example 3:

Input: edges = [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]

Output: 12

Explanation:

All nodes except node 9 are good.

Constraints:

2 <= n <= 105
edges.length == n - 1
edges[i].length == 2
0 <= ai, bi < n
The input is generated such that edges represents a valid tree.

---

# Code

```go
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
```

详细思路解析：

1. 邻接表构建：
   使用 adjs 数组表示树的邻接表，adjs[i] 存储节点 i 的所有相邻节点。我们根据输入的 edges 数组填充这个邻接表。

2. DFS（深度优先搜索）遍历：
   通过递归的 dfs 函数来遍历树，并计算每个节点的子树大小。subTreeSize 数组用于缓存每个节点的子树大小，避免重复计算。

3. 判断“好节点”：
   对于每个节点，如果所有子节点的子树大小相同，那么它就是一个“好节点”。
   我们通过 isGood 变量来跟踪当前节点是否为“好节点”，并在遍历其子节点时进行判断。

4. 计数“好节点”：
   如果某个节点满足“好节点”的条件，我们将 goodCount 计数器增加一。

5. 返回结果：
   最终，通过 dfs(0, -1) 从根节点开始遍历树，并返回计算出的“好节点”总数。这段代码可以有效地计算给定树中的“好节点”数量，并使用详细的注释解释了每一步的实现和思路。
