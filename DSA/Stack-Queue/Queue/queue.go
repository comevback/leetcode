package queue

import (
	"errors" // 导入errors包，用于处理错误
	"fmt"    // 导入fmt包，用于格式化输出
)

func main() {
	queue := NewQueue[int]() // 创建一个新的队列
	queue.Enqueue(1)         // 向队列中添加元素1
	queue.Enqueue(5)         // 向队列中添加元素5
	queue.Enqueue(3)         // 向队列中添加元素3
	queue.Enqueue(6)         // 向队列中添加元素6
	queue.Enqueue(7)         // 向队列中添加元素7
	queue.PrintQueue()       // 打印队列中的所有元素

	queue.Dequeue()    // 从队列中移除一个元素
	queue.Dequeue()    // 再次从队列中移除一个元素
	queue.PrintQueue() // 打印当前队列中的所有元素

	queue.Enqueue(61)  // 向队列中添加元素61
	queue.Enqueue(75)  // 向队列中添加元素75
	queue.PrintQueue() // 打印当前队列中的所有元素

	Find(queue, 3)     // 在队列中查找元素3
	Delete(queue, 61)  // 从队列中删除元素61
	queue.PrintQueue() // 打印当前队列中的所有元素
}

// ===================================================  Queue结构和方法  ==============================================
// ListNode结构定义
type ListNode[T any] struct {
	Val  T            // 节点存储的值
	Next *ListNode[T] // 指向下一个节点的指针
}

// Queue结构定义
type Queue[T any] struct {
	head *ListNode[T] // 指向队列头部的指针
	tail *ListNode[T] // 指向队列尾部的指针
}

// Queue构造函数
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{} // 初始化一个空队列并返回其地址
}

// Enqueue方法：向队列尾中添加元素
func (queue *Queue[T]) Enqueue(value T) {
	newNode := &ListNode[T]{
		Val: value, // 创建一个新节点，值为传入的value
	}

	if queue.head == nil { // 如果队列为空
		queue.tail = newNode // 新节点既是头也是尾
		queue.head = newNode
		return
	}

	queue.tail.Next = newNode // 将新节点添加到队尾
	queue.tail = newNode      // 更新队尾指针
}

// Dequeue方法：从队列头中移除元素
func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.head == nil { // 如果队列为空
		var null T
		return null, errors.New("empty queue") // 返回空队列错误
	}

	res := queue.head      // 获取头节点
	if res == queue.tail { // 如果队列中只有一个元素
		queue.tail = nil // 将尾指针设为空
	}
	queue.head = queue.head.Next // 更新头指针
	return res.Val, nil          // 返回被删除的元素
}

// peek方法：查看队列头部元素
func (queue *Queue[T]) Peek() (T, error) {
	if queue.head == nil { // 如果队列为空
		var null T
		return null, errors.New("empty queue") // 返回空队列错误
	}

	return queue.head.Val, nil // 返回头节点
}

// isEmpty
func (queue *Queue[T]) IsEmpty() bool {
	return queue.head == nil
}

// find方法：在队列中查找元素
func Find(queue *Queue[int], value int) (*ListNode[int], error) {
	if queue.head == nil { // 如果队列为空
		return nil, errors.New("empty queue") // 返回空队列错误
	}

	current := queue.head

	for current != nil {
		if current.Val == value { // 如果找到值匹配的节点
			return current, nil // 返回该节点
		}
		current = current.Next // 继续检查下一个节点
	}

	return nil, errors.New("didn't find") // 如果没找到，返回未找到错误
}

// delete方法：在队列中删除元素
func Delete(queue *Queue[int], value int) error {
	if queue.head == nil { // 如果队列为空
		return errors.New("empty queue") // 返回空队列错误
	}

	current := queue.head
	var preCurrent *ListNode[int] // 前一个节点初始化为nil

	for current != nil {
		if current.Val == value { // 找到需要删除的节点
			// 删除点是头节点
			if preCurrent == nil { // 如果是头节点
				queue.head = current.Next // 更新头节点
				if queue.head == nil {    // 如果队列变空
					queue.tail = nil // 尾节点也设置为nil
				}
				return nil
			}
			// 删除点是尾节点
			if queue.tail == current { // 如果是尾节点
				preCurrent.Next = nil
				queue.tail = preCurrent // 更新尾节点
				return nil
			}
			// 删除点既不是头节点也不是尾节点
			preCurrent.Next = current.Next // 将前节点指向当前节点的下一个节点
			return nil
		}
		preCurrent = current   // 更新前节点为当前节点
		current = current.Next // 移动到下一个节点
	}

	return errors.New("didn't find") // 如果没有找到，返回未找到错误
}

// printQueue方法：打印队列中的所有元素
func (queue *Queue[T]) PrintQueue() error {
	var arr []T

	current := queue.head
	if current == nil { // 如果队列为空
		return errors.New("empty queue") // 返回空队列错误
	}

	for current != nil {
		arr = append(arr, current.Val) // 将节点的值添加到数组
		current = current.Next         // 移动到下一个节点
	}

	fmt.Println(arr) // 打印所有值组成的数组
	return nil
}
