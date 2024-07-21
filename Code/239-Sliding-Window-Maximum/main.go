package main

import "fmt"

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	res := maxSlidingWindow(nums, 3)
	fmt.Println(res)
}

// maxSlidingWindow 函数计算每个大小为 k 的滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	// 初始化 maxQueue 实例
	myQueue := NewMaxQueue()
	for i := 0; i < len(nums); i++ {
		if i < k {
			// 前 k 个元素直接入队
			myQueue.Enqueue(nums[i])
			if i == k-1 {
				// 第一个窗口形成时记录最大值
				res = append(res, myQueue.ViewMax())
			}
		} else {
			// 新元素入队，移除窗口外的元素
			myQueue.Enqueue(nums[i])
			myQueue.Dequeue()
			// 记录当前窗口的最大值
			res = append(res, myQueue.ViewMax())
		}
	}
	return res
}

// ***********************************************  MaxQueue Implementation  ***************************************************
// maxQueue 结构体包含一个普通队列和一个双端队列用于维护最大值
type maxQueue struct {
	queue []int
	maxQ  []int
}

// NewMaxQueue 创建并返回一个新的 maxQueue 实例
func NewMaxQueue() maxQueue {
	return maxQueue{
		queue: []int{},
		maxQ:  []int{},
	}
}

// Enqueue 方法将一个新元素加入队列，同时更新 maxQ 保持递减顺序
func (mq *maxQueue) Enqueue(val int) {
	mq.queue = append(mq.queue, val)
	// 移除 maxQ 中所有小于当前值的元素
	for len(mq.maxQ) > 0 && mq.maxQ[len(mq.maxQ)-1] < val {
		mq.maxQ = mq.maxQ[:len(mq.maxQ)-1]
	}
	// 将新值加入 maxQ
	mq.maxQ = append(mq.maxQ, val)
}

// Dequeue 方法移除队列中的第一个元素，并更新 maxQ
func (mq *maxQueue) Dequeue() {
	if len(mq.queue) == 0 {
		return
	}
	// 移除队列中的第一个元素
	poped := mq.queue[0]
	mq.queue = mq.queue[1:]
	// 如果移除的元素是 maxQ 中的最大值，也移除它
	if poped == mq.maxQ[0] {
		mq.maxQ = mq.maxQ[1:]
	}
}

// ViewMax 方法返回当前队列中的最大值
func (mq *maxQueue) ViewMax() int {
	if len(mq.maxQ) == 0 {
		return -1 // 如果队列为空，返回一个默认值
	}
	return mq.maxQ[0]
}

// ViewTop 方法返回队列中的第一个元素
func (mq *maxQueue) ViewTop() int {
	if len(mq.queue) == 0 {
		return -1 // 如果队列为空，返回一个默认值
	}
	return mq.queue[0]
}
