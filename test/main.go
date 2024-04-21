package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 自分の得意な言語で
	// スキルチェックの基本となる、標準入力で値を取得し、
	// 出力するコードを書いてみよう！
	fmt.Println("enter anything")

	// 这是用fmt的Scan方法的输入 ====
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(input)

	// =============================

	// 用bufio

	// 创建一个输入检测对象scanner
	//scanner := bufio.NewScanner(os.Stdin)

	// 单次输入

	//
	// scanner.Scan()
	// str := scanner.Text()
	// fmt.Println(str)
	scanner := bufio.NewScanner(os.Stdin)

	// 每一个输入，都输出对应的值
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}
		fmt.Print("your input is: ")
		fmt.Println(input)
	}
	//
}
