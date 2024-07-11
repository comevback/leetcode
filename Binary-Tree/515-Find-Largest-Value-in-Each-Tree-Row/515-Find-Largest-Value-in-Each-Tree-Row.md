# 515. Find Largest Value in Each Tree Row

Solved
Medium
Topics
Companies
Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).

Example 1:

Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]
Example 2:

Input: root = [1,2,3]
Output: [1,3]

Constraints:

The number of nodes in the tree will be in the range [0, 104].
-231 <= Node.val <= 231 - 1

---

# Code

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 构建二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 9}

	// 获取每一层的最大值
	res := largestValues(root)
	fmt.Println(res) // 输出结果
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量 res 用于存储每层的最大值
var res []int

// largestValues 函数返回二叉树每一层的最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// 初始化 res 为空切片
	res = []int{}

	// 初始化队列，将根节点加入队列
	queue := []*TreeNode{}
	queue = append(queue, root)

	// 层序遍历
	for len(queue) != 0 {
		tempMax := math.MinInt // 初始化当前层的最大值为最小整数
		length := len(queue)   // 当前层的节点数

		for i := 0; i < length; i++ {
			// 将当前节点的左子节点加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 将当前节点的右子节点加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 更新当前层的最大值
			if queue[i].Val > tempMax {
				tempMax = queue[i].Val
			}
		}
		// 移除当前层的节点
		queue = queue[length:]
		// 将当前层的最大值加入结果集
		res = append(res, tempMax)
	}

	return res
}
```
