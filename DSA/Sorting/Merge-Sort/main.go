package main

import (
	"fmt"

	queue "github.com/comevback/leetcode/DSA/Stack-Queue/Queue"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	res := MergeSort(arr)
	fmt.Println(res)
}

// =====================================================  递归实现  =======================================================
// 归并算法
func MergeSort(arr []int) []int {
	// 如果给的arr没有元素或只有一个元素，直接返回本身
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	return mergeIndex(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

// 切片方式的merge
// 归并两个数组的操作，当两个数组长度都不为零时，每次比较它们的头元素，把小的那个取出加入到结果数组，当有一个数组空时，把另一个直接加入结果数组的尾部。
func mergeCut(arr1 []int, arr2 []int) []int {
	res := []int{}
	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] <= arr2[0] {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		}
	}

	if len(arr1) != 0 {
		res = append(res, arr1...)
	}

	if len(arr2) != 0 {
		res = append(res, arr2...)
	}

	return res
}

// 指针方式的merge，内存更优
func mergeIndex(arr1 []int, arr2 []int) []int {
	res := []int{}
	index1, index2 := 0, 0

	for index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1] <= arr2[index2] {
			res = append(res, arr1[index1])
			index1 += 1
		} else {
			res = append(res, arr2[index2])
			index2 += 1
		}
	}

	if index1 < len(arr1) {
		res = append(res, arr1[index1:]...)
	}

	if index2 < len(arr2) {
		res = append(res, arr2[index2:]...)
	}

	return res
}

// =======================================================  迭代实现  =====================================================
// 迭代方式实现 MergeSort
func MergeSortIterate(arr []int) []int {
	// 如果给的arr没有元素或只有一个元素，直接返回本身
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	length := len(arr)
	// 定义一个输入队列，用于把一整个arr分化成长度为1的数组
	inputQueue := queue.NewQueue[[]int]()
	// 首先把arr推进这个队列
	inputQueue.Enqueue(arr)
	// 定义一个输出队列，用于把所有长度为1的数组归并为一个长度为原数组的排序后数组
	resQueue := queue.NewQueue[[]int]()

	// 当输入队列不为空的时候，继续循环
	for !inputQueue.IsEmpty() {
		// 弹出队列头
		divided, _ := inputQueue.Dequeue()
		// 得到队列头代表的元素的左右半部分
		left := divided[:len(divided)/2]
		right := divided[len(divided)/2:]

		// 如果左右半部分长度为1，加入到输出队列，否则继续加入到输入队列
		if len(left) == 1 {
			resQueue.Enqueue(left)
		} else {
			inputQueue.Enqueue(left)
		}
		if len(right) == 1 {
			resQueue.Enqueue(right)
		} else {
			inputQueue.Enqueue(right)
		}
	}

	// 当输出队列不为空，循环
	for !resQueue.IsEmpty() {
		// 如果队列头元素的数组长度为原数组，说明已经排序好，推出循环
		top, _ := resQueue.Peek()
		if len(top) == length {
			break
		}

		// 弹出头两个元素
		merge1, _ := resQueue.Dequeue()
		merge2, _ := resQueue.Dequeue()

		// 把这两个数组进行merge，加入队尾
		new := mergeIndex(merge1, merge2)
		resQueue.Enqueue(new)
	}

	// 弹出排序好的数组
	res, _ := resQueue.Dequeue()
	return res
}

// ****************************************************  指针方法  ****************************************************
func sortArray(nums []int) []int {
	temp := make([]int, len(nums)) // 提前定义这个数组，避免在递归中重复定义，如果是很大的数组，在递归中定义的数组被重复创建和销毁，很浪费时间
	MergeSortIndex(nums, 0, len(nums)-1, temp)
	return nums
}

func MergeSortIndex(nums []int, lo int, hi int, temp []int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2 // 这种方法可能引起整型溢出
	mid = lo + (hi-lo)/2 // 这种方法不会溢出

	MergeSortIndex(nums, lo, mid, temp)
	MergeSortIndex(nums, mid+1, hi, temp)

	Merge(nums, lo, mid, hi, temp)
}

func Merge(nums []int, lo int, mid int, hi int, temp []int) {
	copy(temp[lo:hi+1], nums[lo:hi+1])

	var leftIndex, rightIndex int = lo, mid + 1
	var current int = lo

	for leftIndex < mid+1 && rightIndex < hi+1 {
		if temp[leftIndex] <= temp[rightIndex] {
			nums[current] = temp[leftIndex]
			leftIndex += 1
		} else {
			nums[current] = temp[rightIndex]
			rightIndex += 1
		}

		current += 1
	}

	for leftIndex < mid+1 {
		nums[current] = temp[leftIndex]
		current += 1
		leftIndex += 1
	}

	for rightIndex < hi+1 {
		nums[current] = temp[rightIndex]
		current += 1
		rightIndex += 1
	}

}
