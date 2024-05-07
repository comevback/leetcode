package main

import (
	"bufio"
	"fmt"
	"os"
)

// 定义一个结构体mapNode，用来存储网格信息
type mapNode struct {
	nodeMap [][]int // 二维切片，存储网格中的每个点的状态
}

// NewMapNode函数，用于创建并初始化一个mapNode结构体实例
func NewMapNode(w int, h int) *mapNode {
	newNode := make([][]int, w) // 创建一个二维切片，宽度为w

	for i := range newNode {
		newNode[i] = make([]int, h) // 对于每一行，分配长度为h的切片
	}

	return &mapNode{
		nodeMap: newNode, // 将初始化好的二维切片赋给mapNode的nodeMap字段
	}
}

// add方法，用于在指定位置添加一个点的状态
func (m *mapNode) add(w int, h int) {
	m.nodeMap[w][h] = 1 // 在网格的w, h位置标记为1
}

func main() {
	var height, width int
	fmt.Scanf("%d %d", &height, &width) // 从标准输入读取网格的高度和宽度
	Nmap := NewMapNode(width, height)   // 创建一个新的网格节点
	h := height - 1                     // 因为需要从底部开始填充网格，设置h为height-1
	var dountCount int = 0              // 初始化甜甜圈计数器

	scanner := bufio.NewScanner(os.Stdin) // 创建一个新的扫描器，用于读取输入行
	buf := make([]byte, 0, 64*1024)       // 分配一个初始大小为64KB的缓冲区
	scanner.Buffer(buf, 1024*1024)        // 设置扫描器的最大容量为1MB
	for scanner.Scan() {                  // 循环读取每一行输入
		str := scanner.Text()        // 获取扫描到的文本
		for i := 0; i < width; i++ { // 遍历每一行的每一个字符
			if str[i] == '.' { // 如果字符是'.'
				Nmap.add(i, h) // 在对应的网格位置标记1
			}
		}

		h -= 1     // 准备填充下一行，行索引递减
		if h < 0 { // 如果行索引小于0，说明所有行都已处理完毕
			break // 退出循环
		}
	}

	// 检查网格中的每一个可能的3x3区域
	for i := 1; i < width-1; i++ {
		for j := 1; j < height-1; j++ {
			// 判断中心点及其周围的8个点是否符合甜甜圈图案的条件
			if Nmap.nodeMap[i][j] == 1 && Nmap.nodeMap[i-1][j-1] == 0 && Nmap.nodeMap[i-1][j] == 0 && Nmap.nodeMap[i-1][j+1] == 0 &&
				Nmap.nodeMap[i][j-1] == 0 && Nmap.nodeMap[i][j+1] == 0 &&
				Nmap.nodeMap[i+1][j-1] == 0 && Nmap.nodeMap[i+1][j] == 0 && Nmap.nodeMap[i+1][j+1] == 0 {
				dountCount += 1 // 如果符合，甜甜圈计数加一
			}
		}
	}

	fmt.Println(dountCount) // 输出甜甜圈的总数
}
