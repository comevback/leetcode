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
