package main

import (
	"container/list"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 输入字符串表示文件和目录结构
	str := "file1.txt\nfile2.txt\nlongfile.txt"
	// 调用 lengthLongestPath 函数计算最长路径的长度
	res := lengthLongestPath(str)
	// 打印计算结果
	fmt.Println(res)
}

// lengthLongestPath 函数计算给定输入字符串中最长的文件路径长度
func lengthLongestPath(input string) int {
	// 初始化最长路径长度为0
	maxLength := 0
	// 初始化当前深度为0
	currentDepth := 0
	// 创建一个映射用于记录每个深度的路径长度
	depthIndex := make(map[int]int)
	// 创建一个字节切片作为栈
	stack := []byte{}

	// 遍历输入字符串的每一个字符
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' { // 如果当前字符是换行符，计算深度
			depth := 0
			j := i + 1
			for j < len(input) && input[j] == '\t' {
				j += 1
				depth += 1
			}
			// 根据深度更新 depthIndex 和 currentDepth
			if depth > currentDepth {
				depthIndex[depth] = len(stack)
				currentDepth = depth
			} else {
				index := depthIndex[depth]
				stack = stack[:index]
				currentDepth = depth
			}
			// 如果深度不为0，将换行符加入栈中
			if depth != 0 {
				stack = append(stack, input[i])
			}
		} else if input[i] == '\t' { // 如果当前字符是制表符，跳过
			continue
		} else if input[i] == '.' { // 只有当出现了点号，才是一个文件，才需要计算路径长度，如果没有点号，表示是一个目录，不需要计算路径长度
			// 如果当前字符是点号，表示文件名
			stack = append(stack, input[i])
			j := i + 1
			length := 0

			// 计算文件后缀的长度
			for j < len(input) && input[j] != '\n' {
				j += 1
				length += 1
			}

			// 更新最长路径长度
			if len(stack)+length > maxLength {
				maxLength = len(stack) + length
			}
		} else { // 如果当前字符不是换行符、制表符和点号，将当前字符加入栈中
			stack = append(stack, input[i])
		}
	}

	// 返回最长路径长度
	return maxLength
}

// ***********************************************************************************************************************
// lengthLongestPath 函数计算给定输入字符串中最长的文件路径长度
func lengthLongestPath1(input string) int {
	// 创建一个链表作为栈，存储之前的父路径长度
	stack := list.New()
	maxLen := 0
	// 将输入字符串按换行符分割成多个部分
	parts := strings.Split(input, "\n")
	for _, part := range parts {
		// 计算当前部分的深度（级别）
		level := strings.LastIndexByte(part, '\t') + 1
		// 让栈中只保留当前目录的父路径
		for stack.Len() > level {
			stack.Remove(stack.Back())
		}
		// 将当前部分去掉前面的制表符，并压入栈中
		stack.PushBack(part[level:])
		// 如果是文件，计算路径长度
		if strings.Contains(part, ".") {
			sum := 0
			for e := stack.Front(); e != nil; e = e.Next() {
				sum += utf8.RuneCountInString(e.Value.(string))
			}
			// 加上父路径的分隔符
			sum += stack.Len() - 1
			// 更新最大长度
			maxLen = max(maxLen, sum)
		}
	}
	return maxLen
}

// max 函数返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
