package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	s := "3[a]2[bc]"
	res := getMatched(s)
	fmt.Println(res)
}

// 1. 栈
// 用一个栈维护重复次数，一个栈装字符
func decodeString(s string) string {
	var numStack []int // 数字栈，用于存储重复次数
	var stack []byte   // 字符栈，用于存储字符串

	var numIndex int             // 记录数字开始的索引
	var numRecoding bool = false // 标记是否正在记录数字

	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' { // 如果当前字符是数字
			if !numRecoding { // 如果没有在记录数字，设置开始记录的索引
				numIndex = i
				numRecoding = true
			}
		} else if s[i] == '[' && numRecoding { // 如果当前字符是 '[' 且正在记录数字
			num, _ := strconv.Atoi(s[numIndex:i]) // 将记录的数字部分转换为整数
			numStack = append(numStack, num)      // 将数字压入数字栈
			stack = append(stack, s[i])           // 将 '[' 压入字符栈
			numRecoding = false                   // 重置数字记录标志
		} else if s[i] == ']' { // 如果当前字符是 ']'
			repeatTimes := numStack[len(numStack)-1] // 获取最近的重复次数
			numStack = numStack[:len(numStack)-1]    // 将该次数从栈中弹出
			var str []byte

			// 从字符栈中找到最近的 '['
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == '[' {
					str = make([]byte, len(stack[i+1:])) // 创建新的切片来存储子字符串，注意这里不能直接 str = stack[i+1:]，因为切片会随原本数组改变。
					copy(str, stack[i+1:])               // 复制子字符串到新的切片
					stack = stack[:i]                    // 将字符栈截断到 '[' 之前
					break
				}
			}

			// 将子字符串重复指定次数并压入字符栈
			for i := 0; i < repeatTimes; i++ {
				stack = append(stack, str...)
			}
		} else { // 如果当前字符是普通字符
			stack = append(stack, s[i]) // 直接将字符压入字符栈
		}
	}

	// 将字符栈中的字符组合成最终字符串
	var builder strings.Builder
	for _, v := range stack {
		builder.WriteByte(v)
	}
	res := builder.String()
	return res
}

// ==========================================================================================================
// 2. 递归
func getMatched(s string) []string {
	re := regexp.MustCompile(`\d+\[\w+\]`)
	finded := re.FindAllString(s, -1)
	return finded
}
