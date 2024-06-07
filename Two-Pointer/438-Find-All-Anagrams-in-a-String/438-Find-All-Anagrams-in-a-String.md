# 438. Find All Anagrams in a String
Solved
Medium
Topics
Companies
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


Example 1:
> Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
> Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
 

Constraints:

1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.

---

# Code
```go
package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}

func findAnagrams(s string, p string) []int {
	// need 用来存储目标字符串 p 中每个字符的出现次数
	need := make(map[byte]int)
	// window 用来存储滑动窗口中每个字符的出现次数
	window := make(map[byte]int)
	// valid 表示滑动窗口中满足要求的字符数
	valid := 0
	// res 用来存储结果，即所有异位词的起始索引
	res := []int{}

	// 初始化 need 字典，记录 p 中每个字符的出现次数
	for i := 0; i < len(p); i++ {
		need[p[i]] += 1
	}

	// 滑动窗口的左右指针
	left, right := 0, 0

	// 开始滑动窗口
	for right < len(s) {
		// 判断当前字符是否是需要的字符
		if need[s[right]] > 0 {
			// 如果是，将其加入到窗口中，并更新计数
			window[s[right]] += 1
			// 如果窗口中的该字符数达到需求，更新 valid 计数
			if window[s[right]] == need[s[right]] {
				valid += 1
			}
			// 右指针右移，扩大窗口
			right += 1

			// 如果窗口大小超过了 p 的长度，需要缩小窗口
			if right-left > len(p) {
				// 判断左指针字符是否在 need 中
				if need[s[left]] > 0 {
					// 如果在，且窗口中的该字符数刚好满足需求，更新 valid 计数
					if window[s[left]] == need[s[left]] {
						valid -= 1
					}
					// 左指针字符移出窗口
					window[s[left]] -= 1
				}
				// 左指针右移，缩小窗口
				left += 1
			}

			// 如果窗口中满足需求的字符数与 need 中字符数相等，则记录起始索引
			if valid == len(need) {
				res = append(res, left)
			}
		} else {
			// 如果当前字符不是需要的字符，直接将右指针移到下一个字符
			right += 1
			// 左指针与右指针同步，并重置 valid 和 window
			left = right
			valid = 0
			window = make(map[byte]int)
		}
	}

	// 返回结果数组
	return res
}
```