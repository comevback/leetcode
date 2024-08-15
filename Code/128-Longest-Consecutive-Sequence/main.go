package main

import "math"

func main() {

}

// longestConsecutive 用于查找数组中最长连续元素序列的长度。
func longestConsecutive(nums []int) int {
	// 创建一个哈希表，用来存储每个数字以及一个标志位表示该数字是否已被访问。
	hsMap := make(map[int]bool)
	// 初始化最小值和最大值，这里使用了math包中的MaxInt和MinInt以处理极端情况。
	min, max := math.MaxInt, math.MinInt

	// 遍历数组，填充哈希表，并更新数组中的最小值和最大值。
	for _, v := range nums {
		hsMap[v] = true
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	// 初始化当前序列长度和最大序列长度变量。
	length := 0
	maxLength := 0

	// 再次遍历数组，寻找最长的连续序列。
	for _, v := range nums {
		// 重置当前长度为0。
		length = 0
		// 查找当前元素v的前一个元素是否存在，如果存在则继续向前查找，直到找到连续序列的起点。
		for hsMap[v-1] {
			v -= 1
		}

		// 从连续序列的起点开始，向后查找并更新连续序列的长度，同时将访问过的数字标记为已访问。
		for hsMap[v] {
			hsMap[v] = false // 标记为已访问
			length += 1
			// 如果当前连续序列长度超过了之前记录的最大长度，则更新最大长度。
			if length > maxLength {
				maxLength = length
			}
			v += 1 // 移动到下一个数字
		}
	}

	// 返回最大连续序列的长度。
	return maxLength
}
