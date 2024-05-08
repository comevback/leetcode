package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	insertionSort(arr)
	fmt.Println(arr)
}

// insertionSort - 插入排序 O(n^2)-O(1)
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// key为当前需要排序的元素，先将其存储起来，以免被覆盖
		key := arr[i]
		// j为已排序的元素的最后一个元素的下标
		j := i - 1

		// 将已排序的元素中比key大的元素后移一位，直到找到比key小的元素
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// 将原本的arr[i]插入到这个比arr[i]小的元素后面
		arr[j+1] = key
	}
}

// 我的实现
func insertSorting(nums []int) {
	// 从第二个元素开始遍历，每个i都是未排名的，i之前的元素都是已排名的
	for i := 1; i < len(nums); i++ {
		// 从i开始，向前遍历，如果i小于i-1，则交换位置
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
