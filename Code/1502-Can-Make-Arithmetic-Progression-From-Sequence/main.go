package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4}
	fmt.Println(canMakeArithmeticProgression(arr))
}

// 排序后遍历差值
func canMakeArithmeticProgression(arr []int) bool {
	// 从小到大排序
	sort.Ints(arr)
	// 初始差值为最后一个元素减倒数第二个元素
	diff := arr[len(arr)-1] - arr[len(arr)-2]
	// 遍历每两个相邻的差值，如果有不相等的，返回false
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	// 如果都想等，返回true
	return true
}
