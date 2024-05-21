package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var Line, staNums int
	// 从标准输入读取两个整数，分别表示线路数和每条线路的站点数
	fmt.Scanf("%d %d", &Line, &staNums)

	// 创建一个二维切片来存储每条线路的站点信息
	var Route [][]int = make([][]int, Line)
	for i := range Route {
		Route[i] = make([]int, staNums)
	}
	index := 0

	// 使用bufio.Scanner来读取多行输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()      // 读取一行文本
		arr := strings.Fields(str) // 将行文本分割成单词（这里指数字字符串）
		for i := range arr {
			Route[index][i], _ = strconv.Atoi(arr[i]) // 将字符串转换为整数并存储在Route中
		}
		index += 1

		Line -= 1
		if Line < 1 {
			break // 读取足够的线路数据后退出循环
		}
	}

	// 读取下一行，获取需要计算成本的路径数
	scanner.Scan()
	str := scanner.Text()
	passed, _ := strconv.Atoi(str)
	preSta := 0 // 初始站点设置为0
	cost := 0   // 初始成本为0

	// 继续读取后续行，每行包含线路和站点编号
	for scanner.Scan() {
		str := scanner.Text()
		arr := strings.Fields(str)

		line, _ := strconv.Atoi(arr[0])
		sta, _ := strconv.Atoi(arr[1])

		line -= 1 // 线路编号从0开始
		sta -= 1  // 站点编号从0开始

		// 计算从前一个站点到当前站点的成本
		num := Route[line][sta] - Route[line][preSta]
		num = abs(num) // 使用自定义的abs函数获取绝对值

		cost += num // 累加成本

		preSta = sta // 更新前一个站点编号

		passed -= 1
		if passed < 1 {
			break // 当所有的路径都计算完毕，退出循环
		}
	}

	// 输出总成本
	fmt.Println(cost)
}

// 自定义函数，计算整数的绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
