package main

import "fmt"

func main() {
	nums := []int{8, 9, 2, 3, 4}
	res := search(nums, 9)
	fmt.Println(res)
}

// 1. 寻找旋转点，分别二分查找
func search(nums []int, target int) int {
	// 寻找旋转点
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		}
	}

	// 旋转点为left-1，如果left为0，说明没有旋转
	var reverIdx int = left - 1
	if left == 0 {
		reverIdx = 0
	}

	// 根据旋转点，分别二分查找
	// 旋转点左边
	left, right = 0, reverIdx

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	// 旋转点右边
	left, right = reverIdx, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

// =========================================================================================================
// 2. 二分查找，不分别找旋转点
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	// 寻找旋转点
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// left 是最小元素的索引，也是旋转点
	rot := left
	left, right = 0, len(nums)-1

	// 在旋转的数组中找到目标
	for left <= right {
		mid := left + (right-left)/2
		realMid := (mid + rot) % len(nums) // 考虑旋转的实际中点
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
