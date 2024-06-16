package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://leetcode.com/problems/peak-index-in-a-mountain-array/description/"

	// 发起 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 输出获取的内容
	fmt.Println(string(body))
}
