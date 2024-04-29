# 345. Reverse Vowels of a String（双指针）

Easy

Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

 

Example 1:
> Input: s = "hello"
Output: "holle"

Example 2:
> Input: s = "leetcode"
Output: "leotcede"
 

Constraints:
> 1 <= s.length <= 3 * 105
s consist of printable ASCII characters.

---

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))
}

// ===============================================================================================================================
// 1. 双指针 O(n)-O(n)
// 设定头尾两个指针，分别从头尾开始遍历，当头指针遇到元音字母时，停止，当尾指针遇到元音字母时，停止，交换两个指针的元素
func reverseVowels(s string) string {
	length := len(s)
	head := 0
	tail := length - 1
	target := "aeiouAEIOU"

	// 标记头尾指针是否找到元音字母，找到为true，未找到为false
	singalHead := false
	singalTail := false

	r := []rune(s)

	// 当头指针小于尾指针时，进行遍历
	for head < tail {
		// 当头尾指针都找到元音字母时，交换两个指针的元素
		if singalHead && singalTail {
			temp := r[head]
			r[head] = r[tail]
			r[tail] = temp
			singalHead = false
			singalTail = false
			head += 1
			tail -= 1
			continue
		}
		// 当头尾指针未找到元音字母时，继续遍历
		if !singalHead {
			if strings.ContainsRune(target, rune(s[head])) {
				singalHead = true
			} else {
				head += 1
			}
		}
		// 当头尾指针未找到元音字母时，继续遍历
		if !singalTail {
			if strings.ContainsRune(target, rune(s[tail])) {
				singalTail = true
			} else {
				tail -= 1
			}
		}
	}

	// 将rune转换为string，返回
	result := string(r)
	return result
}

// ===============================================================================================================================
// solution from [LeetCode The Hard Way] 我的解法的简化板 O(n)-O(n)
func reverseVowels1(s string) string {
	ss := []rune(s)
	n := len(ss)
	l, r := 0, n-1
	for l < r {
		// find the index of the first vowel in the given range
		for l < r && !isVowel(ss[l]) {
			l += 1
		}
		// find the index of last vowel in the given range
		for r > l && !isVowel(ss[r]) {
			r -= 1
		}
		// swap ss[l] and ss[r]
		ss[l], ss[r] = ss[r], ss[l]
		// since we've processed the character ss[l],
		// we move the left pointer to the right
		l += 1
		// since we've processed the character ss[r],
		// we move the right pointer to the left
		r -= 1
	}
	return string(ss)
}

func isVowel(c rune) bool {
	// alternatively, we can just check
	// return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' ||
	//        c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

```