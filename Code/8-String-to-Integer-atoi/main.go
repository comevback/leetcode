package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	s := "   -042"
	// 调用 myAtoi 函数并打印结果
	res := myAtoi(s)
	fmt.Println(res)
}

// myAtoi 将字符串转换为整数
func myAtoi(s string) int {
	numStr := ""

	// 遍历字符串中的每个字符
	for _, v := range s {
		// 跳过前导空格
		if v == ' ' && numStr == "" {
			continue
			// 如果当前字符是数字或符号且 numStr 为空，则将其加入 numStr
		} else if v >= '0' && v <= '9' || (v == '-' || v == '+') && numStr == "" {
			numStr += string(v)
			// 遇到非数字字符后立即停止
		} else {
			break
		}
	}

	// 将 numStr 转换为整数
	num, _ := strconv.Atoi(numStr)

	// 检查整数是否超出 32 位整数范围
	if num < math.MinInt32 {
		num = math.MinInt32
	}
	if num > math.MaxInt32 {
		num = math.MaxInt32
	}

	return num
}

// for _, v := range s {

// 	if v == ' ' {
// 		if numStr == "" && !getSign {
// 			continue
// 		} else {
// 			break
// 		}
// 	} else if v == '-' || v == '+' {
// 		if numStr == "" && !getSign {
// 			if v == '-' {
// 				minus = true
// 			}
// 		} else {
// 			break
// 		}
// 		getSign = true
// 	} else if v < '0' || v > '9' {
// 		break
// 	} else {
// 		numStr += string(v)
// 	}
// }
