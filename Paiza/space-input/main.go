package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Scan(&num)
	var strArr []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(ComSplitter)

	for scanner.Scan() {
		// 获取不带逗号的字符串
		str := scanner.Text()[:len(scanner.Text())]
		// 字符串加入数组
		strArr = append(strArr, str)
		num -= 1
		if num < 1 {
			break
		}
	}

	// 打印数组里的元素
	for _, v := range strArr {
		fmt.Println(v)
	}
}

// 遇到逗号分隔
func ComSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return len(data), data, nil
	}

	// 在data里向后查找
	for i := 0; i < len(data); i++ {
		// 如果发现某个字节为逗号
		if data[i] == ',' {
			// advance = i + 1：切掉包括逗号之前的字符串
			// token = data[:i]：返回从头开始不包括逗号的字符串
			// err = nil：不返回错误
			return i + 1, data[:i], nil
		}
		if data[i] == '\n' {
			return len(data), data, nil
		}
	}
	return 0, nil, nil
}
