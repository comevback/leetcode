# 1302. Deepest Leaves Sum

Solved
Medium
Topics
Companies
Hint
Given the root of a binary tree, return the sum of values of its deepest leaves.

Example 1:

Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15
Example 2:

Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 19

Constraints:

The number of nodes in the tree is in the range [1, 104].
1 <= Node.val <= 100

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

func deepestLeavesSum(root *TreeNode) int {
	// 如果根节点为空，返回 0
	if root == nil {
		return 0
	}
	// 使用切片作为队列进行层序遍历
	queue := []*TreeNode{}
	// 将根节点加入队列
	queue = append(queue, root)

	// 开始层序遍历
	for len(queue) != 0 {
		// 当前层节点值的和
		tempSum := 0
		// 当前层节点数
		length := len(queue)
		// 遍历当前层的每个节点
		for i := 0; i < length; i++ {
			// 如果左子节点不为空，将其加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 如果右子节点不为空，将其加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 累加当前节点的值到 tempSum
			tempSum += queue[i].Val
		}
		// 移除当前层的节点
		queue = queue[length:]
		// 如果队列为空，表示当前层是最深层，返回当前层的节点和
		if len(queue) == 0 {
			return tempSum
		}
	}
	// 如果树为空，返回 0（理论上不会到达这里，因为上面已经处理了 root == nil 的情况）
	return 0
}
```
