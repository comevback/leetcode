# 1011. Capacity To Ship Packages Within D Days
Solved
Medium
Topics
Companies
Hint
A conveyor belt has packages that must be shipped from one port to another within days days.

The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship with packages on the conveyor belt (in the order given by weights). We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within days days.

 

Example 1:

Input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5
Output: 15
Explanation: A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10

Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.
Example 2:

Input: weights = [3,2,2,4,1,4], days = 3
Output: 6
Explanation: A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
1st day: 3, 2
2nd day: 2, 4
3rd day: 1, 4
Example 3:

Input: weights = [1,2,3,1,1], days = 4
Output: 3
Explanation:
1st day: 1
2nd day: 2
3rd day: 3
4th day: 1, 1
 

Constraints:

1 <= days <= weights.length <= 5 * 104
1 <= weights[i] <= 500


# Code
```go
package main

import "fmt"

func main() {
	// 测试用例，包含货物重量数组和期望的运输天数
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := shipWithinDays(weights, 15)
	fmt.Println(res) // 输出计算得到的最小载重量
}

// shipWithinDays 计算在指定天数内完成运输所需的最小载重量
func shipWithinDays(weights []int, days int) int {
	// 确定载重量的搜索范围
	left := findMax(weights)          // 最小载重量不能小于货物中的最大重量
	right := 50000 * findMax(weights) // 一个假设的上界，为了简化计算

	// 二分查找最小载重量
	for left < right {
		mid := left + (right-left)/2
		if countDays(weights, mid) <= days {
			right = mid // 如果当前载重量可以在指定天数内完成运输，则尝试更小的载重量
		} else {
			left = mid + 1 // 如果不能，则增加载重量
		}
	}

	return left // 返回计算得到的最小载重量
}

// countDays 计算给定载重量下完成运输所需的天数
func countDays(weights []int, cap int) int {
	days := 0     // 初始化天数计数
	amount := cap // 初始化当前载重量

	// 遍历所有货物重量，计算总天数
	for i := 0; i < len(weights); i++ {
		if amount-weights[i] >= 0 {
			amount -= weights[i] // 如果当前载重量可以装下货物，则继续装载
		} else {
			days += 1 // 否则，增加一天，并重新开始装载
			amount = cap - weights[i]
		}
	}

	days += 1 // 最后一天的装载
	return days
}

// findMax 找到数组中的最大值，用于确定载重量的下界
func findMax(weight []int) int {
	max := weight[0]
	for _, v := range weight {
		if v > max {
			max = v
		}
	}
	return max
}

// =========================================================================================================
// shipWithinDays1 和 shipWithinDays 函数类似，但是优化了载重量的上界
// 使用了货物总重量作为上界，而不是一个固定的倍数
func shipWithinDays1(weights []int, days int) int {
	var left, right int
	for _, w := range weights {
		left = max(left, w) // 载重量下界
		right += w          // 载重量上界
	}

	// 二分查找
	for left < right {
		mid := left + (right-left)/2
		if countDays1(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// countDays1 类似于 countDays，但为优化版本提供支持
func countDays1(weights []int, cap int) int {
	return countDays(weights, cap) // 直接调用 countDays 函数
}

// max 返回两个整数中的最大值
func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
```