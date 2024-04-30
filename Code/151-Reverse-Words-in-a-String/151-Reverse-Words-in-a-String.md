# 151. Reverse Words in a String

Medium

Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

 

Example 1:
> Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
> Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
> Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
 

Constraints:
> 1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?

---

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "a good   example"
	res := reverseWords1(str)
	fmt.Println(res)
}

// 1.用strings.Fields分割为一个单词的数组 O(n)-O(n)
// 然后从后往前遍历，生产新字符串
func reverseWords(s string) string {
	wordArr := strings.Fields(s)
	res := ""

	for i := len(wordArr) - 1; i >= 0; i-- {
		if i == 0 {
			res += wordArr[0]
		} else {
			res += wordArr[i] + " "
		}
	}
	return res
}

// 2.不用strings.Fields O(n)-O(n)
// 从后往前遍历，遇到单词就记录尾部，遇到空格就记录头部，然后拼接
// 注意最后一个单词的处理，因为最后一个单词后面没有空格，所以要单独处理
func reverseWords1(s string) string {
	// 去掉首空格
	s = strings.TrimLeft(s, " ")
	// 记录去掉首空格后的长度
	length := len(s)

	// 记录单词的尾部的指针
	tail := length - 1
	// 是否遇到单词的信号
	meetWord := false
	res := ""

	// 从后往前遍历
	for i := length - 1; i >= 0; i-- {
		// 如果遇到单词，记录尾部，tail指向单词的最后一个字符，meetWord信号变为true
		if s[i] != ' ' && !meetWord {
			tail = i
			meetWord = true
		}
		// 如果遇到空格，而且meetWord信号为true，说明这个单词结束了到头了，拼接到res中，tail指向空格的前一个字符，meetWord信号变为false
		if (s[i] == ' ') && meetWord {
			res += s[i+1:tail+1] + " "
			tail = i
			meetWord = false
		}

		// 如果遍历到头，而且meetWord信号为true，说明这是第一个单词，拼接到res中，而且后面不加空格
		if i == 0 && meetWord {
			res += s[i : tail+1]
		}
	}

	return res
}

// 3.两次翻转
// 先把整个字符串翻转，然后依次翻转每个单词
// func reverseWords2(s string) string {

// }

// func resOneWord(str string) string {

// }

// 4.把每个单词加入栈，先进后出
```