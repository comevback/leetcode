package main

import "fmt"

func main() {
	records := []int{0, 1, 2, 3, 5}
	res := takeAttendance(records)
	fmt.Println(res)
}

// 1. 二分查找
func takeAttendance(records []int) int {
	left, right := 0, len(records) // 初始化左右指针

	// 当左指针小于右指针时，进行循环
	for left < right {
		mid := left + (right-left)/2 // 计算中间索引

		// 如果中间元素的值大于其索引，则需要在左半部分继续搜索
		if records[mid] > mid {
			right = mid // 将右指针移动到中间位置
		} else if records[mid] == mid {
			// 如果中间元素的值等于其索引，则缺失的元素在右半部分
			left = mid + 1 // 将左指针移动到中间位置的右侧
		}
	}

	// 当循环结束时，left 将指向第一个元素值不等于其索引的位置
	return left
}
