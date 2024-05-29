package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}

	//var list1e *ListNode
	// var list2e *ListNode

	res := mergeKLists3([]*ListNode{list1, list2, list3})
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.我的方法，不断遍历直到每个链表都到达底端，多次循环太慢
func mergeKLists1(lists []*ListNode) *ListNode {
	// 边界条件直接返回
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	// 设虚拟头节点dummy
	var dummy *ListNode = &ListNode{}
	current := dummy
	count := len(lists) // 记录还有多少个链表没有遍历到底端

	// 遍历链表，当所有链表都到达底端时结束，或全部为空时结束
	for count != 0 {
		var tempMin int = math.MaxInt
		var minIndex int = -1
		for i, list := range lists {
			if list == nil {
				continue
			}
			if list.Val < tempMin {
				minIndex = i
				tempMin = list.Val
			}
		}

		// 如果minIndex为-1，说明所有链表都到达底端或给的链表都为空，直接结束
		if minIndex == -1 {
			break
		}

		// 利用lists[minIndex]的值创建新节点，加入到current.Next，然后lists[minIndex]指向下一个节点
		// newNode := &ListNode{
		// 	Val: tempMin,
		// }
		// current.Next = newNode
		// current = current.Next
		// lists[minIndex] = lists[minIndex].Next

		// 优化：直接将lists[minIndex]的节点加入到current.Next，然后断开lists[minIndex]的头节点，将lists[minIndex]指向下一个节点
		current.Next = lists[minIndex]
		current = current.Next

		temp := lists[minIndex].Next
		lists[minIndex].Next = nil
		lists[minIndex] = temp

		// 如果lists[minIndex]为空，count减1
		if lists[minIndex] == nil {
			count -= 1
		}
	}

	// 返回dummy.Next
	return dummy.Next
}

// =====================================================================================================================
// 2. 我的方法，递归分治
func MergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	current := dummy

	c1, c2 := list1, list2

	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			current.Next = c1
			c1 = c1.Next
		} else {
			current.Next = c2
			c2 = c2.Next
		}
		current = current.Next
	}

	if c1 != nil {
		current.Next = c1
	}

	if c2 != nil {
		current.Next = c2
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return MergeTwoList(lists[0], lists[1])
	}

	return MergeTwoList(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

// ****************************************************  优先队列方法  ****************************************************
// 通过把每个链表的头节点加入到优先队列中，每次取出时，取出堆顶的链表的头节点，如果该链表还有节点，将这个节点sink到合适的位置
// 如果取出头节点的堆顶链表为空，pop掉堆顶，继续循环
// 如果所有链表都为空，结束循环
func mergeKLists3(lists []*ListNode) *ListNode {
	// 边界条件直接返回
	if len(lists) == 0 {
		return nil
	}

	// 定义一个优先队列
	pq := NewPQ()
	// 遍历lists，将每个链表的头节点加入到优先队列中
	for _, v := range lists {
		if v == nil { // 如果链表为空，直接跳过
			continue
		}
		pq.push(v)
	}

	// 定义一个虚拟头节点
	var dummy *ListNode = &ListNode{}
	current := dummy // 定义一个current指针指向dummy

	// 循环，直到优先队列里所有链表都为空
	for {
		node, err := pq.getMinNode()
		if err != nil {
			break
		}
		current.Next = node
		current = current.Next
	}

	// 返回dummy.Next
	return dummy.Next
}

// ****************************************************  优先队列实现  ****************************************************
// 定义优先队列
type PQ []*ListNode

// 优先队列构造函数
func NewPQ() *PQ {
	return &PQ{}
}

// 优先队列的上浮操作
func (pq *PQ) swim(index int) {
	parent := index / 2
	for parent >= 1 && (*pq)[parent].Val > (*pq)[index].Val {
		(*pq)[parent], (*pq)[index] = (*pq)[index], (*pq)[parent]
		index = parent
		parent = index / 2
	}
}

// 优先队列的下沉操作
// 找到左右子节点中最小的，然后和当前节点比较，如果当前节点比最小的大，交换，然后继续下沉
// 当左右子节点都比当前节点大，或者当前节点已经没有子节点时，停止下沉
func (pq *PQ) sink(index int) {
	for {
		leftChild := 2 * index
		rightChild := 2*index + 1
		min := index

		if leftChild < len(*pq) && (*pq)[leftChild].Val < (*pq)[min].Val {
			min = leftChild
		}

		if rightChild < len(*pq) && (*pq)[rightChild].Val < (*pq)[min].Val {
			min = rightChild
		}

		if index == min {
			break
		}

		(*pq)[index], (*pq)[min] = (*pq)[min], (*pq)[index]

		index = min
	}
}

// 优先队列的push操作
func (pq *PQ) push(node *ListNode) {
	if node == nil {
		return
	}

	if len(*pq) == 0 {
		(*pq) = append((*pq), nil)

	}

	(*pq) = append((*pq), node)
	pq.swim(len(*pq) - 1)
}

// 优先队列的pop操作
func (pq *PQ) pop() error {
	// 边界条件
	if len(*pq) == 1 {
		return errors.New("empty queue")
	}
	if len(*pq) == 2 {
		(*pq) = (*pq)[:len(*pq)-1]
		return nil
	}

	// 交换堆顶和堆底的元素，然后切掉堆底的元素
	// 然后对堆顶元素下沉
	(*pq)[1] = (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	pq.sink(1)

	return nil
}

// 优先队列的取出堆顶元素
func (pq *PQ) getMinNode() (*ListNode, error) {
	// 边界条件
	if len(*pq) <= 1 {
		return nil, errors.New("empty queue")
	}

	// 创建返回的新节点
	newNode := &ListNode{
		Val: (*pq)[1].Val,
	}

	// 将堆顶元素指向下一个节点
	(*pq)[1] = (*pq)[1].Next
	// 如果堆顶元素为空，pop掉堆顶
	if (*pq)[1] == nil {
		err := pq.pop()
		if err != nil {
			return nil, err
		}
	} else { // 如果堆顶元素不为空，把堆顶元素下沉
		pq.sink(1)
	}

	return newNode, nil
}

// 也可以用一般的优先队列实现，pop时直接pop出整个链表，然后把头节点给current.Next，然后把头节点的下一个节点push进优先队列
