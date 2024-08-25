# 703. Kth Largest Element in a Stream（PQ）

Solved
Easy
Topics
Companies
You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

You are tasked to implement a class which, for a given integer k, maintains a stream of test scores and continuously returns the kth highest test score after a new score has been submitted. More specifically, we are looking for the kth highest score in the sorted list of all scores.

Implement the KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
int add(int val) Adds a new test score val to the stream and returns the element representing the kth largest element in the pool of test scores so far.

Example 1:

Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

Output: [null, 4, 5, 5, 8, 8]

Explanation:

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3); // return 4
kthLargest.add(5); // return 5
kthLargest.add(10); // return 5
kthLargest.add(9); // return 8
kthLargest.add(4); // return 8

Example 2:

Input:
["KthLargest", "add", "add", "add", "add"]
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

Output: [null, 7, 7, 7, 8]

Explanation:

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
kthLargest.add(2); // return 7
kthLargest.add(10); // return 7
kthLargest.add(9); // return 7
kthLargest.add(9); // return 8

Constraints:

1 <= nums.length <= 104
1 <= k <= nums.length + 1
-104 <= nums[i] <= 104
-104 <= val <= 104
At most 104 calls will be made to add.

---

# Code

```go
package main

func main() {

}

// KthLargest 是一个维护流中第 k 大元素的结构
type KthLargest struct {
	pq    *PQ // 优先队列，用来维护最小堆
	limit int // 堆的最大容量
}

// Constructor 初始化 KthLargest 对象
func Constructor(k int, nums []int) KthLargest {
	pq := NewPQ(k) // 创建一个优先队列（最小堆）
	for _, v := range nums {
		if len(pq.queue) < k+1 {
			pq.push(v) // 向堆中添加元素
		} else {
			if v > pq.queue[1] {
				pq.queue[1] = v // 如果新元素大于堆顶元素，则替换堆顶
				pq.sink(1)      // 下沉调整堆
			}
		}
	}

	return KthLargest{
		pq:    &pq,
		limit: k,
	}
}

// Add 添加新的值到流中，并返回当前第 k 大的元素
func (this *KthLargest) Add(val int) int {
	if len(this.pq.queue) < this.limit+1 {
		this.pq.push(val) // 如果堆未满，直接添加
	} else {
		if val > this.pq.queue[1] {
			this.pq.queue[1] = val // 如果新值大于堆顶，替换并调整堆
			this.pq.sink(1)
		}
	}
	return this.pq.queue[1] // 返回堆顶元素，即第 k 大的元素
}

// PQ 是一个优先队列，实现最小堆的功能
type PQ struct {
	queue []int // 使用切片来存储堆元素
}

// NewPQ 创建并初始化一个优先队列
func NewPQ(k int) PQ {
	return PQ{
		queue: []int{}, // 初始化空切片
	}
}

// swim 上浮调整，保持最小堆的性质
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

// sink 下沉调整，保持最小堆的性质
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

		pq.queue[index], pq.queue[min] = pq.queue[min], pq.queue[index] // 交换元素
		index = min
	}
}

// push 向堆中添加新元素，并进行上浮调整
func (pq *PQ) push(val int) {
	if len(pq.queue) == 0 {
		pq.queue = append(pq.queue, 0) // 添加一个哨兵元素，简化计算
	}

	pq.queue = append(pq.queue, val)
	pq.swim(len(pq.queue) - 1)
}

// pop 弹出堆顶元素，并进行下沉调整
func (pq *PQ) pop() int {
	if len(pq.queue) < 2 {
		return -1 // 如果堆为空，返回 -1
	}

	poped := pq.queue[1]                                                            // 保存堆顶元素
	pq.queue[1], pq.queue[len(pq.queue)-1] = pq.queue[len(pq.queue)-1], pq.queue[1] // 交换堆顶与最后一个元素
	pq.queue = pq.queue[:len(pq.queue)-1]                                           // 移除最后一个元素
	pq.sink(1)                                                                      // 下沉调整

	return poped
}
```
