package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := "ab"
	word2 := "pqrs"
	result := mergeAlternately(word1, word2)
	fmt.Println(result)
}

// 1. 从前往后遍历，将两个字符串的字符交替合并, 用strings.Builder拼接
func mergeAlternately(word1 string, word2 string) string {
	// 获取两个字符串的长度, 取最大值为length，作为循环次数
	length1 := len(word1)
	length2 := len(word2)
	var length int
	if length1 >= length2 {
		length = length1
	} else {
		length = length2
	}

	// 创建一个strings.Builder对象，用于拼接字符串
	var builder strings.Builder
	// 遍历两个字符串，将字符交替拼接到builder中
	for i := 0; i < length; i++ {
		if i < length1 {
			builder.WriteByte(word1[i])
		}
		if i < length2 {
			builder.WriteByte(word2[i])
		}
	}

	// 将builder中的内容转换为字符串返回
	res := builder.String()
	return res
}

// ================================================================================================================================
// other solution
func mergeAlternately1(word1 string, word2 string) string {
	var ans string
	for len(word1) > 0 && len(word2) > 0 {
		ans += string(word1[0]) + string(word2[0])
		word1 = word1[1:]
		word2 = word2[1:]
	}
	if len(word1) > 0 {
		ans += word1
	}
	if len(word2) > 0 {
		ans += word2
	}

	return ans
}
