# 1602. 找到二叉树中最近的右侧节点
给定一棵二叉树的根节点 root 和树中的一个节点 u ，返回与 u 所在层中距离最近的右侧节点，当 u 是所在层中最右侧的节点，返回 null 。

示例 1：

输入：root = [1,2,3,null,4,5,6], u = 4
输出：5
解释：节点 4 所在层中，最近的右侧节点是节点 5。
示例 2：

输入：root = [3,null,4,2], u = 2
输出：null
解释：2 的右侧没有节点。
示例 3：

输入：root = [1], u = 1
输出：null
示例 4：

输入：root = [3,4,2,null,null,null,1], u = 4
输出：2
提示:

树中节点个数的范围是 [1, 105] 。
1 <= Node.val <= 105
树中所有节点的值是唯一的。
u 是以 root 为根的二叉树的一个节点。

---

# Code
```go
package main

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findNearestRightNode 查找目标节点 u 的最近右侧邻居节点
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	return BFS(root, u)
}

// BFS 使用广度优先搜索（层次遍历）来查找目标节点的最近右侧邻居节点
func BFS(root *TreeNode, target *TreeNode) *TreeNode {
	queue := []*TreeNode{}      // 当前层节点队列
	newQueue := []*TreeNode{}   // 下一层节点队列
	queue = append(queue, root) // 将根节点添加到当前层队列

	for {
		if len(queue) == 0 && len(newQueue) == 0 {
			break // 如果当前层和下一层队列都为空，结束搜索
		}

		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left) // 将左子节点添加到下一层队列
			}

			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right) // 将右子节点添加到下一层队列
			}

			if queue[i] == target {
				if i+1 < len(queue) {
					return queue[i+1] // 返回目标节点的右侧邻居节点
				}
			}
		}

		queue = newQueue         // 更新当前层队列为下一层队列
		newQueue = []*TreeNode{} // 重置下一层队列
	}

	return nil // 如果没有找到右侧邻居节点，返回 nil
}
```