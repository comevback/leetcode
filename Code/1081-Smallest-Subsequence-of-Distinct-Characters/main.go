package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "cbacdcbc"               // 测试字符串
	res := smallestSubsequence(str) // 获取字典序最小的子序列
	fmt.Println(res)                // 输出结果
}

// smallestSubsequence 函数返回字符串 s 的字典序最小子序列
func smallestSubsequence(s string) string {
	hsMap := make(map[rune][]int) // 创建一个哈希表，用于存储每个字符的所有出现位置
	for i, v := range s {
		hsMap[v] = append(hsMap[v], i) // 将字符位置添加到哈希表中对应的切片
	}
	res := []rune{}        // 创建一个 rune 切片用于存放结果
	inStack := [256]bool{} // 创建一个固定大小的布尔数组，用于标记字符是否已经在结果中

	for i, char := range s { // 遍历字符串中的每个字符
		if inStack[char] { // 如果当前字符已经在结果中，则跳过
			continue
		}

		// 当结果不为空，且结果中最后一个字符大于当前字符，而且该字符在后面还会出现时，弹出结果中的字符
		for len(res) > 0 && res[len(res)-1] > char && hsMap[res[len(res)-1]][len(hsMap[res[len(res)-1]])-1] > i {
			inStack[res[len(res)-1]] = false // 标记为不在结果中
			res = res[:len(res)-1]           // 移除结果中的最后一个字符
		}

		inStack[char] = true    // 标记当前字符为在结果中
		res = append(res, char) // 将当前字符添加到结果中
	}

	var builder strings.Builder // 创建一个字符串构建器，用于构建最终的字符串结果
	for _, v := range res {     // 遍历结果中的每个字符
		builder.WriteRune(v) // 将字符写入字符串构建器
	}

	return builder.String() // 返回构建的字符串
}
