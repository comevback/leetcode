# 1670. Design Front Middle Back Queue

Solved
Medium
Topics
Companies
Hint
Design a queue that supports push and pop operations in the front, middle, and back.

Implement the FrontMiddleBack class:

FrontMiddleBack() Initializes the queue.
void pushFront(int val) Adds val to the front of the queue.
void pushMiddle(int val) Adds val to the middle of the queue.
void pushBack(int val) Adds val to the back of the queue.
int popFront() Removes the front element of the queue and returns it. If the queue is empty, return -1.
int popMiddle() Removes the middle element of the queue and returns it. If the queue is empty, return -1.
int popBack() Removes the back element of the queue and returns it. If the queue is empty, return -1.
Notice that when there are two middle position choices, the operation is performed on the frontmost middle position choice. For example:

Pushing 6 into the middle of [1, 2, 3, 4, 5] results in [1, 2, 6, 3, 4, 5].
Popping the middle from [1, 2, 3, 4, 5, 6] returns 3 and results in [1, 2, 4, 5, 6].

Example 1:

Input:
["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
[[], [1], [2], [3], [4], [], [], [], [], []]
Output:
[null, null, null, null, null, 1, 3, 4, 2, -1]

Explanation:
FrontMiddleBackQueue q = new FrontMiddleBackQueue();
q.pushFront(1); // [1]
q.pushBack(2); // [1, 2]
q.pushMiddle(3); // [1, 3, 2]
q.pushMiddle(4); // [1, 4, 3, 2]
q.popFront(); // return 1 -> [4, 3, 2]
q.popMiddle(); // return 3 -> [4, 2]
q.popMiddle(); // return 4 -> [2]
q.popBack(); // return 2 -> []
q.popFront(); // return -1 -> [] (The queue is empty)

Constraints:

1 <= val <= 109
At most 1000 calls will be made to pushFront, pushMiddle, pushBack, popFront, popMiddle, and popBack.

---

# Code

```go
package main

func main() {

}

// FrontMiddleBackQueue 结构定义了一个自定义队列，可以对队列的前部、中部和后部进行操作。
type FrontMiddleBackQueue struct {
	queue []int // 用切片存储队列元素
}

// Constructor 初始化并返回一个新的 FrontMiddleBackQueue 实例。
func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		queue: []int{},
	}
}

// PushFront 在队列的前部插入一个元素。
func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...) // 将新元素添加到队列的前部
}

// PushMiddle 在队列的中部插入一个元素。
func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	mid := len(this.queue) / 2                                                        // 计算中间位置的索引
	this.queue = append(this.queue[:mid], append([]int{val}, this.queue[mid:]...)...) // 在中间位置插入新元素
}

// PushBack 在队列的后部插入一个元素。
func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.queue = append(this.queue, val) // 将新元素添加到队列的后部
}

// PopFront 从队列的前部移除一个元素并返回该元素。
func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.queue) == 0 {
		return -1 // 如果队列为空，则返回 -1
	}
	poped := this.queue[0]      // 获取前部元素
	this.queue = this.queue[1:] // 移除前部元素
	return poped
}

// PopMiddle 从队列的中部移除一个元素并返回该元素。
func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.queue) == 0 {
		return -1 // 如果队列为空，则返回 -1
	}
	mid := (len(this.queue) - 1) / 2                             // 计算中间位置的索引
	poped := this.queue[mid]                                     // 获取中间元素
	this.queue = append(this.queue[:mid], this.queue[mid+1:]...) // 移除中间元素
	return poped
}

// PopBack 从队列的后部移除一个元素并返回该元素。
func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.queue) == 0 {
		return -1 // 如果队列为空，则返回 -1
	}
	poped := this.queue[len(this.queue)-1]      // 获取后部元素
	this.queue = this.queue[:len(this.queue)-1] // 移除后部元素
	return poped
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
```
