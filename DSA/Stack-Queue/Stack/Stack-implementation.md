# Stack实现 （Array）（Linked List）

## 代码实现
```go
package main

import (
	"errors" // 导入errors包，用于错误处理
	"fmt"    // 导入fmt包，用于打印输出
)

func main() {
	newStack := NewStack_Link[int]() // 创建一个存储int类型的链表实现的栈
	newStack.push(2)                 // 向栈中推入元素2
	newStack.push(3)                 // 向栈中推入元素3
	newStack.push(7)                 // 向栈中推入元素7
	fmt.Println(newStack)            // 打印栈的当前状态
	fmt.Println(newStack.pop())      // 弹出栈顶元素并打印它（7）
	fmt.Println(newStack)            // 打印栈的当前状态
	newStack.peek()                  // 查看栈顶元素并打印它（3）
}

// =================================================  Array Implementation  ====================================================

// 数组实现的栈定义
type Stack_array[T any] struct {
	Val []T // 使用切片来存储栈的元素
}

// 创建一个空的数组实现的栈
func NewStack_array[T any]() *Stack_array[T] {
	return &Stack_array[T]{Val: []T{}} // 初始化一个空切片
}

// 向栈中推入一个元素
func (stack *Stack_array[T]) push(value T) {
	stack.Val = append(stack.Val, value) // 将元素添加到切片的末尾
}

// 从栈中弹出一个元素
func (stack *Stack_array[T]) pop() T {
	if len(stack.Val) == 0 {
		panic("Stack is empty") // 如果栈为空，则抛出异常
	}
	last := stack.Val[len(stack.Val)-1]      // 获取切片的最后一个元素
	stack.Val = stack.Val[:len(stack.Val)-1] // 移除切片的最后一个元素
	return last                              // 返回被弹出的元素
}

// 查看栈顶元素
func (stack *Stack_array[T]) peek() (T, error) {
	if len(stack.Val) == 0 {
		var null T
		return null, errors.New("empty stack") // 如果栈为空，返回错误
	} else {
		return stack.Val[len(stack.Val)-1], nil // 返回栈顶元素
	}
}

// =================================================  LinkedList Implementation  ====================================================

// 链表节点定义
type LinkedNode[T any] struct {
	Val  T              // 节点存储的值
	Next *LinkedNode[T] // 指向下一个节点的指针
}

// 链表实现的栈定义
type Stack_Link[T any] struct {
	Val *LinkedNode[T] // 指向栈顶节点的指针
}

// 创建一个空的链表实现的栈
func NewStack_Link[T any]() *Stack_Link[T] {
	return &Stack_Link[T]{Val: nil} // 初始化栈顶指针为nil
}

// 向栈中推入一个元素
func (stack *Stack_Link[T]) push(value T) {
	var newNode *LinkedNode[T] = &LinkedNode[T]{Val: value} // 创建一个新节点
	newNode.Next = stack.Val                                // 将新节点的Next指向当前的栈顶元素
	stack.Val = newNode                                     // 更新栈顶元素为新节点
}

// 从栈中弹出一个元素
func (stack *Stack_Link[T]) pop() T {
	if stack.Val == nil {
		panic("this Stack is empty") // 如果栈为空，则抛出异常
	} else {
		last := stack.Val.Val      // 获取栈顶节点的值
		stack.Val = stack.Val.Next // 将栈顶节点更新为下一个节点
		return last                // 返回被弹出的元素
	}
}

// 查看栈顶元素
func (stack *Stack_Link[T]) peek() (T, error) {
	if stack.Val == nil {
		var null T
		return null, errors.New("empty stack") // 如果栈为空，返回错误
	} else {
		return stack.Val.Val, nil // 返回栈顶元素的值
	}
}

// ================================================= 具体应用+方法 ====================================================
// 特别注意：对栈进行操作时，需要用一个节点指针来指向栈顶，这样才能保证不丢失栈顶

// 定义链表节点
type ListNode struct {
	Val  string
	Next *ListNode
}

// 定义栈
type Stack struct {
	head *ListNode
}

// 栈的构造函数
func NewStack() *Stack {
	return &Stack{}
}

// 栈的push方法
func (stack *Stack) push(value string) {
	// 创建新节点
	newNode := &ListNode{
		Val:  value,
		Next: stack.head,
	}
	// 新节点作为栈顶
	stack.head = newNode
}

// 栈的find方法，查找返回值为value的节点
func (stack *Stack) find(value string) *ListNode {
	// 创建一个指针指向栈顶
	current := stack.head

	// 遍历栈
	for current != nil {
		// 如果找到了值为value的节点，返回该节点
		if value == current.Val {
			return current
		}
		// 指针后移
		current = current.Next
	}

	// 没有找到返回nil
	return nil
}

// 栈的delete方法，删除节点
func (stack *Stack) delete(node *ListNode) {

	// 如果要删除的节点是栈顶节点，直接将栈顶指针指向下一个节点
	if node == stack.head {
		stack.head = stack.head.Next
		return
	}

	// 创建一个指针指向栈顶
	current := stack.head
	// 遍历栈
	for current.Next != nil {
		// 如果找到了要删除的节点，将该节点的前一个节点的Next指向该节点的Next
		if node == current.Next {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

// 栈的print方法，打印栈
func (stack *Stack) print() {
	current := stack.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
```