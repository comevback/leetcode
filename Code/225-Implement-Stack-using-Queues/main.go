package main

import "fmt"

// 主函数，演示MyStack的用法
func main() {
	myStack := Constructor()   // 创建一个新的MyStack实例
	myStack.Push(1)            // 压入元素1
	myStack.Push(2)            // 压入元素2
	fmt.Println(myStack.Top()) // 输出栈顶元素，应该是2
	myStack.Push(3)            // 压入元素3
	fmt.Println(myStack.Top()) // 输出栈顶元素，应该是3
}

// MyStack 结构体定义，包含一个队列和一个记录栈顶元素的变量
type MyStack struct {
	queue   []int // 用于存储元素的队列
	topItem int   // 用于记录栈顶元素的变量
}

// Constructor 构造函数，初始化并返回一个新的MyStack实例
func Constructor() MyStack {
	return MyStack{
		queue: []int{}, // 初始化队列为空
	}
}

// Push 方法，用于将元素x压入栈中
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x) // 将元素x添加到队列末尾
	this.topItem = x                   // 更新栈顶元素为x
}

// Pop 方法，从栈中移除并返回栈顶元素
func (this *MyStack) Pop() int {
	length := len(this.queue)
	// 将队列中的前 length-2 个元素移到队列末尾
	for i := 0; i < length-2; i++ {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}

	top_item := this.queue[0]                 // 获取当前栈顶元素
	this.topItem = top_item                   // 更新栈顶元素
	this.queue = append(this.queue, top_item) // 将栈顶元素重新添加到队列末尾
	this.queue = this.queue[1:]

	poped := this.queue[0]      // 获取新的栈顶元素
	this.queue = this.queue[1:] // 将栈顶元素从队列中移除
	return poped                // 返回被弹出的元素
}

// Top 方法，返回栈顶元素
func (this *MyStack) Top() int {
	return this.topItem // 返回记录的栈顶元素
}

// Empty 方法，检查栈是否为空
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0 // 如果队列长度为0，则栈为空
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
