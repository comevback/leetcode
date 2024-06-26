package main

import "fmt"

func main() {
	// 测试用例
	nums := []int{1, 1}
	res := splitArray(nums, 1) // 拆分数组，使得每组中的和不超过最大可能值
	fmt.Println(res)           // 输出结果
}

// splitArray 计算在给定组数k的情况下，数组能被分成的最小最大组和
func splitArray(nums []int, k int) int {
	var left, right int
	for _, v := range nums {
		if v > left {
			left = v // 确保分组的最小和至少为数组中的最大值
		}
		right += v // 所有元素和，即不分组时的数组总和
	}

	// 二分查找最小的最大组和
	for left < right {
		mid := left + (right-left)/2 // 计算中点，作为尝试的组和上限
		if getGroups(nums, mid) <= k {
			right = mid // 如果当前最大组和可以满足分组条件，尝试更小的最大组和
		} else if getGroups(nums, mid) > k {
			left = mid + 1 // 如果不能满足，增加最大组和
		}
	}

	return left // 返回最小可能的最大组和
}

// getGroups 计算给定最大组和限制下，数组需要分成的最少组数
func getGroups(nums []int, maxSum int) int {
	days := 0   // 组数计数器
	amount := 0 // 当前组的累计和

	// 遍历数组元素，决定分组
	for i := 0; i < len(nums); i++ {
		if amount+nums[i] <= maxSum {
			amount += nums[i] // 当前组累加当前元素
		} else {
			days += 1        // 开始新的一组
			amount = nums[i] // 当前元素成为新组的第一个元素
		}
	}

	// 处理最后一组（如果有）
	if amount > 0 {
		days += 1
	}

	return days // 返回总组数
}
