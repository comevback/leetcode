package main

import "fmt"

func main() {
	arr := []int{1, 3, 2}
	fmt.Println(isMonotonic(arr))
}

// ========================================================================================
// 1. 分别检测递增递减情况
func isMonotonic(nums []int) bool {
	return increase(nums) || decrease(nums)
}

// 是否递减检测
func increase(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] >= 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

// 是否递增检测
func decrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] <= 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

// ========================================================================================
// 2. other solution by others in leetcode
func isMonotonic1(nums []int) bool {
	increasing := true
	decreasing := true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			decreasing = false
		} else if nums[i] < nums[i-1] {
			increasing = false
		}
		if !increasing && !decreasing {
			return false
		}
	}
	return true
}
