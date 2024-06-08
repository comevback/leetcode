# 3. Longest Substring Without Repeating Characters
Solved
Medium
Topics
Companies
Hint
Given a string s, find the length of the longest 
substring without repeating characters.


Example 1:
> Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
> Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
> Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

---

# Code
```go
package main

import "fmt"

func main() {
	s := "abba"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	// charMap 是一个用来存储每个字符最后出现位置的映射。
	charMap := make(map[byte]int)
	maxlength := 0

	left, right := 0, 0

	for right < len(s) {
		// 如果字符 s[right] 不在映射中，说明它在当前窗口中没有重复。
		if _, exist := charMap[s[right]]; !exist {
			// 将字符 s[right] 及其位置添加到映射中。
			charMap[s[right]] = right
			// 计算当前窗口的长度。
			length := right - left + 1
			// 更新最长子串的长度。
			if length > maxlength {
				maxlength = length
			}
			// 移动右指针。
			right += 1
		} else {
			// 如果字符 s[right] 在映射中存在，说明在当前窗口中有重复字符。
			// 更新左指针以缩小窗口，这一步检查是为了防止 left 回退到之前的位置。
			if charMap[s[right]]+1 >= left {
				left = charMap[s[right]] + 1
			}
			// 更新字符 s[right] 的最新位置。
			charMap[s[right]] = right
			// 计算更新后的窗口长度。
			length := right - left + 1
			// 更新最长子串的长度。
			if length > maxlength {
				maxlength = length
			}
			// 移动右指针。
			right += 1
		}
	}

	// 返回找到的最长无重复字符子串的长度。
	return maxlength
}
```