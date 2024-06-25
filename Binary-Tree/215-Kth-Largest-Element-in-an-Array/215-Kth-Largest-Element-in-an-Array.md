# 215. Kth Largest Element in an Array
Solved
Medium
Topics
Companies
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting? 

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104

---

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	res := findKthLargest1(nums, 4)
	fmt.Println(res) // 输出: 4
}

// ****************************************************  Quick Sort  ************************************************
// 查找第 k 小的元素
func findKthSmallest(nums []int, k int) int {
	pivot := 0
	tail := len(nums) - 1

	// 快速排序划分过程
	for pivot < tail {
		if nums[pivot] > nums[tail] {
			if pivot == tail-1 {
				nums[pivot], nums[tail] = nums[tail], nums[pivot]
			} else {
				nums[pivot], nums[pivot+1], nums[tail] = nums[tail], nums[pivot], nums[pivot+1]
			}
			pivot += 1
		} else {
			tail -= 1
		}
	}

	// 根据 pivot 位置决定是递归左半部分还是右半部分
	if pivot == k-1 {
		return nums[pivot]
	} else if pivot > k-1 {
		return findKthLargest(nums[:pivot], k)
	} else {
		return findKthLargest(nums[pivot+1:], k-pivot-1)
	}
}

// 查找第 k 大的元素
func findKthLargest(nums []int, k int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		p := partion(nums, lo, hi)

		// 根据 pivot 位置决定是继续在左半部分还是右半部分查找
		if p == k-1 {
			return nums[p]
		} else if p > k-1 {
			hi = p - 1
		} else {
			lo = p + 1
		}
	}
	return -1
}

// 快速排序划分函数
func partion(nums []int, lo int, hi int) int {
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] > pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	// 将 pivot 放置在正确位置
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

// ****************************************************  Priority Queue  *********************************************
// 使用优先队列查找第 k 大的元素
func findKthLargest1(nums []int, k int) int {
	pq := newPQ()

	for _, v := range nums {
		// 当优先队列大小小于等于 k 时，直接入队
		if len(pq.queue) <= k {
			pq.Enqueue(v)
		} else if len(pq.queue) > k {
			// 队列满了之后，只有当新元素大于队首元素时才入队
			if pq.queue[1] >= v {
				continue
			} else {
				pq.Enqueue(v)
				pq.Dequeue()
			}
		}
	}

	return pq.queue[1]
}

// 优先队列结构
type PQ struct {
	queue []int
}

// 创建新的优先队列
func newPQ() *PQ {
	return &PQ{
		queue: []int{0},
	}
}

// 上浮操作，确保堆的性质
func (pq *PQ) swim(index int) {
	for {
		parent := index / 2

		if parent < 1 {
			break
		}

		if pq.queue[parent] > pq.queue[index] {
			pq.queue[parent], pq.queue[index] = pq.queue[index], pq.queue[parent]
			index = parent
		} else {
			break
		}
	}
}

// 下沉操作，确保堆的性质
func (pq *PQ) sink(index int) {
	for {
		left := index * 2
		right := index*2 + 1
		min := index

		// 找到子节点中的最小值
		if left < len(pq.queue) && pq.queue[min] > pq.queue[left] {
			min = left
		}

		if right < len(pq.queue) && pq.queue[min] > pq.queue[right] {
			min = right
		}

		if min == index {
			break
		}

		// 交换并继续下沉
		pq.queue[min], pq.queue[index] = pq.queue[index], pq.queue[min]
		index = min
	}
}

// 入队操作
func (pq *PQ) Enqueue(val int) {
	if len(pq.queue) == 0 {
		pq.queue = append(pq.queue, 0)
	}

	pq.queue = append(pq.queue, val)
	pq.swim(len(pq.queue) - 1)
}

// 出队操作
func (pq *PQ) Dequeue() int {
	if len(pq.queue) <= 1 {
		return -1
	}

	temp := pq.queue[1]
	pq.queue[1], pq.queue[len(pq.queue)-1] = pq.queue[len(pq.queue)-1], pq.queue[1]
	pq.queue = pq.queue[:len(pq.queue)-1]
	pq.sink(1)
	return temp
}
```