# 162. Find Peak Element
Solved
Medium
Topics
Companies
A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.

 

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.
 

Constraints:

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i.

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{1}
	res := findPeakElement(nums)
	fmt.Println(res)
}

// 1. 差分数组
func findPeakElement(nums []int) int {
	// 创建一个差分数组
	diff := make([]int, len(nums))
	diff[0] = 1 // 将第一个元素视为增加的
	res := []int{}

	// 构建差分数组，记录每个元素与前一个元素的差值
	for i := 1; i < len(nums); i++ {
		difference := nums[i] - nums[i-1]
		diff[i] = difference
	}

	// 将差分数组的末尾视为减少的，以方便处理边界情况
	diff = append(diff, -1)

	// 找出峰值位置，即差分数组中从正变负的点
	for i := 0; i < len(diff)-1; i++ {
		if diff[i] > 0 && diff[i+1] < 0 {
			res = append(res, i)
		}
	}

	// 如果需要返回最大的峰值，遍历找到最大的峰值索引
	resInx := 0
	for _, v := range res {
		if nums[v] > nums[resInx] {
			resInx = v
		}
	}

	return resInx
}

// 2. 二分查找
func findPeakElement1(nums []int) int {
	left, right := 0, len(nums)

	// 使用二分查找来确定峰值的位置
	for left < right {
		mid := left + (right-left)/2

		// 检查当前中点是否为峰值，即比相邻元素都大
		if mid+1 < len(nums) && mid-1 >= 0 && nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		} else if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			// 如果当前元素大于右侧元素，峰值在左侧或当前
			right = mid
		} else if mid-1 >= 0 && nums[mid] > nums[mid-1] {
			// 如果当前元素大于左侧元素，峰值在右侧
			left = mid + 1
		} else {
			// 否则缩小搜索范围
			right = mid
		}
	}

	// 防止索引越界
	if left == len(nums) {
		left = left - 1
	}

	return left
}
```