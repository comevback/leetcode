# 1679. Max Number of K-Sum Pairs（map）（排序+双指针）

> 在这个问题中，无法用数组代替map和排序。
要用数组代替map或排序，必须条件是，数组people有一个明确的最大值，在（881. Boats to Save People）这个问题里，最大值就是limit，所以这个数组是可预期的有限长度。但是在这个问题里，数组nums没有明确的最大值，nums可能极长，所以无法用数组代替map和排序。

Medium

Hint
You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.

Example 1:
> Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.

Example 2:
> Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.
 

Constraints:
> 1 <= nums.length <= 105
1 <= nums[i] <= 109
1 <= k <= 109

---

## 思路
1. 用和 (1-two-sum) 一样的思路，用map里的key存储nums\[i]，value存储k-nums\[i]的个数，然后遍历nums，如果map里有nums\[i]，则maxOperation+1，并且map里的value-1
2. 将数组排序，用一头一尾两个指针从两端向中间遍历，如果nums\[head]+nums\[tail]==k，则maxOperation+1，head+1，tail-1。

## 代码
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int = []int{3, 5, 1, 5}
	fmt.Println(maxOperations1(nums, 2))
}

// 1. map
// 遍历每个元素，元素值为v，如果hsMap[k-v]存在，这个元素不加入map，同时hsMap[k-v]-1，最大操作数加1
// 如果hsMap[k-v]不存在，则hsMap[v]加1
func maxOperations(nums []int, k int) int {
	var hsMap map[int]int = make(map[int]int)
	maxOp := 0

	for _, v := range nums {
		if hsMap[k-v] > 0 {
			hsMap[k-v] -= 1
			maxOp += 1
		} else {
			hsMap[v] += 1
		}
	}

	return maxOp
}

// 2.双指针 排序
// 先排序，然后使用双指针，head指向头，tail指向尾，从两侧向中间遍历，如果nums[head]+nums[tail] == k，head+1，tail-1，maxOp+1
func maxOperations1(nums []int, k int) int {
	// 定义最大操作数，头尾指针
	maxOp := 0
	var head, tail int = 0, len(nums) - 1
	// 把数组从小到大排序
	sort.Ints(nums)

	// 用双指针从两侧向中间遍历
	for head < tail {
		// 如果头尾指针的和大于k，尾指针向左移动
		if nums[head]+nums[tail] > k {
			tail -= 1
		}

		// 如果头尾指针的和小于k，头指针向右移动
		if nums[head]+nums[tail] < k {
			head += 1
		}

		// 如果前面的遍历操作导致头尾指针相遇，结束遍历
		if head > tail {
			break
		}

		// 如果头尾指针的和等于k，头尾指针同时向中间移动，最大操作数加1
		if nums[head]+nums[tail] == k && head != tail {
			maxOp += 1
			head += 1
			tail -= 1
		}
	}

	// 返回最大操作数
	return maxOp
}

// 在这个问题中，无法用数组代替map和排序，
// 要用数组代替map或排序，必须条件是，数组people有一个明确的最大值，
// 在（881. Boats to Save People）这个问题里，最大值就是limit，所以这个数组是可预期的有限长度。
// 但是在这个问题里，数组nums没有明确的最大值，nums可能极长，所以无法用数组代替map和排序。
```