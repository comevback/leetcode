package main

func main() {
	newStack := New_Stack_LinkedList[int]()
	newStack.push(3)
	newStack.push(2)
	newStack.push(1)
	newStack.push(4)
	newStack.push(5)
}

type ListNode[T int | int64 | float32 | float64 | string] struct {
	Val  T
	Next *ListNode[T]
}

// =============================== LinkedList Implementation ===================================
type Stack_Min_LinkedList[T int | int64 | float32 | float64 | string] struct {
	Val     *ListNode[T]
	Min_Val *ListNode[T]
}

func New_Stack_LinkedList[T int | int64 | float32 | float64 | string]() *Stack_Min_LinkedList[T] {
	return &Stack_Min_LinkedList[T]{}
}

// push方法
func (stack *Stack_Min_LinkedList[T]) push(value T) {
	// 创建一个新节点，值为value
	var newNode *ListNode[T] = &ListNode[T]{
		Val: value,
	}

	// 如果value小于stack.Min_Val.Val，则将newNode插入到stack.Min_Val之前，然后更新stack.Min_Val为newNode
	if stack.Min_Val == nil {
		stack.Min_Val = newNode
	} else if value < stack.Min_Val.Val {
		newNode.Next = stack.Min_Val
		stack.Min_Val = newNode
	} else {
		// 如果value大于等于stack.Min_Val.Val，说明value不是最小值
		// 创造一个新节点，Val为目前的stack.Min_Val.Val，然后插入到stack.Min_Val之前，更新stack.Min_Val为新节点
		var temp *ListNode[T] = &ListNode[T]{
			Val:  stack.Min_Val.Val,
			Next: stack.Min_Val,
		}
		stack.Min_Val = temp
	}

	// 将newNode插入到stack.Val之前
	// 更新newNode为stack.Val
	if stack.Val == nil {
		stack.Val = newNode
	} else {
		newNode.Next = stack.Val
		stack.Val = newNode
	}
}

// =============================== Array Implementation ===================================

type Stack_Min_Array[T any] struct {
	Val     []T
	Min_Val []T
}

// =============================== Array Functions ===================================
