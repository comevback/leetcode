package main

import (
	"fmt"
	"strings"
)

func main() {
	var haystack string = "abc"
	var needle string = "c"

	fmt.Println(strStr(haystack, needle))
}

// 1. 遍历对比
func strStr(haystack string, needle string) int {

	// 如果两个字符串相等，直接返回0
	if haystack == needle {
		return 0
	}

	length := len(needle)

	// i < len(haystack)-length+1 是关键，遍历到最后一位恰恰到达len(haystack)-1
	for i := 0; i < len(haystack)-length+1; i++ {
		// 如果在haystack中有和needle一样的字符串，找到就直接返回
		if haystack[i:i+length] == needle {
			return i
		}
	}
	// 如果没找到，返回-1
	return -1
}

// 2. 用标准库的strings.Index方法
// strings.Index 返回字符串在另一个字符串中第一次出现的索引，如果没有找到，返回-1
func strStrSample1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
