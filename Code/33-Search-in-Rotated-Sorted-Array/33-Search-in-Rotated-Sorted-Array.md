# 33. Search in Rotated Sorted Array
Solved
Medium
Topics
Companies
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1
 

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104

# Code
```go
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
```