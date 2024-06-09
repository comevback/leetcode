package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 5, 2, 6}
	res := numSubarrayProductLessThanK(nums, 100)
	fmt.Println(res)
}

// 1. 我的方法，以每个元素为头，看能乘多少个。比较慢
func numSubarrayProductLessThanK1(nums []int, k int) int {
	arrNum := 0
	tempProdunt := 1

	begin, end := 0, 0

	for begin < len(nums) && end < len(nums) {
		if nums[begin] < k {
			arrNum += 1
			tempProdunt = nums[begin]
			end = begin + 1
			for end < len(nums) && tempProdunt*nums[end] < k {
				arrNum += 1
				tempProdunt *= nums[end]
				end += 1
			}
		}
		begin += 1
		end = begin
	}

	return arrNum
}

// 2.解法
func numSubarrayProductLessThanK(nums []int, k int) int {
	arrNum := 0
	tempProdunt := 1
	left, right := 0, 0

	for right < len(nums) {
		tempProdunt *= nums[right]
		right += 1

		for left < right && tempProdunt >= k {
			tempProdunt /= nums[left]
			left += 1
		}

		arrNum += right - left
	}

	return arrNum
}

// 现在，为什么 arrNum += right - left？

// 假设当前的窗口是从 left 到 right-1（因为 right 已经被增加了1）。那么窗口内的所有子数组都是有效的子数组。具体来说，包含 right-1 位置元素的子数组个数等于 right - left。

// 举个例子：

// 假设 nums = [10, 5, 2, 6]，k = 100。

// 当 right = 1 时，窗口 [10] 是有效子数组，arrNum 增加 1。
// 当 right = 2 时，窗口 [10, 5]，由于 10 * 5 = 50 < 100，有效子数组有 [10, 5] 和 [5]，所以 arrNum 增加 2。
// 当 right = 3 时，窗口 [10, 5, 2]，由于 10 * 5 * 2 = 100 >= 100，我们需要移动 left 使乘积小于 100，新的窗口为 [5, 2]，有效子数组有 [5, 2]，[2]，所以 arrNum 增加 2。
// 当 right = 4 时，窗口 [5, 2, 6]，有效子数组有 [5, 2, 6]，[2, 6]，[6]，所以 arrNum 增加 3。
// 通过 arrNum += right - left，我们计算出包含当前 right 元素的所有有效子数组的数量，并将其累加到 arrNum 中。
