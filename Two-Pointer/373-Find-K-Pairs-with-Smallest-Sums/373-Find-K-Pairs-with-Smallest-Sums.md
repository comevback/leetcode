# 373. Find K Pairs with Smallest Sums
Solved
Medium
Topics
Companies
You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

Define a pair (u, v) which consists of one element from the first array and one element from the second array.

Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.

Example 1:
> Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]]
Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

Example 2:
> Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [[1,1],[1,1]]
Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 

Constraints:

1 <= nums1.length, nums2.length <= 105
-109 <= nums1[i], nums2[i] <= 109
nums1 and nums2 both are sorted in non-decreasing order.
1 <= k <= 104
k <= nums1.length * nums2.length

---

# Code
```go
package main

func main() {
	arr1 := []int{1, 2, 4, 5, 6}
	arr2 := []int{3, 5, 7, 9}
	k := 20

	res := kSmallestPairs(arr1, arr2, k)
	for _, v := range res {
		println(v[0], v[1])
	}
}

// kSmallestPairs 返回前 k 个最小和的数对
// 使用优先队列（最小堆）来高效地找到和最小的前 k 个数对。主要步骤如下：

// 初始化优先队列：将 nums1 中的每个元素与 nums2 的第一个元素配对，并插入优先队列。初始时的优先队列包含 len(nums1) 个元素。
// 每次从优先队列中弹出一个数对，记录到结果列表中。
// 对于弹出的数对 (nums1[i], nums2[j])，插入新的数对 (nums1[i], nums2[j+1]) 到优先队列。
// 重复上述过程，直到找到 k 个数对。
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := NewPQ()

	// 初始化优先队列，将 nums1 中每个元素与 nums2 的第一个元素配对并插入优先队列
	for i, v := range nums1 {
		newItem := Item{
			Arr:   []int{v, nums2[0]}, // 数对
			Val:   v + nums2[0],       // 数对的和
			Index: []int{i, 0},        // 数对中元素的索引
		}
		pq.push(newItem)
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		// 从优先队列中弹出最小的数对
		poped := pq.pop()
		res = append(res, poped.Arr)

		index1, index2 := poped.Index[0], poped.Index[1]

		// 检查是否可以插入新的数对
		if index2+1 < len(nums2) {
			newItem := Item{}
			newItem.Arr = []int{nums1[index1], nums2[index2+1]}
			newItem.Val = nums1[index1] + nums2[index2+1]
			newItem.Index = []int{index1, index2 + 1}
			pq.push(newItem)
		}
	}

	return res
}

// Item 定义了优先队列中的元素
type Item struct {
	Arr   []int // 数对
	Val   int   // 数对的和
	Index []int // 数对中元素的索引
}

// PQ 实现了一个优先队列
type PQ []Item

// NewPQ 创建一个新的优先队列
func NewPQ() *PQ {
	return &PQ{}
}

// swim 执行上浮操作，以维持堆的性质
func (pq *PQ) swim(index int) {
	parent := index / 2

	for parent >= 1 && (*pq)[parent].Val > (*pq)[index].Val {
		(*pq)[parent], (*pq)[index] = (*pq)[index], (*pq)[parent]
		index = parent
		parent = index / 2
	}
}

// sink 执行下沉操作，以维持堆的性质
func (pq *PQ) sink(index int) {
	for {
		left := 2 * index
		right := 2*index + 1
		min := index

		// 找到子节点中较小的节点
		if left < len(*pq) && (*pq)[left].Val < (*pq)[min].Val {
			min = left
		}

		if right < len(*pq) && (*pq)[right].Val < (*pq)[min].Val {
			min = right
		}

		if min == index {
			break
		}

		(*pq)[index], (*pq)[min] = (*pq)[min], (*pq)[index]
		index = min
	}
}

// push 将新元素插入优先队列
func (pq *PQ) push(item Item) {
	if len(*pq) == 0 {
		(*pq) = append((*pq), Item{}) // 占位符，使索引从1开始
	}

	(*pq) = append((*pq), item)
	pq.swim(len(*pq) - 1)
}

// pop 从优先队列中弹出最小的元素
func (pq *PQ) pop() Item {
	if len(*pq) == 0 {
		var null Item
		return null
	}

	poped := (*pq)[1]
	(*pq)[1] = (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	pq.sink(1)
	return poped
}
```