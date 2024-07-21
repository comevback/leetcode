```go
package main

func main() {

}

// ********************************************  LinkedList Implementation  ******************************************
// LinkNode 是一个双向链表节点，包含一个值和指向前后节点的指针
type LinkNode struct {
	Val    int
	Before *LinkNode
	Next   *LinkNode
}

// maxQueue_LK 结构体包含一个普通队列和一个双向链表用于维护最大值
type maxQueue_LK struct {
	queue    []int     // 用于存储队列元素
	ListHead *LinkNode // 双向链表的头节点，维护当前最大值
	ListTail *LinkNode // 双向链表的尾节点，方便添加新元素
}

// NewMaxQueue_LK 创建并返回一个新的 maxQueue_LK 实例
func NewMaxQueue_LK() maxQueue_LK {
	return maxQueue_LK{}
}

// Enqueue 方法将一个新元素加入队列，同时更新双向链表保持递减顺序
func (mq *maxQueue_LK) Enqueue(val int) {
	mq.queue = append(mq.queue, val)
	current := mq.ListTail
	// 通过比较，找到链表中小于新元素的节点
	for current != nil && current.Val < val {
		current = current.Before
	}

	newNode := &LinkNode{Val: val}
	if current == nil {
		// 如果当前链表为空或所有节点值都小于新元素，将新节点作为头和尾节点
		mq.ListHead = newNode
		mq.ListTail = newNode
	} else {
		// 否则，将新节点插入到链表尾部
		current.Next = newNode
		newNode.Before = current
		mq.ListTail = newNode
	}
}

// Dequeue 方法移除队列中的第一个元素，并更新双向链表
func (mq *maxQueue_LK) Dequeue() {
	if len(mq.queue) == 0 {
		return // 如果队列为空，直接返回
	}

	poped := mq.queue[0]
	mq.queue = mq.queue[1:]

	// 如果移除的元素是双向链表的头节点，更新头节点
	if poped == mq.ListHead.Val {
		mq.ListHead = mq.ListHead.Next
		if mq.ListHead != nil {
			mq.ListHead.Before = nil
		}
	}
}

// ViewMax 方法返回当前队列中的最大值
func (mq *maxQueue_LK) ViewMax() int {
	if mq.ListHead == nil {
		return -1 // 如果队列为空，返回一个默认值
	}
	return mq.ListHead.Val
}

// ViewTop 方法返回队列中的第一个元素
func (mq *maxQueue_LK) ViewTop() int {
	if len(mq.queue) == 0 {
		return -1 // 如果队列为空，返回一个默认值
	}
	return mq.queue[0]
}

// ********************************************  Array Implementation  ******************************************
// maxQueue 结构体包含一个普通队列和一个双端队列用于维护最大值
type maxQueue struct {
	queue []int
	maxQ  []int
}

// NewMaxQueue 创建并返回一个新的 maxQueue 实例
func NewMaxQueue() maxQueue {
	return maxQueue{
		queue: []int{},
		maxQ:  []int{},
	}
}

// Enqueue 方法将一个新元素加入队列，同时更新 maxQ 保持递减顺序
func (mq *maxQueue) Enqueue(val int) {
	mq.queue = append(mq.queue, val)
	// 移除 maxQ 中所有小于当前值的元素
	for len(mq.maxQ) > 0 && mq.maxQ[len(mq.maxQ)-1] < val {
		mq.maxQ = mq.maxQ[:len(mq.maxQ)-1]
	}
	// 将新值加入 maxQ
	mq.maxQ = append(mq.maxQ, val)
}

// Dequeue 方法移除队列中的第一个元素，并更新 maxQ
func (mq *maxQueue) Dequeue() {
	if len(mq.queue) == 0 {
		return
	}
	// 移除队列中的第一个元素
	poped := mq.queue[0]
	mq.queue = mq.queue[1:]
	// 如果移除的元素是 maxQ 中的最大值，也移除它
	if poped == mq.maxQ[0] {
		mq.maxQ = mq.maxQ[1:]
	}
}

// ViewMax 方法返回当前队列中的最大值
func (mq *maxQueue) ViewMax() int {
	if len(mq.maxQ) == 0 {
		return -1 // 如果队列为空，返回一个默认值
	}
	return mq.maxQ[0]
}

// ViewTop 方法返回队列中的第一个元素
func (mq *maxQueue) ViewTop() int {
	if len(mq.queue) == 0 {
		return -1 // 如果队列为空，返回一个默认值
	}
	return mq.queue[0]
}
```
