# 1822. Sign of the Product of an Array

Easy

Hint
There is a function signFunc(x) that returns:

1 if x is positive.
-1 if x is negative.
0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).


Example 1:
> Input: nums = [-1,-2,-3,-4,3,2,1]
Output: 1
Explanation: The product of all values in the array is 144, and signFunc(144) = 1

Example 2:
> Input: nums = [1,5,0,2,-3]
Output: 0
Explanation: The product of all values in the array is 0, and signFunc(0) = 0

Example 3:
> Input: nums = [-1,1,-1,1,-1]
Output: -1
Explanation: The product of all values in the array is -1, and signFunc(-1) = -1
 

Constraints:
> 1 <= nums.length <= 1000
-100 <= nums[i] <= 100

---

# Code
```go
package main

func main() {

}

// 直接累乘
func arraySign(nums []int) int {
	res := 1
	for _, v := range nums {
		// 如果遇到0，直接返回0
		if v == 0 {
			return 0
		} else if v > 0 {
			// 如果遇到正数，乘1
			res *= 1
		} else {
			// 如果遇到负数，乘-1
			res *= -1
		}
	}
	// 返回结果
	return res
}

// 计算负数个数
func arraySign1(nums []int) int {
	// 负数个数
	minusCount := 0
	for _, v := range nums {
		// 如果遇到0，直接返回0
		if v == 0 {
			return 0
		}
		// 如果遇到负数，minusCount
		if v < 0 {
			minusCount += 1
		}
	}
	// 如果负数个数是偶数，包括0，累乘一定是正的，返回1
	if minusCount%2 == 0 {
		return 1
	}
	// 否则返回-1
	return -1
}
```
