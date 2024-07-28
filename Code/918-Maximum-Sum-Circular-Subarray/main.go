package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-3, -2, -3}
	res := maxSubarraySumCircular(nums)
	fmt.Println(res) // 打印计算得到的循环子数组的最大和
}

// maxSubarraySumCircular 计算循环数组的最大子数组和
// maxSubarraySumCircular calculates the maximum subarray sum for a circular array
func maxSubarraySumCircular(nums []int) int {
	numLen := len(nums)
	biggestSum := math.MinInt         // 初始化最大和为最小整数
	prefix := make([]int, 2*numLen+1) // 创建前缀和数组，长度为原数组的两倍加一
	prefix[0] = 0                     // 初始化前缀和的第一个元素为0

	// 计算扩展数组的前缀和
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + nums[(i-1)%numLen]
	}

	mystack := NewMinStack() // 创建一个新的最小栈
	mystack.Push(0)          // 推入初始前缀和的索引
	for i := 1; i < len(prefix); i++ {
		temp := prefix[i] - mystack.ViewMin() // 计算当前最大可能子数组和
		if temp > biggestSum {
			biggestSum = temp // 更新最大子数组和
		}
		if mystack.getLength() == numLen {
			mystack.Pop() // 如果栈的长度达到数组长度，移除最老的元素
		}
		mystack.Push(prefix[i]) // 推入当前前缀和
	}

	return biggestSum // 返回计算得到的最大子数组和
}

// ------------------------------------------------------------------------------------------------------------------------

// MinStack 是一个支持查找最小值的栈
// MinStack is a stack that supports finding the minimum value
type MinStack struct {
	stack    []int
	minStack []int
}

// NewMinStack 创建一个新的 MinStack 实例
// NewMinStack creates a new instance of MinStack
func NewMinStack() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push 将值压入栈中，并更新最小栈
// Push adds a value to the stack and updates the min stack accordingly
func (mq *MinStack) Push(val int) {
	mq.stack = append(mq.stack, val)
	for len(mq.minStack) > 0 && mq.minStack[len(mq.minStack)-1] > val {
		mq.minStack = mq.minStack[:len(mq.minStack)-1]
	}
	mq.minStack = append(mq.minStack, val)
}

// Pop 从栈中移除元素，并适当更新最小栈
// Pop removes an element from the stack and updates the min stack appropriately
func (mq *MinStack) Pop() {
	poped := mq.stack[0]
	mq.stack = mq.stack[1:]
	if poped == mq.minStack[0] {
		mq.minStack = mq.minStack[1:]
	}
}

// ViewMin 返回栈中的最小值
// ViewMin returns the minimum value in the stack
func (mq *MinStack) ViewMin() int {
	return mq.minStack[0]
}

// getLength 返回栈的长度
// getLength returns the length of the stack
func (mq *MinStack) getLength() int {
	return len(mq.stack)
}
