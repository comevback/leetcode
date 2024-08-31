# 2415. Reverse Odd Levels of Binary Tree

Solved
Medium
Topics
Companies
Hint
Given the root of a perfect binary tree, reverse the node values at each odd level of the tree.

For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then it should become [18,29,11,7,4,3,1,2].
Return the root of the reversed tree.

A binary tree is perfect if all parent nodes have two children and all leaves are on the same level.

The level of a node is the number of edges along the path between it and the root node.

Example 1:

Input: root = [2,3,5,8,13,21,34]
Output: [2,5,3,8,13,21,34]
Explanation:
The tree has only one odd level.
The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.
Example 2:

Input: root = [7,13,11]
Output: [7,11,13]
Explanation:
The nodes at level 1 are 13, 11, which are reversed and become 11, 13.
Example 3:

Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
Explanation:
The odd levels have non-zero values.
The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.

Constraints:

The number of nodes in the tree is in the range [1, 214].
0 <= Node.val <= 105
root is a perfect binary tree.

---

# Code

```go
package main

func main() {

}

// TreeNode 定义了二叉树节点的结构。
type TreeNode struct {
	Val   int       // 当前节点的值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

// reverseOddLevels 反转二叉树中所有奇数层的节点值。
func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{}      // 使用队列来进行广度优先搜索
	newQueue := []*TreeNode{}   // 用于存储下一层的节点
	queue = append(queue, root) // 初始化队列
	index := 0                  // 用于跟踪当前层的层数

	// 当队列不为空时，循环处理每一层
	for len(queue) != 0 {
		// 遍历当前层的每个节点，将子节点添加到 newQueue
		for _, v := range queue {
			if v.Left != nil {
				newQueue = append(newQueue, v.Left)
			}
			if v.Right != nil {
				newQueue = append(newQueue, v.Right)
			}
		}
		// 如果当前层是奇数层（注意：层数从 0 开始，因此奇数层实际上是偶数索引），则反转该层的节点值
		if index%2 != 0 {
			reverse(queue)
		}
		// 更新队列为下一层的节点，清空 newQueue，增加层数索引
		queue = newQueue
		newQueue = []*TreeNode{}
		index += 1
	}

	return root // 返回修改后的根节点
}

// reverse 用于反转一个 TreeNode 切片中的节点值。
func reverse(arr []*TreeNode) {
	if len(arr) < 2 {
		return // 如果数组长度小于 2，不需要反转
	}

	head, tail := 0, len(arr)-1

	// 使用两个指针，一个从开始一个从末尾，交换这两个指针指向的节点的值
	for head < tail {
		arr[head].Val, arr[tail].Val = arr[tail].Val, arr[head].Val
		head += 1
		tail -= 1
	}
}
```
