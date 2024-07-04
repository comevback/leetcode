# 560. Subarray Sum Equals K
Solved
Medium
Topics
Companies
Hint
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.


Example 1:
> Input: nums = [1,1,1], k = 2
Output: 2

Example 2:
> Input: nums = [1,2,3], k = 3
Output: 2

Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107

---

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{1}
	res := subarraySum(nums, 0)
	fmt.Println(res)
}

func subarraySum(nums []int, k int) int {
	subArrNum := 0                     // 初始化计数器，用于记录满足条件的子数组个数
	prefix := make([]int, len(nums)+1) // 前缀和数组，长度比 nums 长 1，用于存储前缀和
	hsMap := make(map[int]int)         // 哈希表，用于存储前缀和出现的次数
	hsMap[0] += 1                      // 初始化哈希表，前缀和为 0 的次数为 1
	sum := 0                           // 初始化累加和

	// 遍历 nums 数组，计算每个元素的前缀和
	for i := 0; i < len(nums); i++ {
		sum += nums[i]    // 累加当前元素
		prefix[i+1] = sum // 存储当前前缀和到前缀和数组
		need := sum - k   // 计算需要的前缀和，使得当前前缀和减去需要的前缀和等于 k

		// 如果哈希表中存在需要的前缀和，则增加满足条件的子数组个数
		if hsMap[need] > 0 {
			subArrNum += hsMap[need]
		}

		hsMap[sum] += 1 // 将当前前缀和存入哈希表，出现次数加 1
	}

	return subArrNum // 返回满足条件的子数组个数
}
```