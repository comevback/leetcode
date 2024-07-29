package main

import "fmt"

func main() {
	nums := []int{1, -1, -2, 4, -7, 3}
	res := maxResult(nums, 2)
	fmt.Println(res) // 打印最终结果
}

// maxResult 计算可以从数组的第一个元素跳到最后一个元素的最大得分
// maxResult calculates the maximum result you can achieve by jumping from the first element to the last element
func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums)) // 动态规划数组，dp[i] 存储到达 i 位置的最大得分
	dp[0] = nums[0]              // 初始位置的得分就是数组的第一个元素
	myqueue := NewMaxQueue()     // 初始化最大队列
	myqueue.Enqueue(dp[0])       // 将初始得分入队

	for i := 1; i < len(dp); i++ {
		if myqueue.getLength() > k { // 保证队列中元素数量不超过 k，满足跳跃限制
			myqueue.Dequeue() // 如果超过了，就从队列中移除最旧的元素
		}
		dp[i] = myqueue.ViewMax() + nums[i] // 计算当前位置的最大得分
		myqueue.Enqueue(dp[i])              // 将新计算的得分入队
	}

	return dp[len(dp)-1] // 返回到达数组最后一个位置的最大得分
}

// --------------------------------------------------------------------------------------------------------
// MaxQueue 是一个支持快速获取最大值的队列
// MaxQueue is a queue that supports fast retrieval of the maximum value
type MaxQueue struct {
	queue    []int // 普通队列
	maxQueue []int // 用于维护队列最大值的单调递减队列
}

// NewMaxQueue 初始化并返回一个 MaxQueue 实例
// NewMaxQueue initializes and returns an instance of MaxQueue
func NewMaxQueue() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{},
	}
}

// Enqueue 将一个值加入到队列中，并维护最大队列的性质
// Enqueue adds a value to the queue and maintains the properties of the max queue
func (mq *MaxQueue) Enqueue(val int) {
	mq.queue = append(mq.queue, val)
	// 确保 maxQueue 是单调递减的
	for len(mq.maxQueue) > 0 && mq.maxQueue[len(mq.maxQueue)-1] < val {
		mq.maxQueue = mq.maxQueue[:len(mq.maxQueue)-1]
	}
	mq.maxQueue = append(mq.maxQueue, val)
}

// Dequeue 从队列中移除元素，并更新最大队列
// Dequeue removes an element from the queue and updates the max queue accordingly
func (mq *MaxQueue) Dequeue() {
	poped := mq.queue[0]
	mq.queue = mq.queue[1:]
	if poped == mq.maxQueue[0] {
		mq.maxQueue = mq.maxQueue[1:]
	}
}

// ViewMax 返回队列中的最大值
// ViewMax returns the maximum value in the queue
func (mq *MaxQueue) ViewMax() int {
	return mq.maxQueue[0]
}

// getLength 返回队列的当前长度
// getLength returns the current length of the queue
func (mq *MaxQueue) getLength() int {
	return len(mq.queue)
}
