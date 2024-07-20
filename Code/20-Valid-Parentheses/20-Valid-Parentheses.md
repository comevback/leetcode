# 20. Valid Parentheses

Solved
Easy
Topics
Companies
Hint
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

---

# Code

```go
package main

func main() {

}

// isValid 函数用于检查输入字符串 s 是否是有效的括号序列
func isValid(s string) bool {
	// 创建一个字节切片，用作栈
	stack := []byte{}

	// 遍历输入字符串中的每一个字符
	for i := 0; i < len(s); i++ {
		switch s[i] {
		// 如果当前字符是 ')'
		case ')':
			// 如果栈非空且栈顶元素是 '('，则弹出栈顶元素
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				// 否则将当前字符压入栈中
				stack = append(stack, s[i])
			}
		// 如果当前字符是 ']'
		case ']':
			// 如果栈非空且栈顶元素是 '['，则弹出栈顶元素
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				// 否则将当前字符压入栈中
				stack = append(stack, s[i])
			}
		// 如果当前字符是 '}'
		case '}':
			// 如果栈非空且栈顶元素是 '{'，则弹出栈顶元素
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				// 否则将当前字符压入栈中
				stack = append(stack, s[i])
			}
		// 对于其他字符（即 '(', '[', '{'），直接压入栈中
		default:
			stack = append(stack, s[i])
		}
	}

	// 如果栈为空，说明所有括号匹配，返回 true，否则返回 false
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
```
