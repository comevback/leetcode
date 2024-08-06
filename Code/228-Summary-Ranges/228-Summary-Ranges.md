# 228. Summary Ranges

Solved
Easy
Topics
Companies
You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b

Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

Constraints:

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
All the values of nums are unique.
nums is sorted in ascending order.

---

# Code

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{-1}
	res := summaryRanges(nums)
	fmt.Println(res)
}

// summaryRanges 函数返回一个表示连续区间的字符串数组
func summaryRanges(nums []int) []string {
	res := []string{}   // 存储结果的字符串数组
	if len(nums) == 0 { // 如果输入数组为空，直接返回空结果
		return res
	}

	last := "" // 用于记录当前区间的起始数字

	for i := 0; i < len(nums); i++ {
		if last == "" { // 如果当前区间为空，设置起始数字
			numStr := strconv.Itoa(nums[i])
			last = numStr
		}

		// 处理最后一个元素的情况
		if i == len(nums)-1 {
			if i == 0 { // 如果只有一个元素，直接添加到结果中
				res = append(res, last)
				last = ""
				break
			}
			if nums[i] == nums[i-1]+1 { // 如果当前元素和前一个元素是连续的
				numStr := strconv.Itoa(nums[i])
				last += "->" + numStr   // 将当前元素添加到区间中
				res = append(res, last) // 将完整区间添加到结果中
				last = ""
			} else { // 如果当前元素和前一个元素不连续
				res = append(res, last) // 将当前区间添加到结果中
			}
			break
		}

		// 处理非连续情况
		if !(nums[i] == nums[i+1]-1) {
			if i == 0 { // 如果第一个元素不与后续元素连续，直接添加到结果中
				res = append(res, last)
				last = ""
				continue
			}
			if nums[i] == nums[i-1]+1 { // 如果当前元素和前一个元素是连续的
				numStr := strconv.Itoa(nums[i])
				last += "->" + numStr   // 将当前元素添加到区间中
				res = append(res, last) // 将完整区间添加到结果中
				last = ""
			} else { // 如果当前元素和前一个元素不连续
				res = append(res, last) // 将当前区间添加到结果中
				last = ""
			}
		}
	}

	return res
}
```
