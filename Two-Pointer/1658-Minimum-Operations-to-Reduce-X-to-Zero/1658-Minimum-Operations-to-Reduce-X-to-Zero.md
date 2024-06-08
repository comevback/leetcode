# 1658. Minimum Operations to Reduce X to Zero
Solved
Medium
Topics
Companies
Hint
You are given an integer array nums and an integer x. In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x. Note that this modifies the array for future operations.

Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.


Example 1:
> Input: nums = [1,1,4,2,3], x = 5
Output: 2
Explanation: The optimal solution is to remove the last two elements to reduce x to zero.

Example 2:
> Input: nums = [5,6,7,8,9], x = 4
Output: -1

Example 3:
> Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 104
1 <= x <= 109

---

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{5, 1, 4, 2, 3}
	x := 6
	res := minOperations(nums, x)
	fmt.Println(res)
}

// minOperations 计算通过移除数组两端的元素，使其元素和等于 x 所需的最小操作数。
// 如果无法实现，返回 -1。
func minOperations(nums []int, x int) int {
	// 计算数组的总和。
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 需要的子数组的和。
	need := sum - x

	// 如果需要的和小于0，说明无法通过移除两端元素实现，返回 -1。
	if need < 0 {
		return -1
	} else if need == 0 {
		// 如果需要的和等于0，说明需要移除整个数组，返回数组长度。
		return len(nums)
	}

	// 用于临时存储当前窗口的和。
	tempSum := 0
	// 用于记录最大子数组长度。
	maxLength := 0
	// 定义窗口的左右指针。
	left, right := 0, 0

	// 使用滑动窗口技术找到和为 need 的最长子数组。
	for right < len(nums) {
		// 当 tempSum 等于需要的和时，计算当前窗口的长度并更新最大长度。
		if tempSum == need {
			length := right - left
			if length > maxLength {
				maxLength = length
			}
		}
		// 将右指针指向的元素加入 tempSum。
		tempSum += nums[right]
		right += 1
		// 当 tempSum 超过需要的和时，移动左指针以缩小窗口。
		for tempSum > need {
			tempSum -= nums[left]
			left += 1
			// 确保右指针不会落后于左指针。
			if right < left {
				right = left
			}
		}
	}

	// 检查最后一个窗口是否等于需要的和，并更新最大长度。
	if tempSum == need {
		length := right - left
		if length > maxLength {
			maxLength = length
		}
	}

	// 如果找到的最大子数组长度不为 0，返回最小操作数，否则返回 -1。
	if maxLength != 0 {
		return len(nums) - maxLength
	} else {
		return -1
	}
}
```