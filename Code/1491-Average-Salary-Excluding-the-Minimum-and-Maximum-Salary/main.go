package main

import (
	"fmt"
	"math"
)

func main() {
	salary := []int{1000, 2000, 3000}
	res := average(salary)
	fmt.Println(res)
}

// 1. 我的解法，一次循环
// 累加每个整数，在循环中找出最大最小值
// 累加结束后，总和减去最大最小值，然后除以总数减二，得到结果
func average(salary []int) float64 {
	var highest, lowest int = math.MinInt, math.MaxInt
	var sum int
	var res float64

	for _, v := range salary {
		sum += v
		if v > highest {
			highest = v
		}
		if v < lowest {
			lowest = v
		}
	}

	sum = sum - highest - lowest
	res = float64(sum) / float64(len(salary)-2)
	return res
}
