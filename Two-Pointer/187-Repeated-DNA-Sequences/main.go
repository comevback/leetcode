package main

import (
	"fmt"
)

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	res := findRepeatedDnaSequences2(s)
	fmt.Println(res)
}

// 1. 简单解法
func findRepeatedDnaSequences1(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	strMap := make(map[string]int) // 存储每个子串及其出现次数
	res := []string{}              // 存储结果
	left, right := 0, 10           // 双指针用于遍历子串

	for right <= len(s) {
		strMap[s[left:right]] += 1 // 更新当前子串的计数
		right += 1
		left += 1
	}

	for key, value := range strMap { // 遍历map查找计数大于1的子串
		if value > 1 {
			res = append(res, key)
		}
	}

	return res
}

// 2. 头尾拼接
func findRepeatedDnaSequences2(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	strMap := make(map[string]bool) // 存储每个子串及其出现次数
	resMap := make(map[string]bool)
	res := []string{}    // 存储结果
	left, right := 0, 10 // 双指针用于遍历子串
	subStr := s[left:right]

	for right <= len(s) {
		if strMap[subStr] {
			resMap[subStr] = true
		} else {
			strMap[subStr] = true
		}
		if right == len(s) {
			break
		}
		subStr = subStr[1:]
		subStr = subStr + string(s[right])
		right += 1
	}

	for key, value := range resMap { // 遍历map查找计数大于1的子串
		if value {
			res = append(res, key)
		}
	}

	return res
}

// 3. 使用最高最低位处理
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	hsMap := make(map[int]bool)     // 存储子串的哈希值
	resMap := make(map[string]bool) // 存储已发现的重复子串
	res := []string{}               // 存储结果
	left, right := 0, 0
	currentNum := 0 // 当前子串的数值表示

	for right <= len(s) {
		if right-left == 10 {
			if hsMap[currentNum] { // 如果当前哈希值已存在
				resMap[s[left:right]] = true // 标记子串为重复
			} else {
				hsMap[currentNum] = true
			}
			minus := 0
			switch s[left] { // 根据最左边的字符减小数值
			case 'A':
				minus = 1
			case 'C':
				minus = 2
			case 'G':
				minus = 3
			case 'T':
				minus = 4
			}

			currentNum = currentNum - times(minus, 10) // 更新当前数值
			left += 1
		}

		if right < len(s) {
			add := 0
			switch s[right] { // 根据新字符增加数值
			case 'A':
				add = 1
			case 'C':
				add = 2
			case 'G':
				add = 3
			case 'T':
				add = 4
			}

			currentNum = currentNum*10 + add // 更新当前数值
			right += 1
		} else {
			break
		}
	}

	for key, value := range resMap { // 遍历结果map添加到最终结果列表
		if value {
			res = append(res, key)
		}
	}

	return res
}

// 计算数值的n次方
func times(num int, times int) int {
	for i := 1; i < times; i++ {
		num *= 10
	}
	return num
}
