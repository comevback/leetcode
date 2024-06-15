# KMP算法和LPS数组

<https://yeefun.github.io/kmp-algorithm-for-beginners/>

KMP算法是一种用于在字符串中查找子串的算法。它的核心是构建一个最长公共前后缀数组（LPS）。这个数组用于在匹配失败时，跳过一些字符，从而提高查找效率。

---

[problem 28](../../Code/28-Find-the-Index-of-the-First-Occurrence-in-a-String/28-Find-the-Index-of-the-First-Occurrence-in-a-String.md)

# Code

```go
package main

import "fmt"

func main() {
	str := "mississippi"
	res := KMP(str, "issip")
	fmt.Println(res)
}

// 函数实现了 KMP 算法用于在 str 中查找 needle 字符串的位置
func KMP(str string, needle string) int {
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
			// *** 当前字符不匹配情况，和LPS函数类似处理
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
```