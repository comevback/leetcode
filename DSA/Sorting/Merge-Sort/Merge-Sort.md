# Merge-Sort

![Merge-Sort.gif](Merge-Sort.gif)

## 递归的 sort 函数就是二叉树的遍历函数，而 merge 函数就是在每个节点上做的事情

# Code
```go
package main

import (
	"fmt"

	queue "github.com/comevback/leetcode/DSA/Stack-Queue/Queue"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	res := MergeSortIterate(arr)
	fmt.Println(res)
}

// =====================================================  递归实现  =======================================================
// 归并算法
func MergeSort(arr []int) []int {
	// 如果给的arr没有元素或只有一个元素，直接返回本身
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	// 如果arr长度大于1，分化
	if len(arr) == 1 {
		return arr
	} else {
		return mergeIndex(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
	}
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
		left := divided.Val[:len(divided.Val)/2]
		right := divided.Val[len(divided.Val)/2:]

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
		if len(top.Val) == length {
			break
		}

		// 弹出头两个元素
		merge1, _ := resQueue.Dequeue()
		merge2, _ := resQueue.Dequeue()

		// 把这两个数组进行merge，加入队尾
		new := mergeIndex(merge1.Val, merge2.Val)
		resQueue.Enqueue(new)
	}

	// 弹出排序好的数组
	res, _ := resQueue.Dequeue()
	return res.Val
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

// *************************************************  review 6.18  ****************************************************
// 指针实现
func MergeSort_618(nums []int, temp []int, left int, right int) {
	// 如果只有一个元素，直接返回
	if left == right-1 {
		return
	}

	mid := left + (right-left)/2
	MergeSort_618(nums, temp, left, mid)
	MergeSort_618(nums, temp, mid, right)

	merge_618(nums, temp, left, mid, right)
}

// 合并数组中的left到mid和mid到right两个部分
func merge_618(nums []int, temp []int, left int, mid int, right int) {
	copy(temp[left:right], nums[left:right])
	l1, l2 := left, mid
	current := left

	for l1 < mid && l2 < right {
		if temp[l1] <= temp[l2] {
			nums[current] = temp[l1]
			l1 += 1
		} else {
			nums[current] = temp[l2]
			l2 += 1
		}
		current += 1
	}

	for l1 < mid {
		nums[current] = temp[l1]
		l1 += 1
		current += 1
	}

	for l2 < right {
		nums[current] = temp[l2]
		l2 += 1
		current += 1
	}
}
```

---

// ========================================  实例：leetcode 912 Sort-an-Array  ============================================
# 912. Sort an Array

Medium

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.


Example 1:
> Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).

Example 2:
> Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessairly unique.
 

Constraints:

1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104

---

# Code

## 方法一：指针法
```go
package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	res := sortArray(arr)
	fmt.Println(res)
}

func sortArray(nums []int) []int {
	temp := make([]int, len(nums)) // 提前分配空间，通过参数传递，避免递归中频繁开辟空间，在面对大数组时，会有性能提升
	MergeSort(nums, 0, len(nums)-1, temp)
	return nums
}

func MergeSort(nums []int, lo int, hi int, temp []int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2

	MergeSort(nums, lo, mid, temp)
	MergeSort(nums, mid+1, hi, temp)

	Merge(nums, lo, mid, hi, temp)
}

// Merge 合并两个有序数组
func Merge(nums []int, lo int, mid int, hi int, temp []int) {
	// 把需要排序的数组部分复制到临时数组部分中
	copy(temp[lo:hi+1], nums[lo:hi+1])

	// 定义左右指针，分别指向左半部分的起始位置，右半部分的起始位置
	var leftIndex, rightIndex int = lo, mid + 1
	var current int = lo // 当前指针，指向当前需要填充的位置

	// 比较左右两个部分的元素，把较小的元素填充到原数组中
	for leftIndex < mid+1 && rightIndex < hi+1 { // 当左右两个部分都不越界时，继续比较
		if temp[leftIndex] <= temp[rightIndex] { // 如果左半部分的元素小于等于右半部分的元素，把左半部分的元素填充到原数组中
			nums[current] = temp[leftIndex]
			leftIndex += 1
		} else {
			nums[current] = temp[rightIndex] // 如果右半部分的元素小于左半部分的元素，把右半部分的元素填充到原数组中
			rightIndex += 1
		}
		current += 1 // 填充完一个元素后，当前指针后移
	}

	// 如果左半部分还有剩余元素，把剩余元素填充到原数组中，如果右半部分还有剩余元素，把剩余元素填充到原数组中
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
```

## 方法二：切片法
```go
func sortArray(nums []int) []int {
	res := MergeSort(nums)
	return res
}
// 归并算法
func MergeSort(arr []int) []int {
	// 如果给的arr没有元素或只有一个元素，直接返回本身
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	// 如果arr长度大于1，分化
	if len(arr) == 1 {
		return arr
	} else {
		return mergeIndex(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
	}
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
```