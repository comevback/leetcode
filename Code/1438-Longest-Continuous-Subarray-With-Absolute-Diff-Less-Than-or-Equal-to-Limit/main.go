package main

import "fmt"

func main() {
	nums := []int{4, 2, 2, 2, 4, 4, 2, 2}
	res := longestSubarray(nums, 0)
	fmt.Println(res)
}

// 此函数用于确定最长的子数组，其中最小元素与最大元素之间的绝对差不超过给定限制。
// Function to determine the longest subarray where the absolute difference between
// the minimum and maximum elements does not exceed the given limit.
func longestSubarray(nums []int, limit int) int {
	longest := 0
	// 初始化一个新的MyStack对象。
	// Initialize a new MyStack object.
	mystack := NewMyStack()

	// 遍历输入切片中的元素。
	// Iterate over the elements of the input slice.
	for _, v := range nums {
		// 将当前元素推入栈中。
		// Push the current element onto the stack.
		mystack.Push(v)
		// 如果栈中的最大元素与最小元素之差超过限制，则持续弹出栈。
		// If the difference between the max and min elements in the stack exceeds the limit,
		// keep popping the stack.
		for mystack.ViewMax()-mystack.ViewMin() > limit {
			mystack.Pop()
		}

		// 更新目前找到的最长长度。
		// Update the longest length found so far.
		if mystack.getLength() > longest {
			longest = mystack.getLength()
		}
	}

	return longest
}

// MyStack是一个栈数据结构，增加了高效检索最小值和最大值的支持。
// MyStack represents a stack data structure with additional support for retrieving
// minimum and maximum values efficiently.
type MyStack struct {
	stack    []int
	minStack []int
	maxStack []int
}

// MyStack的构造函数初始化栈、minStack和maxStack的空切片。
// Constructor for MyStack initializes empty slices for the stack, minStack, and maxStack.
func NewMyStack() MyStack {
	return MyStack{
		stack:    []int{},
		minStack: []int{},
		maxStack: []int{},
	}
}

// Push向栈中添加新值，同时更新minStack和maxStack。
// Push adds a new value to the stack for updating the minStack and maxStack.
func (ms *MyStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	// 维护maxStack：从maxStack顶部移除小于新值的元素。
	// Maintain maxStack: remove elements from the top of maxStack that are less than the new value.
	for len(ms.maxStack) > 0 && ms.maxStack[len(ms.maxStack)-1] < val {
		ms.maxStack = ms.maxStack[:len(ms.maxStack)-1]
	}
	// 维护minStack：从minStack顶部移除大于新值的元素。
	// Maintain minStack: remove elements from the top of minStack that are greater than the new value.
	for len(ms.minStack) > 0 && ms.minStack[len(ms.minStack)-1] > val {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}

	ms.maxStack = append(ms.maxStack, val)
	ms.minStack = append(ms.minStack, val)
}

// Pop从栈前端移除元素，并相应地调整minStack和maxStack。
// Pop removes the element at the front of the stack, and adjusts minStack and maxStack accordingly.
func (ms *MyStack) Pop() {
	poped := ms.stack[0]
	ms.stack = ms.stack[1:]

	if poped == ms.maxStack[0] {
		ms.maxStack = ms.maxStack[1:]
	}

	if poped == ms.minStack[0] {
		ms.minStack = ms.minStack[1:]
	}
}

// getLength返回栈中当前的元素数量。
// getLength returns the current number of elements in the stack.
func (ms *MyStack) getLength() int {
	return len(ms.stack)
}

// ViewMax返回栈中当前的最大值。
// ViewMax returns the current maximum value in the stack.
func (ms *MyStack) ViewMax() int {
	return ms.maxStack[0]
}

// ViewMin返回栈中当前的最小值。
// ViewMin returns the current minimum value in the stack.
func (ms *MyStack) ViewMin() int {
	return ms.minStack[0]
}
