package main

import "math"

func main() {

}

func constrainedSubsetSum(nums []int, k int) int {
	maxSum := math.MinInt        // 初始化 maxSum 为最小整数，以便任何元素都能比它大。
	dp := make([]int, len(nums)) // dp 数组存储到当前索引为止的最大子序列和。
	myqueue := NewMaxQueue()     // 初始化一个 MaxQueue，用于维护 k 个元素的最大值。

	// 初始化 dp[0] 和更新 maxSum。
	if nums[0] > maxSum {
		maxSum = nums[0]
	}

	if nums[0] < 0 {
		dp[0] = 0
	} else {
		dp[0] = nums[0]
	}

	myqueue.Enqueue(dp[0]) // 将 dp[0] 加入队列。

	// 遍历 nums 数组更新 dp 值。
	for i := 1; i < len(dp); i++ {
		if myqueue.getLength() > k {
			myqueue.Dequeue() // 如果队列长度超过 k，则移除最旧的元素。
		}

		max := myqueue.ViewMax() // 获取当前队列中的最大值。

		if max < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = max + nums[i]
		}

		if dp[i] > maxSum {
			maxSum = dp[i] // 更新 maxSum。
		}

		myqueue.Enqueue(dp[i]) // 将当前 dp[i] 加入队列。
	}

	return maxSum // 返回最大子序列和。
}

// ---------------------------------------------------------------------------------------------------------
// MaxQueue 是一个支持常数时间内获取最大值的队列。
type MaxQueue struct {
	queue    []int // 主队列。
	maxQueue []int // 辅助队列，用来维持最大值。
}

// NewMaxQueue 初始化并返回一个 MaxQueue 实例。
func NewMaxQueue() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{},
	}
}

// Enqueue 将一个值加入队列，并更新 maxQueue 以保持最大值信息。
func (mq *MaxQueue) Enqueue(val int) {
	mq.queue = append(mq.queue, val)
	// 确保 maxQueue 的最后一个元素比 val 大。
	for len(mq.maxQueue) > 0 && mq.maxQueue[len(mq.maxQueue)-1] < val {
		mq.maxQueue = mq.maxQueue[:len(mq.maxQueue)-1]
	}
	mq.maxQueue = append(mq.maxQueue, val)
}

// Dequeue 移除队列最前端的值，并在必要时更新 maxQueue。
func (mq *MaxQueue) Dequeue() {
	poped := mq.queue[0]
	mq.queue = mq.queue[1:]
	// 如果被移除的元素是 maxQueue 的最大值，则也从 maxQueue 中移除。
	if poped == mq.maxQueue[0] {
		mq.maxQueue = mq.maxQueue[1:]
	}
}

// ViewMax 返回队列中的最大值。
func (mq *MaxQueue) ViewMax() int {
	return mq.maxQueue[0]
}

// getLength 返回 queue 的当前长度。
func (mq *MaxQueue) getLength() int {
	return len(mq.queue)
}
