# 28. Find the Index of the First Occurrence in a String（KMP）

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

// 3.KMP算法 
// strings.Index 返回字符串在另一个字符串中第一次出现的索引，如果没有找到，返回-1
func strStrSample1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
```