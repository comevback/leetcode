package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage(arr, 4))
}

// 多次循环，时间复杂度太高
func findMaxAverage1(nums []int, k int) float64 {
	maxAverage := float64(-10000)
	if len(nums) == 0 {
		panic("empty")
	}
	for i := 0; i < len(nums)-k+1; i++ {
		if maxAverage < getAverage(nums[i:i+k]) {
			maxAverage = getAverage(nums[i : i+k])
		}
	}
	return maxAverage
}

func getAverage(nums []int) float64 {
	var sum float64
	for _, v := range nums {
		sum += float64(v)
	}

	average := sum / float64(len(nums))
	return average
}

// ================================================================================================================

// 滑动窗口，每次循环直接减去前一个元素，加上后一个元素
func findMaxAverage(nums []int, k int) float64 {
	// 求前k个元素的和
	var sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	// 求前k个元素的平均值，作为初始化的最大平均值
	maxAverage := sum / float64(k)
	// 定义当前平均值
	currentAve := maxAverage

	// 滑动窗口，每次循环直接减去前一个元素，加上后一个元素，如果这个平均值大于最大平均值，则更新最大平均值
	for i := 0; i < len(nums)-k; i++ {
		currentAve = currentAve - float64(nums[i])/float64(k) + float64(nums[i+k])/float64(k)
		if currentAve > maxAverage {
			maxAverage = currentAve
		}
	}

	// 返回最大平均值
	return maxAverage
}
