package main

import (
	"fmt"
	"strings"
)

func main() {
	// 得到花环的长度
	var length int
	fmt.Scan(&length)

	// 得到两个花环
	var f1, f2 string
	fmt.Scan(&f1)
	fmt.Scan(&f2)

	// 将花环1转换为字符切片
	var arr1 []rune
	for _, v := range f1 {
		arr1 = append(arr1, v)
	}

	// 循环length次，每次将花环1旋转i步，然后与花环2比较，如果相等则输出Yes
	for i := 0; i < len(f1); i++ {
		temp := rotate(arr1, i)
		if temp == f2 {
			fmt.Println("Yes")
			return
		}
	}

	// 如果循环结束都没有找到相等的，则输出No
	fmt.Println("No")
}

// 旋转一个字符切片 step 步
func rotate(arr []rune, step int) string {
	temp := []rune{}
	temp = append(temp, arr...)

	// 整体旋转
	reverse(temp)
	// 旋转前 step 个字符
	reverse(temp[:step])
	// 旋转后 len(temp) - step 个字符
	reverse(temp[step:])

	// 将字符切片转换为字符串
	var builder strings.Builder
	for _, v := range temp {
		builder.WriteRune(v)
	}

	return builder.String()
}

// 翻转数组
func reverse(arr []rune) {
	var head, tail int = 0, len(arr) - 1
	for head < tail {
		arr[head], arr[tail] = arr[tail], arr[head]
		head += 1
		tail -= 1
	}
}
