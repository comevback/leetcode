package main

import "fmt"

func main() {
	s := "ABBB"
	res := characterReplacement(s, 2)
	fmt.Println(res)
}

// characterReplacement 计算在最多可以替换 k 个字符的情况下，字符串 s 中最长的重复字符子串的长度。
func characterReplacement(s string, k int) int {
	// 用于记录窗口内各个字符的出现次数
	charCount := [26]int{}
	// 窗口内单个字符的最大出现次数
	maxCount := 0
	// 最长重复字符子串的长度
	maxLength := 0

	// 窗口的左右边界
	left, right := 0, 0

	// 遍历字符串 s
	for right < len(s) {
		// 增加窗口右边界的字符计数
		charCount[s[right]-'A'] += 1
		// 更新窗口内单个字符的最大出现次数
		if charCount[s[right]-'A'] > maxCount {
			maxCount = charCount[s[right]-'A']
		}
		// 扩大窗口右边界
		right += 1

		// 如果需要替换的字符数超过了 k，缩小窗口左边界
		for right-left-maxCount > k {
			// 减少窗口左边界的字符计数
			charCount[s[left]-'A'] -= 1
			// 缩小窗口左边界
			left += 1
		}
		// 这里不用更新 maxCount
		// 因为只有 maxCount 变得更大的时候才可能获得更长的重复子串，才会更新 maxLength

		// 计算当前窗口的长度
		length := right - left
		// 更新最长重复字符子串的长度
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
