package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义一个逆波兰表达式的字符串切片
	tokens := []string{"2", "1", "+", "3", "*"}
	// 调用 evalRPN 函数计算表达式的值
	res := evalRPN(tokens)
	// 打印计算结果
	fmt.Println(res)
}

// evalRPN 函数通过栈计算逆波兰表达式的值
func evalRPN(tokens []string) int {
	// 创建一个字符串切片，用作栈
	stack := []string{}

	// 遍历 tokens 中的每一个字符串
	for _, v := range tokens {
		// 如果当前字符串是运算符
		if v == "+" || v == "-" || v == "*" || v == "/" {
			// 从栈中弹出两个操作数
			left, _ := strconv.Atoi(stack[len(stack)-2])
			right, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			var res int
			// 根据运算符进行相应的计算
			switch v {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "*":
				res = left * right
			case "/":
				res = left / right
			}
			// 将计算结果转换为字符串并压入栈中
			resStr := strconv.Itoa(res)
			stack = append(stack, resStr)
		} else {
			// 如果当前字符串是数字，则直接压入栈中
			stack = append(stack, v)
		}
	}

	// 栈中最后一个元素即为计算结果，转换为整数并返回
	res, _ := strconv.Atoi(stack[0])
	return res
}

// evalRPN1 函数通过栈计算逆波兰表达式的值（优化版本，使用整数栈）
func evalRPN1(tokens []string) int {
	// 创建一个整数切片，用作栈
	stack := []int{}

	// 遍历 tokens 中的每一个字符串
	for _, v := range tokens {
		// 如果当前字符串是运算符
		if v == "+" || v == "-" || v == "*" || v == "/" {
			// 从栈中弹出两个操作数
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var res int
			// 根据运算符进行相应的计算
			switch v {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "*":
				res = left * right
			case "/":
				res = left / right
			}
			// 将计算结果压入栈中
			stack = append(stack, res)
		} else {
			// 如果当前字符串是数字，则转换为整数后压入栈中
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	// 栈中最后一个元素即为计算结果，返回结果
	res := stack[0]
	return res
}
