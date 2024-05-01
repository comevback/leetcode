# 334. Increasing Triplet Subsequence（双指针）

Medium

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
 

Example 1:
> Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:
> Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.

Example 3:
> Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
 

Constraints:
> 1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1
 

Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?

---

```go
package main

import "fmt"

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	fmt.Println(increasingTriplet(nums))
	nums = []int{1, 5, 0, 4, 1, 3}
	fmt.Println(increasingTriplet(nums))
}

// ================================================================================================================================
// 错误尝试：
// 1.想在循环中定义最大最小值，如果最小值在前面，然后又遇到新的最大值，则是有这样的数组，但这三个数并不一定是由最小值开始，到最大数构成，所以错误，例如 [20,100,10,12,5,13]。
func increasingTriplet1(nums []int) bool {
	length := len(nums)
	var smallest int = nums[0]
	var smallestIndex int = 0
	var biggest int = nums[0]
	//var biggestIndex int = 0

	for i := 1; i < length; i++ {
		if nums[i] < smallest {
			smallest = nums[i]
			smallestIndex = i
		}
		if i-1 > smallestIndex && nums[i] > biggest {
			return true
		}
		if nums[i] > biggest {
			biggest = nums[i]
			// biggestIndex = i
		}

	}

	return false
}

// ================================================================================================================================
// 2.我的解法：
// 设置一个最小值和一个中间值，当发现一个递增关系时，给这两个值赋值为 smallest 和 middle，信号 preFound 为true,才开始赋值。
// 当循环中而且信号为 true 时：
// 如果发现一个数比 middle 大，则返回 true，说明已经找到一个数大于 middle 大于 smallest 且递增
// 如果发现一个数比 smallest 大且比 middle 小，则更新 middle 为这个数
// 如果发现一个数比 smallest 小，则更新 smallest 为这个数
func increasingTriplet2(nums []int) bool {
	var smallest int = nums[0]
	var middle int = nums[0]
	var preFound bool = false

	for i := 1; i < len(nums); i++ {
		if preFound {
			if nums[i] > middle {
				return true
			}
			if nums[i] > smallest && nums[i] < middle {
				middle = nums[i]
			}
			if nums[i] < smallest {
				smallest = nums[i]
			}
		}
		if nums[i] > nums[i-1] {
			if preFound {
				if nums[i] > middle {
					return true
				} else {
					middle = nums[i]
					smallest = nums[i-1]
				}
			} else {
				middle = nums[i]
				smallest = nums[i-1]
				preFound = true
			}
		}
	}

	return false

}

// ================================================================================================================================
// 3.官方解法
// 1.初始化最小值 smallest 为int类型可能的最大值
// 2.初始化次小值 middle 也为int类型可能的最大值
// 3.遍历数组中的每个元素
// 4.如果当前数字小于或等于最小值，更新最小值
// 5.如果当前数字小于或等于次小值但大于最小值，更新次小值
// 6.如果找到一个数字大于已记录的最小值和次小值，返回true
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false // 如果数组长度小于3，直接返回false，因为不可能形成三元组
	}

	smallest := int(^uint(0) >> 1) // 初始化最小值为int类型可能的最大值
	middle := int(^uint(0) >> 1)   // 初始化次小值也为int类型可能的最大值

	for _, num := range nums { // 遍历数组中的每个元素
		if num <= smallest {
			smallest = num // 更新最小值，如果当前数字小于或等于最小值
		} else if num <= middle {
			middle = num // 更新次小值，如果当前数字小于或等于次小值但大于最小值
		} else {
			return true // 如果找到一个数字大于已记录的最小值和次小值，返回true
		}
	}

	return false // 如果遍历完成后没有找到符合条件的三元组，返回false
}

// 我的解法和官方解法之前的区别在于，我的解法是在发现递增关系时才开始赋值，而官方解法是在遍历数组时就开始赋值，这样可以减少一些判断。
// 之所以可以一开始就赋值，是因为 smallest 和 middle 的初始值是int类型可能的最大值，所以
if nums[i] > middle {
    return true
}
// 这个判断不会在一开始就被触发，也就没必要等到发现递增关系再赋值。
```
