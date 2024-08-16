package main

func main() {

}

// reverseBits 反转一个 int32 类型整数的所有位
func reverseBits(num int32) int32 {
	var result int32 = 0
	for i := 0; i < 32; i++ {
		// 将 result 向左移一位，为下一位腾出空间
		result <<= 1
		// 获取 num 的最低位
		lowest := num & 1
		// 将这个位添加到 result 的最低位
		result |= lowest
		// 将 num 向右移动一位，准备下一次迭代
		num >>= 1
	}
	return result
}
