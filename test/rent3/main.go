package main

import "fmt"

func main() {
	res := calPossibility()
	fmt.Println(res)
}

func calPossibility() int {
	num := 1
	var nowPoss float64 = 1.0
	var allDiffPoss float64 = 1.0
	for allDiffPoss > 0.5 {
		nowPoss = (365 - float64(num)) / 365
		allDiffPoss *= nowPoss
		num += 1
	}

	return num
}
