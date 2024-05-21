package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cardType int
	fmt.Scan(&cardType)                             // 从标准输入读取卡片类型的数量
	times := cardType * 2                           // 计算需要读取的卡片总数（每种类型2张）
	index := 1                                      // 用于记录当前卡片的索引位置
	var cardMap map[int][]int = make(map[int][]int) // 创建一个映射来存储每种类型的卡片和它们的索引位置
	sum := 0                                        // 用于累加结果

	scanner := bufio.NewScanner(os.Stdin) // 创建一个新的Scanner对象
	for scanner.Scan() {                  // 循环读取标准输入
		str := scanner.Text()       // 获取当前行的文本
		num, _ := strconv.Atoi(str) // 将读取的文本转换为整数

		cardMap[num] = append(cardMap[num], index) // 将当前卡片的索引添加到对应类型的切片中

		index += 1 // 索引递增
		times -= 1 // 减少剩余需要读取的次数
		if times < 1 {
			break // 如果已读取足够的卡片，退出循环
		}
	}

	for _, value := range cardMap { // 遍历每种卡片类型的索引列表
		add := value[1] - value[0] - 1 // 计算两张同类型卡片之间的间隔数
		sum += add                     // 累加间隔数到总和
	}

	fmt.Println(sum) // 输出所有间隔数的总和
}
