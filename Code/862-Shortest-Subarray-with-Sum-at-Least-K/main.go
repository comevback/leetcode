package main

import (
	"fmt"
	"math"
)

func main() {
	// nums := []int{-34, 37, 51, 3, -12, -50, 51, 100, -47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6}
	nums := []int{2, -1, 2}
	res := shortestSubarray(nums, 3)
	fmt.Println(res)
}

func shortestSubarray(nums []int, k int) int {
	minLength := math.MaxInt   // 初始化最小长度为最大整数
	hsMap := make(map[int]int) // 用于存储前缀和及其索引
	hsMap[0] = -1              // 初始前缀和
	sum := 0
	mystack := NewMyStack() // 创建一个新的栈
	mystack.Push(0)         // 推入初始值

	// 遍历数组，计算前缀和
	for i, v := range nums {
		sum += v
		hsMap[sum] = i
		mystack.Push(sum)
		// 当栈中最大值与最小值之差大于等于 k 时，尝试更新最短长度
		for mystack.ViewMax()-mystack.ViewMin() >= k {
			mx := mystack.ViewMax()
			mi := mystack.ViewMin()
			if hsMap[mx]-hsMap[mi] > 0 && hsMap[mx]-hsMap[mi] < minLength {
				minLength = hsMap[mx] - hsMap[mi]
			}
			mystack.Pop()
		}
	}

	// 如果没有找到符合条件的子数组，返回 -1
	if minLength == math.MaxInt {
		return -1
	} else {
		return minLength
	}
}

// -----------------------------------------------------------------------------------------------------------------------
// MyStack 定义了一个栈结构，支持最大值和最小值的快速检索
// MyStack defines a stack structure with support for fast retrieval of maximum and minimum values
type MyStack struct {
	stack    []int
	maxStack []int
	minStack []int
}

// NewMyStack 创建一个新的 MyStack 实例
// NewMyStack creates a new instance of MyStack
func NewMyStack() MyStack {
	return MyStack{
		stack:    []int{},
		maxStack: []int{},
		minStack: []int{},
	}
}

// Push 将值压入栈中，并更新最大栈和最小栈
// Push adds a value to the stack and updates the max and min stacks accordingly
func (mq *MyStack) Push(val int) {
	mq.stack = append(mq.stack, val)
	for len(mq.maxStack) > 0 && mq.maxStack[len(mq.maxStack)-1] < val {
		mq.maxStack = mq.maxStack[:len(mq.maxStack)-1]
	}
	for len(mq.minStack) > 0 && mq.minStack[len(mq.minStack)-1] > val {
		mq.minStack = mq.minStack[:len(mq.minStack)-1]
	}
	mq.maxStack = append(mq.maxStack, val)
	mq.minStack = append(mq.minStack, val)
}

// Pop 从栈中移除元素，并适当更新最大栈和最小栈
// Pop removes an element from the stack and updates the max and min stacks appropriately
func (mq *MyStack) Pop() {
	poped := mq.stack[0]
	mq.stack = mq.stack[1:]
	if poped == mq.maxStack[0] {
		mq.maxStack = mq.maxStack[1:]
	}
	if poped == mq.minStack[0] {
		mq.minStack = mq.minStack[1:]
	}
}

// ViewMax 返回栈中的最大值
// ViewMax returns the maximum value in the stack
func (mq *MyStack) ViewMax() int {
	max := mq.maxStack[0]
	return max
}

// ViewMin 返回栈中的最小值
// ViewMin returns the minimum value in the stack
func (mq *MyStack) ViewMin() int {
	min := mq.minStack[0]
	return min
}

// getLength 返回栈的长度
// getLength returns the length of the stack
func (mq *MyStack) getLength() int {
	res := len(mq.stack)
	return res
}
