# 852. Peak Index in a Mountain Array
Solved
Medium
Topics
Companies
You are given an integer mountain array arr of length n where the values increase to a peak element and then decrease.

Return the index of the peak element.

Your task is to solve it in O(log(n)) time complexity.

 

Example 1:

Input: arr = [0,1,0]

Output: 1

Example 2:

Input: arr = [0,2,1,0]

Output: 1

Example 3:

Input: arr = [0,10,5,2]

Output: 1

 

Constraints:

3 <= arr.length <= 105
0 <= arr[i] <= 106
arr is guaranteed to be a mountain array.

# Code
```go
package main

import "fmt"

func main() {
	arr := []int{0, 3, 5, 12, 2}
	res := peakIndexInMountainArray(arr)
	fmt.Println(res)
}

// 1. 二分查找
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr) // 初始化左右指针

	// 使用二分查找来定位峰值
	for left < right {
		mid := left + (right-left)/2 // 计算中间索引
		// 如果当前中间元素是峰值，则返回该索引
		if mid-1 >= 0 && mid+1 < len(arr) && arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if mid-1 >= 0 && arr[mid] > arr[mid-1] || mid-1 < 0 { // 如果中间元素大于前一个元素或者是数组的第一个元素，移动左指针
			left = mid + 1
		} else if mid+1 < len(arr) && arr[mid] > arr[mid+1] || mid+1 == len(arr) { // 如果中间元素大于后一个元素或者是数组的最后一个元素，移动右指针
			right = mid
		}
	}

	// 如果搜索结束后左指针超出数组范围，将左指针调整为数组的最后一个元素的索引
	if left == len(arr) {
		left = left - 1
	}

	return left // 返回左指针，此时应指向峰值元素
}
```