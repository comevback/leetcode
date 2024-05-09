# 709. To Lower Case

Easy

Hint
Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.

Example 1:
> Input: s = "Hello"
Output: "hello"

Example 2:
> Input: s = "here"
Output: "here"

Example 3:
> Input: s = "LOVELY"
Output: "lovely"
 

Constraints:
> 1 <= s.length <= 100
s consists of printable ASCII characters.

---

# Code
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "Hello"
	fmt.Println(toLowerCase2(a))
}

// 1. 遍历每个字符，如果是大写字母范围的，转换为小写，加入新字符串，返回新字符串
func toLowerCase(s string) string {
	res := ""
	for _, v := range s {
		if int(v) >= 65 && int(v) <= 90 {
			res += string(rune(int(v) + 32))
		} else {
			res += string(v)
		}
	}
	return res
}

// 2. 用 strings.Builder 构造string
func toLowerCase1(s string) string {
	var builder strings.Builder
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			builder.WriteRune(v + 32)
		} else {
			builder.WriteRune(v)
		}
	}
	return builder.String()
}

// 3. 遍历每个字符，如果是大写字母范围的，转换为小写，加入新字符串，返回新字符串
func toLowerCase2(s string) string {
	var str string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			str += string(v + 32)
		} else {
			str += string(v)
		}
	}
	return str
}
```