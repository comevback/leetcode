# 987. Vertical Order Traversal of a Binary Tree
Solved
Hard
Topics
Companies
Given the root of a binary tree, calculate the vertical order traversal of the binary tree.

For each node at position (row, col), its left and right children will be at positions (row + 1, col - 1) and (row + 1, col + 1) respectively. The root of the tree is at (0, 0).

The vertical order traversal of a binary tree is a list of top-to-bottom orderings for each column index starting from the leftmost column and ending on the rightmost column. There may be multiple nodes in the same row and same column. In such a case, sort these nodes by their values.

Return the vertical order traversal of the binary tree.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[9],[3,15],[20],[7]]
Explanation:
Column -1: Only node 9 is in this column.
Column 0: Nodes 3 and 15 are in this column in that order from top to bottom.
Column 1: Only node 20 is in this column.
Column 2: Only node 7 is in this column.
Example 2:

Input: root = [1,2,3,4,5,6,7]
Output: [[4],[2],[1,5,6],[3],[7]]
Explanation:
Column -2: Only node 4 is in this column.
Column -1: Only node 2 is in this column.
Column 0: Nodes 1, 5, and 6 are in this column.
          1 is at the top, so it comes first.
          5 and 6 are at the same position (2, 0), so we order them by their value, 5 before 6.
Column 1: Only node 3 is in this column.
Column 2: Only node 7 is in this column.
Example 3:

Input: root = [1,2,3,4,6,5,7]
Output: [[4],[2],[1,5,6],[3],[7]]
Explanation:
This case is the exact same as example 2, but with nodes 5 and 6 swapped.
Note that the solution remains the same since 5 and 6 are in the same location and should be ordered by their values.

Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 1000

---

# Code
```go
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
```