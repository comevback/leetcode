package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 定义项目数量和每一项的长度
	var H, N int
	fmt.Scanf("%d %d", &H, &N)
	scanner := bufio.NewScanner(os.Stdin)
	// 定义结果数组，为一个 string 数组
	res := make([]string, 0, H)

	// 循环读取输入
	for scanner.Scan() {
		// 读取一行输入
		str := scanner.Text()
		// 将输入按空格分割为一个字符串数组
		arr := strings.Split(str, " ")

		// 定义一个 int 数组，长度为 N
		numArr := make([]int, 0, N)
		// 遍历字符串数组，将每个字符串转换为 int 类型，如果大于等于 128 则转换为 1，否则转换为 0，然后添加到 int 数组中
		for _, v := range arr {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			if num >= 128 {
				num = 1
			} else {
				num = 0
			}
			numArr = append(numArr, num)
		}

		// 定义一个字符串，用于存储这一次的结果
		resString := ""
		// 遍历 int 数组，将每个元素转换为字符串，然后拼接到结果字符串中，如果是最后一个元素则不加空格
		for i, v := range numArr {
			var str string
			if i == len(numArr)-1 {
				str = strconv.Itoa(v)
			} else {
				str = strconv.Itoa(v) + " "
			}
			resString += str
		}
		// 将结果字符串添加到结果数组中
		res = append(res, resString)

		// H 减 1，如果 H 小于 1 则退出循环
		H -= 1
		if H < 1 {
			break
		}
	}

	// 遍历结果数组，输出每个结果
	for _, v := range res {
		fmt.Println(v)
	}
}
