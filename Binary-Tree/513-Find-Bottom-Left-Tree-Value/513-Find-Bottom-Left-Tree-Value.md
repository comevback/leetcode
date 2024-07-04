# 513. Find Bottom Left Tree Value
Solved
Medium
Topics
Companies
Given the root of a binary tree, return the leftmost value in the last row of the tree.

Example 1:

Input: root = [2,1,3]
Output: 1
Example 2:

Input: root = [1,2,3,4,null,5,6,null,null,7]
Output: 7

Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1

---

# Code
```go
package main

func main() {

}

// TreeNode 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findBottomLeftValue 找到二叉树最左下角的值
func findBottomLeftValue(root *TreeNode) int {
	// 调用 BFS 函数获取最左下角的节点
	res := BFS(root)
	// 返回该节点的值
	return res.Val
}

// BFS 执行广度优先搜索以找到最左下角的节点
func BFS(root *TreeNode) *TreeNode {
	// 初始化队列
	queue := []*TreeNode{}
	// 初始化新队列
	newQueue := []*TreeNode{}

	// 将根节点添加到队列
	queue = append(queue, root)

	// 无限循环，直到找到最左下角的节点
	for {
		// 遍历当前队列中的所有节点
		for i := 0; i < len(queue); i++ {
			// 如果当前节点有左子节点，则将其添加到新队列
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}

			// 如果当前节点有右子节点，则将其添加到新队列
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}

		// 如果新队列为空，说明当前队列中的第一个节点是最左下角的节点
		if len(newQueue) == 0 {
			return queue[0]
		}

		// 将新队列赋值给当前队列，并清空新队列
		queue = newQueue
		newQueue = []*TreeNode{}
	}
}
```