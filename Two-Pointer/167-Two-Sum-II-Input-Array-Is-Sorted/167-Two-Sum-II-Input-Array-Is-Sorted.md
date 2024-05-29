# 167. Two Sum II - Input Array Is Sorted
Solved
Medium
Topics
Companies
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.


Example 1:
> Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
> Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
> Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
 

Constraints:

2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers is sorted in non-decreasing order.
-1000 <= target <= 1000
The tests are generated such that there is exactly one solution.

---

# Code
```go
package main

func main() {

}

// 1.我的解法，快慢指针同时从左开始走
func twoSum(numbers []int, target int) []int {
	res := []int{}

	slow, fast := 0, 1

	for fast < len(numbers) {
		if numbers[slow]+numbers[fast] >= target {
			if numbers[slow]+numbers[fast] == target {
				res = append(res, slow+1)
				res = append(res, fast+1)
				break
			} else {
				slow -= 1
			}
		} else {
			fast += 1
			slow += 1
		}
	}

	return res
}

// 2. 前后指针分别从左右端向中间靠拢
func twoSum1(nums []int, target int) []int {
	// 一左一右两个指针相向而行
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			// 题目要求的索引是从 1 开始的
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++ // 让 sum 大一点
		} else if sum > target {
			right-- // 让 sum 小一点
		}
	}
	return []int{-1, -1}
}
```