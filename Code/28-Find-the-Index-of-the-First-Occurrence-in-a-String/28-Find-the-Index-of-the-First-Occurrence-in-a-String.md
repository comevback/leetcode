# 28. Find the Index of the First Occurrence in a String（KMP算法）（求LPS数组）

Easy

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

 

Example 1:
> Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
> Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
 

Constraints:
> 1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.

---

# 思路
在字符串 haystack 以needle的长度为窗口滑动，每次比较窗口内的字符串是否等于 needle，如果相等则返回当前窗口的起始位置。

# 代码
```go
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

// ============================================================================================================
// 函数实现了 KMP 算法用于在 str 中查找 needle 字符串的位置
func strStr(str string, needle string) int {
	// needle 的最长公共前后缀数组
	next := LPS(needle)
	// start 用于标记 needle 中当前匹配的位置
	start := 0

	for i := 0; i < len(str); i++ {
		// 如果当前字符匹配
		if needle[start] == str[i] {
			if start == len(needle)-1 { // 如果已经匹配到 needle 的最后一个字符
				return i - start // 返回结果，也就是匹配的起始位置
			}
			start += 1 // 匹配成功时，移动 needle 的匹配位置
		} else {
			//*** 当前字符不匹配情况，和LPS函数类似处理
			for start > 0 && str[i] != needle[start] { // 使用 LPS 数组进行跳转，直到找到可能的匹配位置或回退到初始位置
				start = next[start-1]
			}
			// 如果匹配了，start进一位
			if str[i] == needle[start] {
				start += 1 // 匹配成功时，移动 needle 的匹配位置
			}
		}
	}
	return -1 // 如果没有找到匹配，返回 -1
}

// LPS 函数计算字符串 needle 的最长公共前后缀数组
func LPS(needle string) []int {

	prefixList := make([]int, len(needle)) // LPS数组
	length := 0                            // length 用于跟踪当前最长的前后缀长度

	// 从 needle 的第二个字符开始遍历
	for i := 1; i < len(needle); i++ {
		// *** 如果当前字符与前后缀的下一个字符不匹配
		for length > 0 && needle[length] != needle[i] {
			length = prefixList[length-1] // 回退到上一个有效的前后缀长度
		}

		// 如果找到匹配的前后缀字符
		if needle[length] == needle[i] {
			length += 1            // 增加当前的前后缀长度
			prefixList[i] = length // 将当前长度存储在 LPS 数组对应位置
		}
	}
	return prefixList // 返回构建的 LPS 数组
}

// ============================================================================================================
// 1. 遍历对比
func strStr1(haystack string, needle string) int {

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

// ============================================================================================================
// 2. 用标准库的strings.Index方法
// strings.Index 返回字符串在另一个字符串中第一次出现的索引，如果没有找到，返回-1
func strStrSample1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
```