# 209. Minimum Size Subarray Sum

Solved
Medium
Topics
Companies
Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

Constraints:

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104

Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

---

# Code

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(7, nums)
	fmt.Println(res) // 输出最小子数组的长度
}

// minSubArrayLen 寻找至少和为 target 的最小连续子数组长度
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0   // 初始化左右指针
	minLen := math.MaxInt // 最小长度初始化为最大整数值，用于后续比较
	sum := 0              // 当前子数组的和

	// 主循环，直到右指针越过数组末端
	for right <= len(nums) {
		if sum < target {
			// 如果当前子数组和小于目标值，右指针右移，扩展窗口
			if right < len(nums) {
				sum += nums[right] // 累加新包含的元素
			}
			right += 1 // 右指针右移
		} else {
			// 当前子数组和大于等于目标值
			if right-left < minLen {
				minLen = right - left // 更新最小长度
			}
			if left < len(nums) {
				sum -= nums[left] // 从和中减去最左边的元素
			}
			left += 1 // 左指针右移，缩小窗口
		}
	}

	// 如果 minLen 仍然是最大整数值，说明没有找到合适的子数组
	if minLen == math.MaxInt {
		return 0
	} else {
		return minLen // 返回找到的最小子数组长度
	}
}
```
