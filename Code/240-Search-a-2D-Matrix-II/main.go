package main

import "fmt"

func main() {
	// 初始化一个二维矩阵
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5 // 指定搜索的目标值
	// 使用第二种方法进行搜索
	res := searchMatrix(matrix, target)
	fmt.Println(res) // 输出搜索结果
}

// searchMatrix 使用优先队列将矩阵转换为排序后的一维数组并进行二分查找
func searchMatrix(matrix [][]int, target int) bool {
	arr := make([]int, len(matrix)*len(matrix[0])) // 创建用于存储矩阵元素的一维数组
	myPQ := newPQ()                                // 初始化一个优先队列

	// 将每一行的数据加入到优先队列中
	for i := 0; i < len(matrix); i++ {
		myPQ.Enqueue(matrix[i])
	}

	// 从优先队列中逐个提取元素，构造有序数组
	for i := 0; i < len(arr); i++ {
		add := myPQ.Dequeue() // 从优先队列中获取最小元素的行
		arr[i] = add[0]       // 添加最小元素到数组
		add = add[1:]         // 移除已经处理的最小元素
		if len(add) > 0 {
			myPQ.Enqueue(add) // 如果当前行还有剩余元素，再次加入优先队列
		}
	}

	// 二分查找目标值
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2 // 计算中间位置
		if arr[mid] == target {
			return true // 找到目标值，返回true
		} else if arr[mid] < target {
			left = mid + 1 // 目标值较大，调整左边界
		} else {
			right = mid // 目标值较小，调整右边界
		}
	}

	return false // 未找到目标值
}

// PQ 定义一个优先队列结构
type PQ struct {
	queue [][]int // 使用二维数组实现最小堆，存储行的引用
}

// newPQ 初始化并返回一个新的优先队列实例
func newPQ() *PQ {
	return &PQ{queue: [][]int{}} // 预留第一个位置，方便索引计算
}

// sink 下沉操作，用于维护最小堆性质
func (pq *PQ) sink(index int) {
	for {
		left, right := index*2, index*2+1 // 计算左右子节点的索引
		min := index                      // 假设当前节点为最小

		// 找到三者中的最小值
		if left < len(pq.queue) && pq.queue[left][0] < pq.queue[min][0] {
			min = left
		}
		if right < len(pq.queue) && pq.queue[right][0] < pq.queue[min][0] {
			min = right
		}

		// 如果当前节点是最小的，停止下沉
		if min == index {
			break
		}

		// 否则，交换当前节点和最小节点，继续下沉
		pq.queue[min], pq.queue[index] = pq.queue[index], pq.queue[min]
		index = min
	}
}

// swim 上浮操作，用于插入元素后维护最小堆性质
func (pq *PQ) swim(index int) {
	for {
		parent := index / 2 // 计算父节点的索引
		// 如果当前节点小于父节点，需要交换
		if parent >= 1 && pq.queue[index][0] < pq.queue[parent][0] {
			pq.queue[index], pq.queue[parent] = pq.queue[parent], pq.queue[index]
			index = parent
		} else {
			break
		}
	}
}

// Enqueue 将一个数组加入到优先队列中
func (pq *PQ) Enqueue(nums []int) {
	pq.queue = append(pq.queue, nums) // 将新数组添加到堆的末尾
	pq.swim(len(pq.queue) - 1)        // 执行上浮操作，维护堆的性质
}

// Dequeue 从优先队列中移除并返回具有最小值的元素的数组
func (pq *PQ) Dequeue() []int {
	if len(pq.queue) <= 1 {
		return nil // 如果队列为空，返回nil
	}
	// 将堆顶元素（最小元素所在数组）与最后一个元素交换
	pq.queue[1], pq.queue[len(pq.queue)-1] = pq.queue[len(pq.queue)-1], pq.queue[1]
	temp := pq.queue[len(pq.queue)-1]     // 保存堆顶元素
	pq.queue = pq.queue[:len(pq.queue)-1] // 移除堆顶元素
	pq.sink(1)                            // 执行下沉操作，维护堆的性质
	return temp                           // 返回原堆顶元素
}

// ***************************************************  右上开始搜索  *********************************************
// searchMatrix1 从矩阵的右上角开始搜索目标值
func searchMatrix1(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1 // 初始化位置为矩阵的右上角

	// 搜索目标值
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true // 直接找到目标值
		} else if matrix[i][j] > target {
			j-- // 目标值较小，向左移动
		} else {
			i++ // 目标值较大，向下移动
		}
	}

	return false // 搜索结束，未找到目标值
}
