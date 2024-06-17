# 81. Search in Rotated Sorted Array II
Solved
Medium
Topics
Companies
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.

 

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
 

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums is guaranteed to be rotated at some pivot.
-104 <= target <= 104
 

Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?

---

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3}
	res := search(nums, 3)
	fmt.Println(res)
}

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	// 处理重复元素：确保搜索区间两端不包含重复元素
	for left < right {

		// 跳过左侧重复的元素
		for left < right && nums[left] == nums[left+1] {
			left += 1
		}

		// 跳过右侧重复的元素
		for left < right && nums[right] == nums[right-1] {
			right -= 1
		}

		// 执行二分搜索
		mid := left + (right-left)/2

		// 如果用left来判断，未翻转和最后一个翻转会返回同样的值
		// if nums[mid] >= nums[left] {
		// 	left = mid + 1
		// } else if nums[mid] < nums[left] {
		// 	right = mid
		// }

		// 判断中间元素与右端元素的关系，确定旋转的位置
		if nums[mid] >= nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		}
	}

	// reverseIdx 是找到的可能的旋转点
	reverseIdx := left
	// 特殊情况：如果整个数组是有序的，则reverseIdx应为0
	if nums[0] < nums[len(nums)-1] {
		reverseIdx = 0
	}

	// 重置left和right用于标准的二分搜索
	left, right = 0, len(nums)

	// 执行旋转调整过后的二分搜索
	for left < right {
		mid := left + (right-left)/2
		// 计算实际的中间索引，考虑到旋转
		realMid := (mid + reverseIdx) % len(nums)
		if nums[realMid] == target {
			return true // 找到目标值
		} else if nums[realMid] < target {
			left = mid + 1
		} else if nums[realMid] > target {
			right = mid
		}
	}

	// 如果未找到目标值，返回 false
	return false
}
```