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
