package main

import "fmt"

func main() {
	height := []int{2, 0, 2}
	res := trap(height)
	fmt.Println(res) // 输出应为 2，因为中间的低洼可以存 2 单位的水
}

// trap 函数计算在一系列直方图（表示高度）中可以“困住”多少雨水。
func trap(height []int) int {
	curHeight := 0 // 当前遇到的最高高度
	sum := 0       // 存储总共可以困住的水量
	tempSum := 0   // 临时存储在当前最高高度之前困住的水量

	// 第一次从左到右遍历，计算可能困住的水
	for _, v := range height {
		if v >= curHeight { // 如果当前柱子高度大于或等于目前遇到的最高高度
			sum += tempSum // 将当前积累的临时水量加到总水量中
			tempSum = 0    // 重置临时水量
			curHeight = v  // 更新当前的最高高度
		} else {
			tempSum += curHeight - v // 计算当前柱子顶部到当前最高高度的水量，并累加到临时水量
		}
	}

	// 重置变量以从右到左再次遍历
	curHeight = 0
	tempSum = 0
	// 第二次从右到左遍历，校正之前可能的过估计
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > curHeight { // 如果当前柱子高于目前遇到的最高高度
			sum += tempSum        // 加上从这个方向计算的临时水量
			tempSum = 0           // 重置临时水量
			curHeight = height[i] // 更新当前最高高度
		} else {
			tempSum += curHeight - height[i] // 计算并累加临时水量
		}
	}

	return sum // 返回总共困住的水量
}

func trap1(height []int) int {

}

type MaxQueue struct {
	queue    []int
	maxQueue []int
}

func NewMaxQueue() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{},
	}
}

func (mq *MaxQueue) Enqueue(value int) {
	mq.queue = append(mq.queue, value)
	for len(mq.maxQueue) > 0 && value > mq.maxQueue[len(mq.maxQueue)-1] {
		mq.maxQueue = mq.maxQueue[:len(mq.maxQueue)-1]
	}

}

func (mq *MaxQueue) Dequeue() {

}

func (mq *MaxQueue) ViewMax() int {
	return mq.maxQueue[0]
}
