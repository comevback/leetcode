# 2441. Largest Positive Integer That Exists With Its Negative（哈希表）

Easy

Hint
Given an integer array nums that does not contain any zeros, find the largest positive integer k such that -k also exists in the array.

Return the positive integer k. If there is no such integer, return -1.

Example 1:
> Input: nums = [-1,2,-3,3]
Output: 3
Explanation: 3 is the only valid k we can find in the array.

Example 2:
> Input: nums = [-1,10,6,7,-7,1]
Output: 7
Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.

Example 3:
> Input: nums = [-10,8,6,7,-2,-3]
Output: -1
Explanation: There is no a single valid k, we return -1.
 

Constraints:
> 1 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
nums[i] != 0

---

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int = []int{-2, -3, -1, 1}
	fmt.Println(findMaxK(nums))
}

// 1.用map储存，把每个元素的值作为key，下标作为value
// 在遍历过程中，如果这个值的相反数作为key在map里出现过，看看这个值是否大于等于现有的max（因为初始值设为-1，如果数组中最大的有相反数正数为1，如果必须要求大于-1，那就不会替换max）
// 如果找到这个数，作为max返回，否则返回初始值-1

// map值为int
func findMaxK(nums []int) int {

	var hsMap map[int]int = make(map[int]int)
	var max int = -1

	for i, v := range nums {
		if _, exist := hsMap[-v]; exist {
			if v*v >= max*max {
				if v < 0 {
					max = -v
				} else {
					max = v
				}
			}
		}
		hsMap[v] = i
	}

	return max
}

// map值为bool
func findMaxK_bool(nums []int) int {

	var hsMap map[int]bool = make(map[int]bool)
	var max int = -1

	for _, v := range nums {
		if hsMap[-v] {
			if v*v >= max*max {
				if v < 0 {
					max = -v
				} else {
					max = v
				}
			}
		}
		hsMap[v] = true
	}

	return max
}

// 排序 from others on leetcode
// 也可以自己写排序算法
// 首先把数组排序，然后两个指针，从一头一尾开始，向中间遍历
// 因为一排序，所以左侧都是负数，右侧都是正数
// 两数相加，如果大于零，说明右侧正数过大，尾指针减一，向中间移动
// 两数相加，如果小于零，说明左侧负数过小，头指针加一，向中间移动
// 如果两数相加等于零，说明找到了最大有相反数的正数，返回
// 如果遍历完成还没有找到，说明每个数都没有相反数，返回-1
func findMaxK_sort(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] > 0 {
			j--
		} else if nums[i]+nums[j] < 0 {
			i++
		} else {
			return nums[j]
		}
	}
	return -1
}
```