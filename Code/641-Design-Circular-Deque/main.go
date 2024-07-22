package main

func main() {

}

// MyCircularDeque 结构体定义了一个使用切片的循环双端队列。
type MyCircularDeque struct {
	queue []int // 切片，用于存储双端队列的元素。
	head  int   // 双端队列前端元素的索引。
	tail  int   // 双端队列尾端元素的索引。
	size  int   // 双端队列当前的元素数量。
	cap   int   // 双端队列可以容纳的最大元素数量。
}

// Constructor 初始化一个具有给定容量的 MyCircularDeque 实例。
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k), // 以 k 的大小初始化切片。
		head:  -1,             // 初始时没有元素，head 设置为 -1。
		tail:  -1,             // 初始时没有元素，tail 设置为 -1。
		size:  0,              // 初始大小为 0。
		cap:   k,              // 设置容量为 k。
	}
}

// InsertFront 在双端队列前端插入一个元素。
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false // 如果队列已满，返回 false。
	}
	if this.IsEmpty() {
		this.head = 0
		this.tail = 0
	} else {
		this.head = (this.head - 1 + this.cap) % this.cap // 循环移动头指针。
	}
	this.queue[this.head] = value // 在头部位置插入新元素。
	this.size += 1                // 增加元素数量。
	return true
}

// InsertLast 在双端队列尾端插入一个元素。
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false // 如果队列已满，返回 false。
	}
	if this.IsEmpty() {
		this.head = 0
		this.tail = 0
	} else {
		this.tail = (this.tail + 1) % this.cap // 循环移动尾指针。
	}
	this.queue[this.tail] = value // 在尾部位置插入新元素。
	this.size += 1                // 增加元素数量。
	return true
}

// DeleteFront 从双端队列前端删除一个元素。
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false // 如果队列为空，返回 false。
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	} else {
		this.head = (this.head + 1) % this.cap // 循环移动头指针。
	}
	this.size -= 1 // 减少元素数量。
	return true
}

// DeleteLast 从双端队列尾端删除一个元素。
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false // 如果队列为空，返回 false。
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	} else {
		this.tail = (this.tail - 1 + this.cap) % this.cap // 循环移动尾指针。
	}
	this.size -= 1 // 减少元素数量。
	return true
}

// GetFront 获取双端队列前端的元素。
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1 // 如果队列为空，返回 -1。
	}
	return this.queue[this.head] // 返回头部元素。
}

// GetRear 获取双端队列尾端的元素。
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1 // 如果队列为空，返回 -1。
	}
	return this.queue[this.tail] // 返回尾部元素。
}

// IsEmpty 检查双端队列是否为空。
func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0 // 如果大小为 0，则为空。
}

// IsFull 检查双端队列是否已满。
func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.cap // 如果当前大小等于容量，则为满。
}
