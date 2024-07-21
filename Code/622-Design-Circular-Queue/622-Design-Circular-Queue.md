# 622. Design Circular Queue（循环队列）

Solved
Medium
Topics
Companies
Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle, and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Implement the MyCircularQueue class:

MyCircularQueue(k) Initializes the object with the size of the queue to be k.
int Front() Gets the front item from the queue. If the queue is empty, return -1.
int Rear() Gets the last item from the queue. If the queue is empty, return -1.
boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
boolean isEmpty() Checks whether the circular queue is empty or not.
boolean isFull() Checks whether the circular queue is full or not.
You must solve the problem without using the built-in queue data structure in your programming language.

Example 1:

Input
["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Output
[null, true, true, true, false, 3, true, true, true, 4]

Explanation
MyCircularQueue myCircularQueue = new MyCircularQueue(3);
myCircularQueue.enQueue(1); // return True
myCircularQueue.enQueue(2); // return True
myCircularQueue.enQueue(3); // return True
myCircularQueue.enQueue(4); // return False
myCircularQueue.Rear(); // return 3
myCircularQueue.isFull(); // return True
myCircularQueue.deQueue(); // return True
myCircularQueue.enQueue(4); // return True
myCircularQueue.Rear(); // return 4

Constraints:

1 <= k <= 1000
0 <= value <= 1000
At most 3000 calls will be made to enQueue, deQueue, Front, Rear, isEmpty, and isFull.

---

# Code

```go
package main

func main() {
	obj := Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	obj.Rear()
	obj.IsFull()
	obj.DeQueue()
	obj.EnQueue(4)
	obj.Rear()
}

// MyCircularQueue 是一个循环队列的结构体，使用头尾指针、容量和大小来管理队列
type MyCircularQueue struct {
	queue []int // 存储队列元素的切片
	head  int   // 指向队列头部的指针
	tail  int   // 指向队列尾部的指针
	cap   int   // 队列的容量
	size  int   // 队列中元素的数量
}

// Constructor 初始化队列，设置队列长度为 k
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k), // 分配长度为 k 的切片
		head:  -1,             // 初始化头指针
		tail:  -1,             // 初始化尾指针
		cap:   k,              // 设置队列容量
		size:  0,              // 初始化队列大小
	}
}

// EnQueue 向循环队列插入一个元素。如果成功插入则返回 true
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false // 如果队列已满，插入失败
	}

	if this.IsEmpty() {
		this.head = 0 // 如果队列为空，设置头指针
	}

	this.tail = (this.tail + 1) % this.cap // 移动尾指针
	this.queue[this.tail] = value          // 插入元素
	this.size += 1                         // 更新队列大小
	return true
}

// DeQueue 从循环队列中删除一个元素。如果成功删除则返回 true
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false // 如果队列为空，删除失败
	}

	if this.head == this.tail {
		this.head = -1 // 如果只有一个元素，重置头尾指针
		this.tail = -1
	}

	this.head = (this.head + 1) % this.cap // 移动头指针
	this.size -= 1                         // 更新队列大小
	return true
}

// Front 从队首获取元素。如果队列为空，返回 -1
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1 // 如果队列为空，返回 -1
	}

	return this.queue[this.head] // 返回队首元素
}

// Rear 获取队尾元素。如果队列为空，返回 -1
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1 // 如果队列为空，返回 -1
	}
	return this.queue[this.tail] // 返回队尾元素
}

// IsEmpty 检查循环队列是否为空
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

// IsFull 检查循环队列是否已满
func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}

// ******************************************************  不使用size和cap  *****************************************************
// MyCircularQueueV2 是一个循环队列的结构体，使用头尾指针和容量来管理队列
type MyCircularQueueV2 struct {
	data  []int // 存储队列元素的切片
	head  int   // 指向队列头部的指针
	tail  int   // 指向队列尾部的指针
	limit int   // 队列的容量
}

// ConstructorV2 初始化队列，设置队列长度为 k
func ConstructorV2(k int) MyCircularQueueV2 {
	return MyCircularQueueV2{
		data:  make([]int, k+1), // 多分配一个空间区分队列满和空
		head:  0,                // 初始化头指针
		tail:  0,                // 初始化尾指针
		limit: k + 1,            // 设置队列容量
	}
}

// ** 因为如果(q.tail+1)%q.limit == q.head，也就是说tail下一个是head，那么就是满的，所以不会出现总共有k+1个元素的情况
func (q *MyCircularQueueV2) EnQueueV2(value int) bool {
	if q.IsFullV2() {
		return false // 如果队列已满，插入失败
	}
	q.data[q.tail] = value          // 插入元素
	q.tail = (q.tail + 1) % q.limit // 移动尾指针
	return true
}

// DeQueueV2 从循环队列中删除一个元素。如果成功删除则返回 true
func (q *MyCircularQueueV2) DeQueueV2() bool {
	if q.IsEmptyV2() {
		return false // 如果队列为空，删除失败
	}
	q.head = (q.head + 1) % q.limit // 移动头指针
	return true
}

// FrontV2 从队首获取元素。如果队列为空，返回 -1
func (q *MyCircularQueueV2) FrontV2() int {
	if q.IsEmptyV2() {
		return -1 // 如果队列为空，返回 -1
	}
	return q.data[q.head] // 返回队首元素
}

// RearV2 获取队尾元素。如果队列为空，返回 -1
func (q *MyCircularQueueV2) RearV2() int {
	if q.IsEmptyV2() {
		return -1 // 如果队列为空，返回 -1
	}
	// 因为 tail 指向下一个插入的位置，所以需要取前一个位置的元素
	return q.data[(q.tail-1+q.limit)%q.limit]
}

// IsEmptyV2 检查循环队列是否为空
func (q *MyCircularQueueV2) IsEmptyV2() bool {
	return q.head == q.tail
}

// IsFullV2 检查循环队列是否已满
func (q *MyCircularQueueV2) IsFullV2() bool {
	return (q.tail+1)%q.limit == q.head
}
```
