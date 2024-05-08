# 242. Valid Anagram

**类似题目：[389-Find-the-Difference.md](../../Memos/389-Find-the-Difference.md)**

Easy

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:
> Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
> Input: s = "rat", t = "car"
Output: false
 

Constraints:
> 1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
 

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

---

# 思路
这道题与[389-Find-the-Difference.md](../../Memos/389-Find-the-Difference.md)有所不同，这道题中，两个字符串可以由
> s = "ac"， t = "bb"

或

> s = "aa"， t = "bb"

组成，如果是这种情况，异或和求ASCII码的方法就不适用了。
只能用map和Array来统计两个字符串中每个字符出现的次数。

# 代码
```go
package main

func main() {

}

// 1.异或
// func isAnagram(s string, t string) bool {
// 	var XOR rune
// 	for _, v := range s + t {
// 		XOR ^= v
// 	}

// 	if XOR == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// 如果遇到 s = "aa"， t = "bb"这种情况，就会判断错误

// 2.map
func isAnagram(s string, t string) bool {
	hsMap := make(map[rune]int)

	for _, v := range s {
		hsMap[v] += 1
	}

	for _, v := range t {
		hsMap[v] -= 1
	}

	for _, val := range hsMap {
		if val != 0 {
			return false
		}
	}

	return true
}

// 3.ASCII值相差
// func isAnagram1(s string, t string) bool {
// 	sum := 0
// 	for _, v := range s {
// 		sum += int(v)
// 	}

// 	for _, v := range t {
// 		sum -= int(v)
// 	}

// 	if sum == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// 如果遇到 s = "ac"， t = "bb"这种情况，就会判断错误

// 4. Array代替map
func isAnagram1(s string, t string) bool {
	chars := make([]int, 26)

	for _, v := range s {
		i := int(v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := int(v - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
```
