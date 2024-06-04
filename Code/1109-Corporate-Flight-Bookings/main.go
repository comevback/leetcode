package main

import "fmt"

func main() {
	res := corpFlightBookings([][]int{{3, 3, 5}, {1, 3, 20}, {1, 2, 15}}, 3)
	fmt.Println(res)
}

func corpFlightBookings(bookings [][]int, n int) []int {
	// 初始化一个长度为n的差分数组
	diff := make([]int, n)

	// 遍历每个预订记录，并更新差分数组
	for _, v := range bookings {
		// 将预订人数加到起始位置（v[0]-1）
		diff[v[0]-1] += v[2]
		// 如果结束位置在数组范围内，将预订人数从结束位置+1处减去
		if v[1] < len(diff) {
			diff[v[1]] -= v[2]
		}
	}

	// 初始化结果数组
	res := make([]int, len(diff))
	// 第一个位置直接等于差分数组的第一个位置
	res[0] = diff[0]
	// 通过累加差分数组计算原数组的每个位置的值
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}

	// 返回结果数组
	return res
}
