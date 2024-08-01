# 290. Word Pattern

Solved
Easy
Topics
Companies
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false

Constraints:

1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.

---

# Code

```go
package main

import "strings"

func main() {
	// 主函数，程序的入口点。这里没有具体实现，因为主要逻辑在 wordPattern 函数中。
}

// wordPattern 检查字符串 s 是否按照给定的模式 pattern 来组织。
// pattern 是一个字符串，其中的每个字符表示一个单词的位置。
// s 是一个用空格分隔的单词序列。
// 返回一个布尔值，指示字符串 s 是否符合模式 pattern。
func wordPattern(pattern string, s string) bool {
	strArr := strings.Split(s, " ") // 将字符串 s 按空格分割成单词数组。
	hsMap := make(map[byte]string)  // 创建一个映射，从 pattern 的字符映射到对应的单词。
	hsMap2 := make(map[string]byte) // 创建另一个映射，从单词映射回 pattern 的字符。

	// 如果分割后的单词数量与模式中的字符数量不相等，直接返回 false。
	if len(strArr) != len(pattern) {
		return false
	}

	// 遍历 pattern 中的每个字符。
	for i := 0; i < len(pattern); i++ {
		_, exist := hsMap[pattern[i]]  // 检查当前模式字符是否已有映射的单词。
		_, exist2 := hsMap2[strArr[i]] // 检查当前单词是否已有映射的模式字符。

		// 如果两者都未映射，则互相映射。
		if !exist && !exist2 {
			hsMap[pattern[i]] = strArr[i]
			hsMap2[strArr[i]] = pattern[i]
		} else {
			// 如果已映射，检查当前字符和单词的映射是否一致。
			if !(hsMap[pattern[i]] == strArr[i] && hsMap2[strArr[i]] == pattern[i]) {
				return false // 如果不一致，返回 false。
			}
		}
	}

	return true // 如果所有字符和单词都一致映射，返回 true。
}
```
