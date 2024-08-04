# 29. Divide Two Integers

Solved
Medium
Topics
Companies
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = -2.33333.. which is truncated to -2.

Constraints:

-231 <= dividend, divisor <= 231 - 1
divisor != 0

---

# Code

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	res := divide(-2147483648, 4)
	fmt.Println(res) // 输出: -536870912
}

// divide 实现了整数除法，返回值为商
func divide(dividend int, divisor int) int {
	// 判断结果的符号
	minus := false
	if dividend*divisor < 0 {
		minus = true
	}

	// 取被除数和除数的绝对值
	a := abs(dividend)
	b := abs(divisor)

	res := 0

	// 使用位移操作进行快速除法
	for a >= b {
		temp, multi := b, 1
		// 找到最大的倍数，使得 a >= temp << 1
		for a >= (temp << 1) {
			temp = temp << 1   // temp 倍增
			multi = multi << 1 // multi 倍增
		}
		a -= temp    // 减去找到的最大倍数
		res += multi // 将倍数加到结果中
	}

	// 如果结果是负数，则取负
	if minus {
		res = -res
	}

	// 防止溢出
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	if res < math.MinInt32 {
		res = math.MinInt32
	}

	return res
}

// abs 返回整数的绝对值
func abs(nums int) int {
	if nums < 0 {
		return -nums
	}
	return nums
}
```
