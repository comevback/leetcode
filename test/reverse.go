package main

import "fmt"

func main_reverse() {
	var str string = "abcd"
	res := reverse1(str)
	fmt.Println(res)
}

// 1.拆分成rune字符然后赋给另外一个rune数组，最后转换成string，返回
func reverse(str string) string {
	length := len(str)
	var r []rune = make([]rune, length)
	for i, v := range str {
		r[length-1-i] = v
	}

	result := string(r)
	return result
}

// 2. 类似1
func reverse1(str string) string {
	length := len(str)
	r := []rune(str)
	fmt.Println(r)
	// 这里一定是i < length/2而不是i <= length/2
	for i := 0; i < length/2; i++ {
		r[i], r[length-1-i] = r[length-1-i], r[i]
	}
	fmt.Println(r)
	result := string(r)
	return result
}
