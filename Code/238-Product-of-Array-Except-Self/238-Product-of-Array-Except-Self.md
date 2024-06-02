# 238. Product of Array Except Self（乘积，优化时间复杂度）

Medium

Hint
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

 

Example 1:
> Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
> Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
 

Constraints:
> 2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

---

```go
package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{-1, 1, 0, -3, 3}
	res := productExceptSelf(nums)
	fmt.Println(res)
}

// ===============================================================================================================================
// 1. 多次分别循环，从前往后循环，计算每个元素前面的乘积，从后往前循环，计算每个元素后面的乘积，最后都结果数组中，每个元素等于它前面的乘积乘以后面的乘积
func productExceptSelf1(nums []int) []int {
	var zeroMap map[int]bool = make(map[int]bool)
	var zeroPos int
	var signal bool
	var res []int = make([]int, len(nums))
	productFront := 1
	productBack := 1

	for i, v := range nums {
		if v == 0 {
			if zeroMap[v] {
				return make([]int, len(nums))
			} else {
				zeroPos = i
				zeroMap[v] = true
				signal = true
			}
		} else {
			res[i] = productFront
			productFront *= v
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			res[i] *= productBack
			productBack *= nums[i]
		}
	}

	if signal {
		res[zeroPos] = productFront
		return res
	} else {
		return res
	}

}

// 优化空间，不使用额外的map，直接使用变量记录0的个数
func productExceptSelf(nums []int) []int {
	// 0的个数
	var zeroCount int = 0
	// 前面的乘积
	var frontProduct int = 1
	// 后面的乘积
	var backProduct int = 1
	// 定义结果数组，长度和输入数组一样，每个元素都初始化为1
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}

	// 定义头尾指针，头指针从0开始，尾指针从数组长度-1开始
	var head, tail int = 0, len(nums) - 1

	// 当头指针小于数组长度时，循环
	for head < len(nums) {
		// 如果0的个数等于2，直接返回全0数组
		if zeroCount == 2 {
			return make([]int, len(nums))
		}
		// 如果头指针的值等于0，0的个数加1
		if nums[head] == 0 {
			zeroCount += 1
		}
		// 前面的乘积乘以头指针的值，赋值给结果数组的头指针位置
		res[head] *= frontProduct
		frontProduct *= nums[head]
		// 后面的乘积乘以尾指针的值，赋值给结果数组的尾指针位置
		res[tail] *= backProduct
		backProduct *= nums[tail]
		// 头指针加1，尾指针减1
		head += 1
		tail -= 1
	}

	return res
}

// ============================================================================================================
// review 2024.6.2
func productExceptSelf3(nums []int) []int {
	product := 1 // 储存所有元素的乘积
	res := make([]int, len(nums)) // 结果数组
	zeroCount := 0 		// 0的个数
	zeroIndex := -1		// 0的位置

	// 计算所有元素的乘积，如果有0，第一次遇到0，记录0的位置，不加入乘积，第二次遇到0，直接返回全0数组
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			product *= nums[i]
		} else {
			if zeroCount == 0 {
				zeroCount += 1
				zeroIndex = i
			} else {
				return res
			}
		}
	}

	// 如果0的个数等于1，位于zeroIndex的位置的元素等于product
	if zeroCount == 1 {
		res[zeroIndex] = product
		return res
	}

	// 遍历数组，每个元素等于product除以当前元素的值
	for i := 0; i < len(nums); i++ {
		res[i] = product / nums[i]
	}

	// 返回结果数组
	return res
}
```