package main

import "fmt"

func main() {
	res := calProSum(3, 3) // テストの例
	fmt.Println(res)
}

func calProSum(m int, n int) int {
	if m < 0 || n < 0 {
		fmt.Println("invaild input")
		return -1
	}

	sum := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			temp := i * j
			sum += temp
		}
	}

	return sum
}
