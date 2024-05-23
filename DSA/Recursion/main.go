package main

import "fmt"

func main() {
	res := ReverseString("hello")
	fmt.Println(res)
}

// 阶乘
func Factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return Factorial(value-1) * value
	}
}

// 斐波那契数列
func Fibonacci(value int) int {
	if value <= 1 {
		return value
	}
	return Fibonacci(value-1) + Fibonacci(value-2)
}

// 递归翻转字符串
func ReverseString(str string) string {
	if len(str) == 1 {
		return str
	} else {
		return ReverseString(str[1:]) + string(str[0])
	}
}
