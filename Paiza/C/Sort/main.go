package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)      // 读取用户输入的行数N
	var res *ListNode // 声明链表的头节点

	scanner := bufio.NewScanner(os.Stdin) // 创建从标准输入读取的Scanner
	for scanner.Scan() {                  // 循环读取每一行输入
		str := scanner.Text() // 获取输入的字符串
		newNode := &ListNode{ // 创建一个新的链表节点
			Val: str,
		}
		thisAmount := count(str) // 计算当前字符串的权重

		current := res      // 从链表头开始遍历
		if current == nil { // 如果链表为空，则直接插入新节点
			res = newNode
			N -= 1
			if N < 1 { // 如果已达到输入的行数，终止循环
				break
			}
			continue
		}

		var parentNode *ListNode // 用于记录当前节点的前一个节点
		for current != nil {
			Amount := count(current.Val) // 计算当前遍历到的节点的权重
			// 比较权重，决定插入位置
			if thisAmount[1] > Amount[1] || thisAmount[1] == Amount[1] && thisAmount[0] > Amount[0] {
				break
			}
			parentNode = current // 向前移动
			current = current.Next
		}

		// 插入新节点
		if parentNode == nil {
			newNode.Next = res
			res = newNode
		} else {
			newNode.Next = current
			parentNode.Next = newNode
		}

		N -= 1
		if N < 1 {
			break
		}
	}

	current := res // 从头节点开始输出链表
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// 把字符串转换为一个两个元素的整数型切片，方便比较
func count(str string) []int {
	fort := strings.Split(str, " ") // 将字符串按空格分割
	res := []int{}
	for _, v := range fort {
		num, err := strconv.Atoi(v) // 将字符串转换为整数
		if err != nil {
			panic(err) // 转换错误则抛出异常
		}
		res = append(res, num) // 将转换后的整数加入数组
	}

	return res
}

// 链表定义
type ListNode struct {
	Val  string
	Next *ListNode
}
