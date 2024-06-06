package main

import (
	"fmt"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	// 初始化需求字符和窗口字符的映射
	reqMap := make(map[byte]int)
	windowMap := make(map[byte]int)
	// 有效字符计数
	valid := 0
	// 初始化最小窗口为一个足够长的字符串
	minWindow := s + t

	// 构建需求字符映射
	for i := 0; i < len(t); i++ {
		reqMap[t[i]] += 1
	}

	// 左右指针初始化
	left, right := 0, 0

	// 开始滑动窗口
	for right < len(s) || left < len(s) {
		// 如果有效字符数小于需求字符数且右指针超出范围，退出循环
		if valid < len(reqMap) && right >= len(s) {
			break
		}

		// 如果有效字符数等于需求字符数
		if valid == len(reqMap) {
			// 检查左指针字符是否在需求字符映射中
			if reqMap[s[left]] > 0 {
				// 记录当前窗口
				window := s[left:right]
				// 更新最小窗口
				if len(window) < len(minWindow) {
					minWindow = window
				}
				// 减少窗口字符映射中的计数
				windowMap[s[left]] -= 1
				// 如果窗口字符计数小于需求字符计数，减少有效字符计数
				if windowMap[s[left]] < reqMap[s[left]] {
					valid -= 1
				}
			} else {
				// 当前窗口不包含需求字符
				window := s[left:right]
				if len(window) < len(minWindow) {
					minWindow = window
				}
			}
			// 左指针右移
			left += 1
		} else {
			// 如果右指针字符在需求字符映射中
			if right < len(s) && reqMap[s[right]] > 0 {
				// 增加窗口字符映射中的计数
				windowMap[s[right]] += 1
				// 如果窗口字符计数等于需求字符计数，增加有效字符计数
				if windowMap[s[right]] == reqMap[s[right]] {
					valid += 1
				}
			}
			// 右指针右移
			right += 1
		}
	}

	// 如果最小窗口未更新，返回空字符串
	if minWindow == s+t {
		return ""
	} else {
		return minWindow
	}
}
