# 2390. Removing Stars From a String

Medium

Hint
You are given a string s, which contains stars *.

In one operation, you can:

Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.


Example 1:
> Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".

Example 2:
> Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.
 

Constraints:
> 1 <= s.length <= 105
s consists of lowercase English letters and stars *.
The operation above can be performed on s.


---

# Code
```go
package main

import (
	"errors"
	"strings"
)

func main() {

}

// 1. 我的方法，用链表实现Stack
func removeStars(s string) string {
	// 创造一个栈实例
	stack := NewStack()
	// 创造一个逆序栈实例，用于存储最终结果
	reverseStack := NewStack()
	// 创建一个字符串构建器
	var builder strings.Builder

	// 遍历字符串，如果是星号则弹出栈顶元素，否则压入栈
	for _, v := range s {
		if v != '*' {
			stack.push(v)
		} else {
			stack.pop()
		}
	}

	// 将栈中元素逆序压入逆序栈
	for stack.Top != nil {
		char, err := stack.pop()
		if err != nil {
			panic(err)
		}
		reverseStack.push(char.Val)
	}

	// 将逆序栈中元素压入构建器
	for reverseStack.Top != nil {
		char, err := reverseStack.pop()
		if err != nil {
			panic(err)
		}
		builder.WriteRune(char.Val)
	}

	// 将构建器转换为字符串，返回
	res := builder.String()
	return res
}

// 链表节点定义
type LinkedNode struct {
	Val  rune
	Next *LinkedNode
}

// 栈定义
type Stack struct {
	Top *LinkedNode
}

// 栈的构造函数
func NewStack() *Stack {
	return &Stack{}
}

// 栈的push方法
func (stack *Stack) push(value rune) {
	newNode := &LinkedNode{
		Val:  value,
		Next: stack.Top,
	}

	stack.Top = newNode
}

// 栈的pop方法
func (stack *Stack) pop() (*LinkedNode, error) {
	if stack.Top == nil {
		return nil, errors.New("empty stack")
	}

	res := stack.Top
	stack.Top = stack.Top.Next
	return res, nil
}

// ===================================================================================================================
// 2. 数组实现Stack
func removeStars1(s string) string {
	stack := []rune{}
	var builder strings.Builder

	for _, v := range s {
		if v != '*' {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	for _, v := range stack {
		builder.WriteRune(v)
	}

	res := builder.String()
	return res
}
```