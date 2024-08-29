# 2583. Kth Largest Sum in a Binary Tree

Solved
Medium
Topics
Companies
Hint
You are given the root of a binary tree and a positive integer k.

The level sum in the tree is the sum of the values of the nodes that are on the same level.

Return the kth largest level sum in the tree (not necessarily distinct). If there are fewer than k levels in the tree, return -1.

Note that two nodes are on the same level if they have the same distance from the root.

Example 1:

Input: root = [5,8,9,2,1,3,7,4,6], k = 2
Output: 13
Explanation: The level sums are the following:

- Level 1: 5.
- Level 2: 8 + 9 = 17.
- Level 3: 2 + 1 + 3 + 7 = 13.
- Level 4: 4 + 6 = 10.
  The 2nd largest level sum is 13.
  Example 2:

Input: root = [1,2,null,3], k = 1
Output: 3
Explanation: The largest level sum is 3.

Constraints:

The number of nodes in the tree is n.
2 <= n <= 105
1 <= Node.val <= 106
1 <= k <= n

---

# Code

```go
package main

func main() {

}

// TreeNode 定义了二叉树的节点结构。
type TreeNode struct {
	Val   int       // 当前节点的值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

// kthLargestLevelSum 使用广度优先搜索（BFS）来计算并返回树中第 k 大的层次总和。
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{} // 使用切片作为队列以支持 BFS
	pq := NewPQ()          // 创建一个新的优先队列作为最小堆
	queue = append(queue, root)

	// 遍历树的每一层
	for len(queue) != 0 {
		temp := 0
		length := len(queue)
		for i := 0; i < length; i++ {
			// 将当前节点的左右子节点添加到队列中
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			temp += queue[i].Val // 计算当前层的总和
		}

		// 维护一个大小为 k 的最小堆，以保留最大的 k 个层次总和
		if len(pq.queue) < k+1 {
			pq.push(temp)
		} else {
			if temp > pq.queue[1] {
				pq.queue[1] = temp
				pq.sink(1)
			}
		}

		queue = queue[length:] // 移除已处理的层的节点
	}

	// 如果堆中的元素少于 k 个，则返回 -1
	if len(pq.queue) < k+1 {
		return -1
	} else {
		return int64(pq.queue[1]) // 返回第 k 大的层次总和
	}
}

// ==============================================================================================================================

// PQ 实现了一个最小堆
type PQ struct {
	queue []int // 存储堆元素的切片
}

// NewPQ 创建并初始化一个新的优先队列
func NewPQ() PQ {
	return PQ{
		queue: []int{0}, // 初始化包含哨兵元素的切片以简化索引计算
	}
}

// swim 实现堆的上浮调整，以维持最小堆的性质
func (pq *PQ) swim(index int) {
	for {
		parent := index / 2

		if parent >= 1 && pq.queue[parent] > pq.queue[index] {
			pq.queue[parent], pq.queue[index] = pq.queue[index], pq.queue[parent] // 交换元素
			index = parent
		} else {
			break
		}
	}
}

// sink 实现堆的下沉调整，以维持最小堆的性质
func (pq *PQ) sink(index int) {
	for {
		left := index * 2
		right := index*2 + 1
		min := index

		if left < len(pq.queue) && pq.queue[left] < pq.queue[min] {
			min = left
		}

		if right < len(pq.queue) && pq.queue[right] < pq.queue[min] {
			min = right
		}

		if min == index {
			break
		}

		pq.queue[min], pq.queue[index] = pq.queue[index], pq.queue[min] // 交换元素
		index = min
	}
}

// push 向最小堆中添加新元素，并进行上浮调整
func (pq *PQ) push(val int) {
	if len(pq.queue) < 1 {
		pq.queue = append(pq.queue, 0) // 确保哨兵元素存在
	}

	pq.queue = append(pq.queue, val)
	pq.swim(len(pq.queue) - 1)
}

// pop 从最小堆中移除堆顶元素，并进行下沉调整
func (pq *PQ) pop() int {
	if len(pq.queue) < 2 {
		return -1 // 如果堆为空，返回 -1
	}

	poped := pq.queue[1]
	pq.queue[1], pq.queue[len(pq.queue)-1] = pq.queue[len(pq.queue)-1], pq.queue[1] // 交换堆顶和最后一个元素
	pq.queue = pq.queue[:len(pq.queue)-1]                                           // 移除最后一个元素
	pq.sink(1)
	return poped
}
```
