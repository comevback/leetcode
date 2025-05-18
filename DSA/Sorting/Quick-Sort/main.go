package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	QuickSort(arr)
	fmt.Println(arr)
}

// *****************************************************  递归实现  *****************************************************
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	key := len(arr) - 1 // 这里是选择最后一个元素作为轴，如果用的是其他的元素，可以把这个元素和最后一个元素调换，然后开始遍历
	head := 0
	for head < key {
		if arr[head] > arr[key] {
			// 如果要交换的两个元素相邻，直接交换
			if head+1 == key {
				arr[key], arr[head] = arr[head], arr[key]
			} else {
				// 如果不相邻，head，key-1，key三个数组互相交换
				arr[key], arr[head], arr[key-1] = arr[head], arr[key-1], arr[key]
			}
			key -= 1
		} else {
			head += 1
		}
	}

	// 如果key左边有元素，对左边元素进行递归，右侧同理
	if key > 0 {
		QuickSort(arr[:key])
	}
	if key < len(arr)-1 {
		QuickSort(arr[key+1:])
	}
}

func QuickSort1(nums []int) {
	if len(nums) < 2 {
		return
	}

	pivot := rand.Intn(len(nums))
	nums[len(nums)-1], nums[pivot] = nums[pivot], nums[len(nums)-1]
	pivot = len(nums) - 1
	head := 0

	for head < pivot {
		if nums[pivot] < nums[head] {
			if pivot == head+1 {
				nums[pivot], nums[head] = nums[head], nums[pivot]
			} else {
				nums[pivot], nums[head], nums[pivot-1] = nums[head], nums[pivot-1], nums[pivot]
			}
			pivot -= 1
		} else {
			head += 1
		}
	}

	if pivot+1 < len(nums) {
		QuickSort1(nums[pivot+1:])
	}
	QuickSort1(nums[:pivot])
}

// *****************************************************  迭代实现  *****************************************************
