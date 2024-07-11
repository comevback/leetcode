# 1609. Even Odd Tree

Solved
Medium
Topics
Companies
Hint
A binary tree is named Even-Odd if it meets the following conditions:

The root of the binary tree is at level index 0, its children are at level index 1, their children are at level index 2, etc.
For every even-indexed level, all nodes at the level have odd integer values in strictly increasing order (from left to right).
For every odd-indexed level, all nodes at the level have even integer values in strictly decreasing order (from left to right).
Given the root of a binary tree, return true if the binary tree is Even-Odd, otherwise return false.

Example 1:

Input: root = [1,10,4,3,null,7,9,12,8,6,null,null,2]
Output: true
Explanation: The node values on each level are:
Level 0: [1]
Level 1: [10,4]
Level 2: [3,7,9]
Level 3: [12,8,6,2]
Since levels 0 and 2 are all odd and increasing and levels 1 and 3 are all even and decreasing, the tree is Even-Odd.
Example 2:

Input: root = [5,4,2,3,3,7]
Output: false
Explanation: The node values on each level are:
Level 0: [5]
Level 1: [4,2]
Level 2: [3,3,7]
Node values in level 2 must be in strictly increasing order, so the tree is not Even-Odd.
Example 3:

Input: root = [5,9,1,3,5,7]
Output: false
Explanation: Node values in the level 1 should be even integers.

Constraints:

The number of nodes in the tree is in the range [1, 105].
1 <= Node.val <= 106

---

# Code

```go
package main

import "math"

func main() {
	// 示例用例
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 12}

	result := isEvenOddTree(root)
	println(result) // 输出 true 或 false
}

// TreeNode 结构体表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isEvenOddTree 判断二叉树是否为偶数奇数树
func isEvenOddTree(root *TreeNode) bool {
	// 如果根节点为空，返回 false
	if root == nil {
		return false
	}
	// 使用切片作为队列进行层序遍历
	queue := []*TreeNode{}
	// 将根节点加入队列
	queue = append(queue, root)
	// 标记当前层是否为偶数层
	even := true
	// 定义前一个节点值和当前节点值
	var pre, cur int

	// 开始层序遍历
	for len(queue) != 0 {
		// 根据当前层是否为偶数层初始化 pre 和 cur
		if even {
			pre, cur = math.MinInt, math.MinInt
		} else {
			pre, cur = math.MaxInt, math.MaxInt
		}
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
			// 更新前一个节点值
			pre = cur
			// 更新当前节点值
			cur = queue[i].Val
			// 判断当前层的节点值是否满足条件
			if even {
				// 偶数层：节点值必须是奇数且严格递增
				if queue[i].Val%2 == 0 || pre >= cur {
					return false
				}
			} else {
				// 奇数层：节点值必须是偶数且严格递减
				if queue[i].Val%2 != 0 || pre <= cur {
					return false
				}
			}
		}
		// 切换到下一层
		even = !even
		// 移除当前层的节点
		queue = queue[length:]
	}
	// 如果所有层的节点值都满足条件，返回 true
	return true
}
```
