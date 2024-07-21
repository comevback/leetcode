```go
package main

func main() {
	// 创建一个新的泛型最小值栈
	newStack := New_Stack_LinkedList[int]()
	// 将元素推入栈中
	newStack.push(3)
	newStack.push(2)
	newStack.push(1)
	newStack.push(4)
	newStack.push(5)
}

// ListNode 是一个泛型单向链表节点，包含一个值和指向下一个节点的指针
type ListNode[T int | int64 | float32 | float64 | string] struct {
	Val  T
	Next *ListNode[T]
}

// =============================== LinkedList Implementation ===================================

// Stack_Min_LinkedList 结构体包含两个链表指针：一个指向当前栈顶元素，另一个指向当前最小值
type Stack_Min_LinkedList[T int | int64 | float32 | float64 | string] struct {
	Val     *ListNode[T]
	Min_Val *ListNode[T]
}

// New_Stack_LinkedList 创建并返回一个新的 Stack_Min_LinkedList 实例
func New_Stack_LinkedList[T int | int64 | float32 | float64 | string]() *Stack_Min_LinkedList[T] {
	return &Stack_Min_LinkedList[T]{}
}

// push 方法将一个新元素加入栈中，同时更新最小值链表
func (stack *Stack_Min_LinkedList[T]) push(value T) {
	// 创建一个新节点，值为 value
	var newNode *ListNode[T] = &ListNode[T]{
		Val: value,
	}

	// 如果 Min_Val 链表为空，或 value 小于当前最小值，将 newNode 插入 Min_Val 链表
	if stack.Min_Val == nil {
		stack.Min_Val = newNode
	} else if value < stack.Min_Val.Val {
		newNode.Next = stack.Min_Val
		stack.Min_Val = newNode
	} else {
		// 如果 value 大于等于当前最小值，将当前最小值复制一份，插入到 Min_Val 链表中
		var temp *ListNode[T] = &ListNode[T]{
			Val:  stack.Min_Val.Val,
			Next: stack.Min_Val,
		}
		stack.Min_Val = temp
	}

	// 将 newNode 插入到栈顶
	if stack.Val == nil {
		stack.Val = newNode
	} else {
		newNode.Next = stack.Val
		stack.Val = newNode
	}
}

// =============================== Array Implementation ===================================

// maxStack 结构体包含两个切片：一个用于存储数据，另一个用于存储最大值
type maxStack struct {
	stack    []int
	maxStack []int
}

// NewMaxStack 创建并返回一个新的 maxStack 实例
func NewMaxStack() maxStack {
	return maxStack{
		stack:    []int{},
		maxStack: []int{},
	}
}

// push 方法将一个新元素加入栈中，同时更新最大值栈
func (ms *maxStack) push(val int) {
	ms.stack = append(ms.stack, val)
	// 如果 maxStack 为空或 val 大于等于当前最大值，将 val 加入 maxStack
	if len(ms.maxStack) == 0 || val >= ms.maxStack[len(ms.maxStack)-1] {
		ms.maxStack = append(ms.maxStack, val)
	}
}

// pop 方法移除栈顶元素，并更新最大值栈
func (ms *maxStack) pop() {
	if len(ms.stack) == 0 {
		return // 如果栈为空，直接返回
	}
	// 如果栈顶元素等于 maxStack 的栈顶元素，将 maxStack 的栈顶元素移除
	if ms.stack[len(ms.stack)-1] == ms.maxStack[len(ms.maxStack)-1] {
		ms.maxStack = ms.maxStack[:len(ms.maxStack)-1]
	}
	// 移除栈顶元素
	ms.stack = ms.stack[:len(ms.stack)-1]
}

// max 方法返回当前栈中的最大值
func (ms *maxStack) max() int {
	if len(ms.maxStack) == 0 {
		return -1 // 如果栈为空，返回一个默认值
	}
	return ms.maxStack[len(ms.maxStack)-1]
}
```
