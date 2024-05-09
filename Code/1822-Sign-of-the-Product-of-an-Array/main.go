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
