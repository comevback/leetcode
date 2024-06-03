# 974. Subarray Sums Divisible by K
Solved
Medium
Topics
Companies
Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.

A subarray is a contiguous part of an array.


Example 1:
> Input: nums = [4,5,0,-2,-3,1], k = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by k = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

Example 2:
> Input: nums = [5], k = 9
Output: 0

Constraints:

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
2 <= k <= 104

---

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{-1, 2, 9}
	res := subarraysDivByK(nums, 2)
	fmt.Println(res)
}

func subarraysDivByK(nums []int, k int) int {
	subNums := 0                    // 初始化计数器，用于记录满足条件的子数组个数
	hsMap := make(map[int][]int)    // 哈希表，用于存储前缀和对 k 取模的结果及其对应的索引
	hsMap[0] = append(hsMap[0], -1) // 初始化哈希表，前缀和为 0 的初始索引为 -1
	sum := 0                        // 初始化累加和

	// 遍历 nums 数组，计算每个元素的前缀和
	for i := 0; i < len(nums); i++ {
		sum += nums[i] // 累加当前元素

		// 计算当前前缀和对 k 取模的结果，并确保结果为非负数
		left := ((sum % k) + k) % k

		// 如果哈希表中存在相同的前缀和对 k 取模的结果，则增加满足条件的子数组个数
		if len(hsMap[left]) > 0 {
			subNums += len(hsMap[left])
		}

		// 将当前前缀和对 k 取模的结果及其索引存入哈希表
		hsMap[left] = append(hsMap[left], i)
	}

	return subNums // 返回满足条件的子数组个数
}
```