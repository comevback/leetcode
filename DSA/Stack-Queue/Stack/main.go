package stack

import (
	"errors" // 导入errors包，用于错误处理
	"fmt"    // 导入fmt包，用于打印输出
)

func main() {
	newStack := NewStack_Link[int]() // 创建一个存储int类型的链表实现的栈
	newStack.Push(2)                 // 向栈中推入元素2
	newStack.Push(3)                 // 向栈中推入元素3
	newStack.Push(7)                 // 向栈中推入元素7
	fmt.Println(newStack)            // 打印栈的当前状态
	fmt.Println(newStack.Pop())      // 弹出栈顶元素并打印它（7）
	fmt.Println(newStack)            // 打印栈的当前状态
	newStack.Peek()                  // 查看栈顶元素并打印它（3）
}

// =================================================  LinkedList Implementation  ====================================================

// 链表节点定义
type LinkedNode[T any] struct {
	Val  T              // 节点存储的值
	Next *LinkedNode[T] // 指向下一个节点的指针
}

// 链表实现的栈定义
type Stack_Link[T any] struct {
	Top *LinkedNode[T] // 指向栈顶节点的指针
}

// 创建一个空的链表实现的栈
func NewStack_Link[T any]() *Stack_Link[T] {
	return &Stack_Link[T]{} // 初始化栈顶指针为nil
}

// 向栈中推入一个元素
func (stack *Stack_Link[T]) Push(value T) {
	var newNode *LinkedNode[T] = &LinkedNode[T]{Val: value} // 创建一个新节点

	if stack.Top == nil { // 如果当前stack没有节点，新节点就是stack.Top
		stack.Top = newNode
		return
	}

	newNode.Next = stack.Top // 将新节点的Next指向当前的栈顶元素
	stack.Top = newNode      // 更新栈顶元素为新节点
}

// 从栈中弹出一个元素
func (stack *Stack_Link[T]) Pop() (T, error) {
	if stack.Top == nil {
		var null T
		return null, errors.New("empty stack") // 如果栈为空，则抛出异常
	} else {
		last := stack.Top.Val      // 获取栈顶节点的值
		stack.Top = stack.Top.Next // 将栈顶节点更新为下一个节点
		return last, nil           // 返回被弹出的元素
	}
}

// 查看栈顶元素
func (stack *Stack_Link[T]) Peek() (T, error) {
	if stack.Top == nil {
		var null T
		return null, errors.New("empty stack") // 如果栈为空，返回错误
	} else {
		return stack.Top.Val, nil // 返回栈顶元素的值
	}
}

// 是否为空方法
func (stack *Stack_Link[T]) IsEmpty() bool {
	return stack.Top == nil
}

// Find 方法在栈中查找一个指定的值
func Find(stack *Stack_Link[int], value int) (*LinkedNode[int], error) {
	// 如果栈为空，则返回错误
	if stack.Top == nil {
		return nil, errors.New("empty stack")
	}

	// 从栈顶开始遍历
	current := stack.Top

	// 循环遍历栈中的每一个节点
	for current != nil {
		// 如果找到了一个节点，其值与要查找的值相等
		if current.Val == value {
			// 返回这个节点和一个nil错误（表示没有错误）
			return current, nil
		}
		// 移动到下一个节点
		current = current.Next
	}

	// 如果遍历完整个栈都没有找到指定的值，返回错误
	return nil, errors.New("didn't find")
}

// Delete 方法从栈中删除一个指定的值
func Delete(stack *Stack_Link[int], value int) error {
	// 如果栈为空，则返回错误
	if stack.Top == nil {
		return errors.New("empty stack")
	}

	// 从栈顶开始遍历
	current := stack.Top
	// parentNode 用于记录当前节点的前一个节点
	var parentNode *LinkedNode[int]

	// 循环遍历栈中的每一个节点
	for current != nil {
		// 如果找到了一个节点，其值与要删除的值相等
		if current.Val == value {
			// 如果要删除的节点是栈顶节点
			if parentNode == nil {
				// 将栈顶节点更新为下一个节点
				stack.Top = stack.Top.Next
			} else {
				// 否则，将前一个节点的 Next 指向当前节点的 Next，从链表中断开当前节点
				parentNode.Next = current.Next
			}
			// 删除成功，返回nil表示没有错误
			return nil
		}

		// 更新parentNode为当前节点，current移动到下一个节点
		parentNode = current
		current = current.Next
	}

	// 如果遍历完整个栈都没有找到指定的值，返回错误
	return errors.New("didn't find")
}

// 打印栈中所有元素
func (stack *Stack_Link[T]) PrintStack() {
	var res []T // 用于存储栈中所有元素的切片

	current := stack.Top // 从栈顶开始遍历

	for current != nil { // 循环直到没有更多元素
		res = append(res, current.Val) // 将当前节点的值追加到结果切片中
		current = current.Next         // 移动到下一个节点
	}

	fmt.Println(res) // 打印包含所有栈元素的切片
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
func (stack *Stack_array[T]) Push(value T) {
	stack.Val = append(stack.Val, value) // 将元素添加到切片的末尾
}

// 从栈中弹出一个元素
func (stack *Stack_array[T]) Pop() (T, error) {
	if len(stack.Val) == 0 {
		var null T
		return null, errors.New("empty stack")
	}
	last := stack.Val[len(stack.Val)-1]      // 获取切片的最后一个元素
	stack.Val = stack.Val[:len(stack.Val)-1] // 移除切片的最后一个元素
	return last, nil                         // 返回被弹出的元素
}

// 查看栈顶元素
func (stack *Stack_array[T]) Peek() (T, error) {
	if len(stack.Val) == 0 {
		var null T
		return null, errors.New("empty stack") // 如果栈为空，返回错误
	} else {
		return stack.Val[len(stack.Val)-1], nil // 返回栈顶元素
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
func (stack *Stack) Push(value string) {
	// 创建新节点
	newNode := &ListNode{
		Val:  value,
		Next: stack.head,
	}
	// 新节点作为栈顶
	stack.head = newNode
}

// 栈的find方法，查找返回值为value的节点
func (stack *Stack) Find(value string) *ListNode {
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
func (stack *Stack) Delete(node *ListNode) {

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
func (stack *Stack) Print() {
	current := stack.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
