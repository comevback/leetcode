```go
package main

import (
	"fmt"
)

func main() {
	pq := NewPriorityQueue()
	// 向优先队列中添加几个元素
	pq.push(Item{Val: 10, Priority: 1})
	pq.push(Item{Val: 20, Priority: 2})
	pq.push(Item{Val: 15, Priority: 3})

	// 从优先队列中移除元素并打印它们的值
	fmt.Println(pq.pop()) // 应该输出最小值 10
	fmt.Println(pq.pop()) // 接下来是 20
	fmt.Println(pq.pop()) // 最后是 15
}

// Item 定义了优先队列中的元素类型。
type Item struct {
	Val      int // 元素的值
	Priority int // 元素的优先级
}

// PriorityQueue 是基于 Item 的切片，实现了堆数据结构。
type PriorityQueue []Item

// NewPriorityQueue 初始化一个新的优先队列。
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

// swim 方法通过上浮调整确保堆的性质，即父节点的值总是小于其子节点的值。
func (pq *PriorityQueue) swim(index int) {
	parent := index / 2
	// 当当前节点的值小于其父节点的值时，进行交换，并继续上浮。
	for parent >= 1 && (*pq)[parent].Val > (*pq)[index].Val {
		(*pq)[parent], (*pq)[index] = (*pq)[index], (*pq)[parent]
		index = parent
		parent = index / 2
	}
}

// sink 方法通过下沉调整确保堆的性质，修正删除根节点后可能破坏的堆结构。
func (pq *PriorityQueue) sink(index int) {
	for {
		leftChild := 2 * index
		rightChild := 2*index + 1
		min := index

		// 选择值较小的子节点进行比较。
		if leftChild < len(*pq) && (*pq)[leftChild].Val < (*pq)[min].Val {
			min = leftChild
		}

		if rightChild < len(*pq) && (*pq)[rightChild].Val < (*pq)[min].Val {
			min = rightChild
		}

		// 如果当前节点已经是最小的，停止下沉。
		if min == index {
			break
		}

		// 否则，与最小的子节点交换。
		(*pq)[index], (*pq)[min] = (*pq)[min], (*pq)[index]
		index = min
	}
}

// push 方法将一个新的值加入到优先队列中，并进行上浮调整。
func (pq *PriorityQueue) push(item Item) {
	// 在切片为空时，首先添加一个哨兵元素。
	if len(*pq) == 0 {
		(*pq) = append((*pq), Item{})
	}
	// 添加新元素到切片尾部。
	(*pq) = append(*pq, item)
	// 对新元素进行上浮调整。
	pq.swim(len(*pq) - 1)
}

// pop 方法移除并返回优先队列中的最小元素。
func (pq *PriorityQueue) pop() int {
	if len(*pq) <= 1 {
		return -1 // 返回 -1 表示队列为空。
	}

	// 将最后一个元素和第一个元素交换。
	(*pq)[1], (*pq)[len(*pq)-1] = (*pq)[len(*pq)-1], (*pq)[1]
	// 保存将要返回的元素值。
	res := (*pq)[len(*pq)-1].Val
	// 移除最后一个元素。
	(*pq) = (*pq)[:len(*pq)-1]
	// 对新的根元素进行下沉调整。
	pq.sink(1)

	return res
}
```
