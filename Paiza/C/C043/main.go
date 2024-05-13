package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var playerNum int
	fmt.Scan(&playerNum) // 读取第一行输入，表示后续要处理的数字个数

	maxTimes := 0              // 记录最大出现次数
	hsMap := make(map[int]int) // 使用哈希表存储每个数字出现的次数
	most := []int{}            // 存储出现次数最多的数字

	scanner := bufio.NewScanner(os.Stdin) // 创建一个从标准输入读取的scanner
	scanner.Split(bufio.ScanWords)        // 设置scanner的分割函数，按单词（数字）分割

	for scanner.Scan() {
		numStr := scanner.Text()         // 读取一个单词（数字的字符串形式）
		num, err := strconv.Atoi(numStr) // 将字符串转换为整数
		if err != nil {
			panic(err) // 如果转换出错，则抛出异常
		}

		hsMap[num] += 1 // 记录或更新该数字出现的次数
		if hsMap[num] > maxTimes {
			maxTimes = hsMap[num]    // 更新最大出现次数
			most = []int{}           // 重置存储最多次数数字的切片
			most = append(most, num) // 添加当前数字
		} else if hsMap[num] == maxTimes {
			exist := false // 检查该数字是否已经在最多次数列表中
			for _, v := range most {
				if v == num {
					exist = true
					break
				}
			}
			if !exist {
				most = append(most, num) // 如果不在列表中，则添加
			}
		}

		playerNum -= 1 // 减少待处理的数字计数
		if playerNum < 1 {
			break // 如果已处理完所有数字，跳出循环
		}
	}

	var resStr string = "" // 初始化结果字符串
	sort.Ints(most)        // 对出现次数最多的数字进行排序

	for i, v := range most {
		var str string
		if i == len(most)-1 {
			str = strconv.Itoa(v) // 最后一个数字后不加空格
		} else {
			str = strconv.Itoa(v) + " " // 其他数字后添加空格
		}
		resStr += str // 将数字字符串添加到结果字符串
	}

	fmt.Println(resStr) // 输出结果字符串
}
