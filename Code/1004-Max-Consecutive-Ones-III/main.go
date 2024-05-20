package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{0, 0, 0, 1}
	res := longestOnes(nums, 0)
	fmt.Println(res)
}

// ====================================================================================================================
// 1. 我的解法 数组
// 当k > 0时，可以继续把0翻转为1，继续记录长度，但当k=0时，不能再翻转0了，把start转到当前窗口的第一个0的下一个位置
func longestOnes(nums []int, k int) int {
	// start 记录当前窗口的起始位置
	start := 0
	// zeros 记录当前窗口内的0的位置index数组
	zeros := []int{}
	// max 记录最大长度
	max := 0

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 当遇到0时
		if nums[i] == 0 {
			if k > 0 { // 如果k > 0，可以继续翻转0，把新的0的位置加入zeros数组
				k -= 1
				zeros = append(zeros, i)
			} else { // 如果k=0，不能再翻转0了，把start转到当前窗口的第一个0的下一个位置，同时把zeros数组的第一个0的位置移除
				zeros = append(zeros, i)
				start = zeros[0] + 1
				zeros = zeros[1:]
			}
		}

		// 计算当前窗口的长度，如果大于max，更新max
		length := i - start + 1
		if length > max {
			max = length
		}
	}

	// 返回最大长度
	return max
}

// =====================================================================================================================
// 2. 我的解法 队列
// 与上一个解法类似，只是用队列来存储0的位置，每次遇到0时，把0的位置加入队列，当k=0时，把队列的第一个0的位置移除
func longestOnes1(nums []int, k int) int {
	start := 0
	zeros := NewQueue()
	max := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if k > 0 {
				k -= 1
				zeros.EnQueue(i)
			} else {
				zeros.EnQueue(i)
				newIndex, _ := zeros.DeQueue()
				start = newIndex.Val + 1
			}
		}

		length := i - start + 1
		if length > max {
			max = length
		}
	}

	return max
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue struct {
	head *ListNode
	tail *ListNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) EnQueue(value int) {
	newNode := &ListNode{
		Val: value,
	}

	if queue.head == nil {
		queue.head = newNode
		queue.tail = newNode
		return
	}

	queue.tail.Next = newNode
	queue.tail = newNode
}

func (queue *Queue) DeQueue() (*ListNode, error) {
	if queue.tail == nil {
		return nil, errors.New("empty queue")
	}

	res := queue.head
	queue.head = queue.head.Next
	if queue.head == nil {
		queue.tail = nil
	}
	return res, nil
}
