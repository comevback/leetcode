# 704. Binary Search
Solved
Easy
Topics
Companies
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
 

Constraints:

1 <= nums.length <= 104
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 5, 7, 8}
	res := search2(nums, 3)
	fmt.Println(res)
}

// search 寻找特定值
// 采用左闭右闭区间进行查找，返回目标值的任一索引
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // 找到目标，返回索引
		} else if nums[mid] > target {
			right = mid - 1 // 目标值在左侧，移动右边界
		} else if nums[mid] < target {
			left = mid + 1 // 目标值在右侧，移动左边界
		}
	}

	return -1 // 未找到目标值
}

// search1 搜索左边界
// 采用左闭右开区间进行查找，返回目标值的最左索引
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid // 找到目标，移动右边界以收缩范围
		} else if nums[mid] > target {
			right = mid // 目标值在左侧，移动右边界
		} else if nums[mid] < target {
			left = mid + 1 // 目标值在右侧，移动左边界
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1 // 检查越界或值不匹配的情况
	}

	return left // 返回最左边的索引
}

// search2 搜索右边界
// 采用左闭右闭区间进行查找，返回目标值的最右索引
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1 // 找到目标，移动左边界以收缩范围
		} else if nums[mid] > target {
			right = mid - 1 // 目标值在左侧，移动右边界
		} else if nums[mid] < target {
			left = mid + 1 // 目标值在右侧，移动左边界
		}
	}

	if right < 0 || nums[right] != target {
		return -1 // 检查越界或值不匹配的情况
	}

	return right // 返回最右边的索引
}
```