# bufio.Scanner 自定义分割

一般用bufio的scanner持续读取输入时，可以使用预定义的分隔符，如bufio.ScanLines、bufio.ScanWords、bufio.ScanRunes等。但是如果输入的数据是自定义的分隔符，就需要自定义分隔符了。



### 标准分隔符
```go
scanner.Split(bufio.ScanLines)  // 按行分割
scanner.Split(bufio.ScanWords)  // 按单词分割
scanner.Split(bufio.ScanRunes)  // 按字符分割
scanner.Split(bufio.ScanBytes)  // 按字节分割
```

### 自定义分隔符
使用bufio.Scanner的Split方法，传入一个自定义的分隔符函数，该函数的签名为：
```go
func(data []byte, atEOF bool) (advance int, token []byte, err error)
```
- data: 输入的待处理数据。
- atEOF: 标志表示是否已到达数据的末尾。
- **advance**: 表示切掉多少字节。
- **token**: 表示返回的下一个令牌（如果有的话），返回的不一定等于切掉的全部字节。
- err: 处理中遇到的任何错误。

#### 例子：用逗号分隔
```go
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
```