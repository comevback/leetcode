# 392. Is Subsequence（双指针）

Easy

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

 

Example 1:
> Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
> Input: s = "axc", t = "ahbgdc"
Output: false
 

Constraints:
> 0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
 
Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?

---

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

// 1. 双指针 O(n)-O(1)
// 定义两个指针 posS 和 posT 分别代表 s 和 t 的下标
// 遍历 t，每当 t[posT] == s[posS] 时，posS += 1，也就是每当在 t 中找到 s 的当前下标的字符时，s 的下标向后移动一位
// 如果在过程中，s 的下标移动到了末尾，说明 s 里的所有字符在 t 中都可以按顺序找到，返回 true
// 反之如果遍历完 t 后，s 的下标还没有移动到末尾，说明 s 中的字符没有在 t 中全部按顺序找到，返回 false
func isSubsequence(s string, t string) bool {

	// 边界条件，s 为空时，返回 true，因为 s 为空时，t 中一定包含 s
	if len(s) == 0 {
		return true
	}
	// 边界条件，s 长度大于 t 时，返回 false，因为 s 长度大于 t 时，s 一定不是 t 的子序列
	if len(s) > len(t) {
		return false
	}

	// 设置两个指针 posS 和 posT 分别代表 s 和 t 的下标
	var posS, posT int = 0, 0

	// 遍历 t
	for posT < len(t) {
		// 如果 t[posT] == s[posS]，说明找到了 s 的一个字符，s 的下标向后移动一位
		if s[posS] == t[posT] {
			posS += 1
			// 如果 s 的下标移动到了末尾，说明 s 里的所有字符在 t 中都可以按顺序找到，返回 true
			if posS == len(s) {
				return true
			}
		}
		// t 的下标向后移动一位
		posT += 1
	}

	// 遍历完 t 后，s 的下标还没有移动到末尾，说明 s 中的字符没有在 t 中全部按顺序找到，返回 false
	return false
}

// ===============================================================================================================================
// 2. other solution
func isSubsequence1(s string, t string) bool {
	sChars := []rune(s) // 将字符串 s 转换成 rune 列表，支持多字节字符
	str := t            // 将字符串 t 赋值给变量 str，用于后续操作

	for _, r := range sChars { // 遍历 s 中的每个字符
		i := strings.IndexRune(str, r) // 在 str 中查找字符 r 的索引位置 i
		if i == -1 {                   // 如果找不到字符，返回 false
			return false
		}
		str = str[i+1:] // 更新 str 为从找到的字符位置的下一个字符开始的子字符串
	}
	return true // 如果所有字符都能按顺序在 t 中找到，返回 true
}
```