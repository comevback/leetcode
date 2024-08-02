package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	nums := 0
	res := reverse(nums)
	fmt.Println(res)
}

// reverse takes an integer and returns its reverse, accounting for overflow.
func reverse(x int) int {
	minus := false // 标志位，用于记录数字是否为负数
	if x < 0 {
		minus = true // 如果x是负数，设置minus为true
		x = -x       // 将x转为正数，简化反转过程
	}

	numStr := strconv.Itoa(x) // 将整数x转换为字符串
	numArr := []rune(numStr)  // 将字符串转换为rune切片，以支持可能的多字节字符操作

	// 反转rune切片中的元素
	reverseArr(numArr)

	// 移除反转后前导的'0'字符，除非这是唯一一个字符
	for len(numArr) > 1 && numArr[0] == '0' {
		numArr = numArr[1:]
	}

	newStr := string(numArr)               // 将处理后的rune切片转回字符串
	reversedNum, _ := strconv.Atoi(newStr) // 将字符串转换回整数，忽略转换错误（因为输入保证是数字）

	if minus {
		reversedNum = -reversedNum // 如果原始数字是负数，将结果转换为负数
	}

	// 检查反转后的数字是否溢出32位整数范围
	if reversedNum < math.MinInt32 || reversedNum > math.MaxInt32 {
		return 0 // 如果溢出，返回0
	}

	return reversedNum // 返回反转后的数字
}

// reverseArr reverses an array of runes in place.
func reverseArr(arr []rune) {
	if len(arr) < 2 {
		return // 如果数组长度小于2，不需要反转
	}

	head, tail := 0, len(arr)-1 // 使用头尾指针进行反转
	for head < tail {
		arr[head], arr[tail] = arr[tail], arr[head] // 交换头尾指针指向的元素
		head += 1
		tail -= 1
	}
}
