package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "edebbed"                   // 输入的字符串
	res := removeDuplicateLetters(s) // 调用函数移除重复字母并保证字典序最小
	fmt.Println(res)                 // 打印结果
}

// removeDuplicateLetters 移除字符串中的重复字母并返回字典序最小的结果
func removeDuplicateLetters(s string) string {
	hsMap := make(map[rune]int)    // 用于记录每个字符在字符串中出现的次数
	stack := []rune{}              // 栈用来构建结果字符串
	instack := make(map[rune]bool) // 记录字符是否已经在栈（结果字符串）中

	// 计算每个字符出现的次数
	for _, v := range s {
		hsMap[v] += 1
	}

	// 遍历字符串中的每个字符
	for _, v := range s {
		if instack[v] { // 如果字符已在栈中，跳过，并减少计数
			hsMap[v] -= 1
			continue
		}

		// 栈非空且栈顶字符大于当前字符，且栈顶字符后面还会再出现
		for len(stack) > 0 && stack[len(stack)-1] > v && hsMap[stack[len(stack)-1]] > 0 {
			instack[stack[len(stack)-1]] = false // 标记栈顶字符不在栈中
			stack = stack[:len(stack)-1]         // 弹出栈顶字符
		}

		// 当前字符入栈，并更新状态
		stack = append(stack, v)
		hsMap[v] -= 1
		instack[v] = true
	}

	// 使用字符串构建器来组合最终的结果
	var builder strings.Builder
	for _, v := range stack {
		builder.WriteRune(v)
	}

	return builder.String() // 返回构建的字符串
}
