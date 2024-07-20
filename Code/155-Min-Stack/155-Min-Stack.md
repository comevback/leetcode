# 155. Min Stack

Solved
Medium
Topics
Companies
Hint
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top(); // return 0
minStack.getMin(); // return -2

Constraints:

-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 \* 104 calls will be made to push, pop, top, and getMin.

---

# Code

```go
package main

func main() {
	// obj := Constructor();
	// obj.Push(val);
	// obj.Pop();
	// param_3 := obj.Top();
	// param_4 := obj.GetMin();
}

// MinStack 结构体包含两个栈：一个用于存储数据，另一个用于存储最小值
type MinStack struct {
	queue    []int // 存储数据的栈
	minQueue []int // 存储最小值的栈
}

// Constructor 返回一个新的 MinStack 实例
func Constructor() MinStack {
	return MinStack{}
}

// Push 方法将一个元素压入栈中，并更新最小值栈
func (this *MinStack) Push(val int) {
	this.queue = append(this.queue, val) // 将元素压入数据栈
	// 如果最小值栈为空，或者新元素小于等于最小值栈顶元素，则将新元素压入最小值栈
	if len(this.minQueue) == 0 || val < this.minQueue[len(this.minQueue)-1] {
		this.minQueue = append(this.minQueue, val)
	} else {
		// 否则将最小值栈顶元素再次压入最小值栈，保持栈的长度一致
		this.minQueue = append(this.minQueue, this.minQueue[len(this.minQueue)-1])
	}
}

// Pop 方法从栈中弹出一个元素，并更新最小值栈
func (this *MinStack) Pop() {
	this.queue = this.queue[:len(this.queue)-1]          // 弹出数据栈顶元素
	this.minQueue = this.minQueue[:len(this.minQueue)-1] // 弹出最小值栈顶元素
}

// Top 方法返回栈顶元素
func (this *MinStack) Top() int {
	return this.queue[len(this.queue)-1] // 返回数据栈顶元素
}

// GetMin 方法返回栈中的最小值
func (this *MinStack) GetMin() int {
	return this.minQueue[len(this.minQueue)-1] // 返回最小值栈顶元素
}

/**
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
```
