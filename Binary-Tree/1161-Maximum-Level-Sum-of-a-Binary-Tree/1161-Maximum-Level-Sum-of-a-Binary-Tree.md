# 1161. Maximum Level Sum of a Binary Tree

Solved
Medium
Topics
Companies
Hint
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

Example 1:

Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.
Example 2:

Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2

Constraints:

The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105

---

# Code

```go
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
```
