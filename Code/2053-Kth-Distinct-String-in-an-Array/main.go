package main

import "fmt"

func main() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	res := kthDistinct(arr, 2)
	fmt.Println(res)
}

// kthDistinct 函数返回数组中第 k 个唯一字符串
func kthDistinct(arr []string, k int) string {
	// 用于存储唯一字符串的结果列表
	res := []string{}
	// 用于记录字符串出现次数的哈希映射
	hsMap := make(map[string]int)

	// 遍历数组，统计每个字符串的出现次数
	for _, v := range arr {
		hsMap[v] += 1
	}

	// 再次遍历数组，找到唯一的字符串并按顺序存储到结果列表中
	for _, v := range arr {
		// 如果字符串的出现次数为1，表示唯一
		if hsMap[v] == 1 {
			res = append(res, v)
			// 当结果列表中的唯一字符串个数等于k时，返回该字符串
			if len(res) == k {
				return v
			}
		}
	}

	// 如果没有找到第k个唯一字符串，返回空字符串
	return ""
}
