# 325. 和等于 k 的最长子数组长度
给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。

示例 1:

输入: nums = [1,-1,5,-2,3], k = 3
输出: 4 
解释: 子数组 [1, -1, 5, -2] 和等于 3，且长度最长。
示例 2:

输入: nums = [-2,-1,2,1], k = 1
输出: 2 
解释: 子数组 [-1, 2] 和等于 1，且长度最长。
提示：

1 <= nums.length <= 2 * 105
-104 <= nums[i] <= 104
-109 <= k <= 109

---

# Code
```go
package main

import "fmt"

func main() {
	nums := []int{1, -1, 5, -2, 3}
	k := 3
	res := MaxSubArrayLen(nums, k)
	fmt.Println(res)
}

func MaxSubArrayLen(nums []int, k int) int {
	maxLength := 0
	hsMap := make(map[int][]int)
	sum := 0
	hsMap[0] = append(hsMap[0], -1)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		need := sum - k
		if len(hsMap[need]) > 0 {
			length := i - hsMap[need][0]
			if length > maxLength {
				maxLength = length
			}
		}
		hsMap[sum] = append(hsMap[sum], i)
	}

	return maxLength
}
```