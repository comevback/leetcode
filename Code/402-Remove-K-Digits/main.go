package main

import (
	"fmt"
	"strings"
)

func main() {
	num := "10"
	res := removeKdigits(num, 1)
	fmt.Println(res)
}

// removeKdigits 函数接受一个数字字符串和一个整数k，返回移除k个数字后的最小字符串形式的数字。
func removeKdigits(num string, k int) string {
	stack := []byte{} // 使用 byte 类型的 slice 作为栈，存储最终保留的数字

	for i := range num { // 遍历数字字符串的每一个字符
		// 如果还有可移除的数字（k>0），并且栈非空，且栈顶元素大于当前字符，执行循环
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1] // 移除栈顶元素
			k -= 1                       // 已移除一个数字，k减1
		}
		stack = append(stack, num[i]) // 将当前字符压入栈
	}

	stack = stack[:len(stack)-k] // 如果遍历结束后k仍然大于0，从栈顶移除剩余的k个元素

	res := ""                   // 初始化结果字符串
	var builder strings.Builder // 使用字符串构建器来拼接结果，提高性能
	for _, v := range stack {
		builder.WriteByte(v) // 将栈中的元素逐个添加到字符串构建器中
	}

	res = builder.String() // 从构建器获取最终字符串
	if res == "" {         // 如果结果为空（所有字符都被移除），返回"0"
		return "0"
	}
	// 移除字符串开头的所有前导零
	for res[0] == '0' && len(res) > 1 {
		res = res[1:]
	}
	return res // 返回处理后的字符串
}
