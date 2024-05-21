package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var workNums int
	fmt.Scan(&workNums) // 从标准输入读取工作区间的数量

	// 创建一个大数组，用于标记每一天是否是工作日
	var arrWorkDays []int = make([]int, 100001)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()      // 读取一行输入
		arr := strings.Fields(str) // 将输入行按空格分割

		start, _ := strconv.Atoi(arr[0]) // 将开始日期从字符串转换为整数
		end, _ := strconv.Atoi(arr[1])   // 将结束日期从字符串转换为整数

		for i := start; i <= end; i++ {
			arrWorkDays[i] = 1 // 将开始到结束日期间的所有天标记为工作日
		}

		workNums -= 1
		if workNums < 1 {
			break // 当所有的工作区间都读取完毕后退出循环
		}
	}

	continues := 0 // 用于计算当前连续工作日的长度
	max := 0       // 用于存储最长连续工作日的长度
	for i := 1; i < len(arrWorkDays); i++ {
		if arrWorkDays[i] == 1 {
			continues += 1 // 如果当前天是工作日，则连续天数加一
		} else {
			if continues > max {
				max = continues // 如果当前连续天数大于已记录的最大值，则更新最大值
			}
			continues = 0 // 重置连续工作日计数
		}

		// 处理最后一段连续的工作日，确保在数组结束时也能正确比较
		if continues > max {
			max = continues
		}
	}

	fmt.Println(max) // 输出最长连续工作日的长度
}
