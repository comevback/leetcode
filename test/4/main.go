package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := F(5)
	fmt.Println(res)
}

func Solution(N int) int {
	str := strconv.Itoa(N)
	length := len(str)

	if length == 1 {
		return 1
	}

	haveZero := false

	allTimes := F(length)
	zeroStarts := 0

	hsMap := make(map[rune]int)
	for _, v := range str {
		hsMap[v] += 1
	}

	for key, value := range hsMap {
		if key == '0' {
			haveZero = true
		}
		allTimes /= F(value)
	}

	if haveZero {
		zeroStarts = F(length - 1)
		hsMap['0'] -= 1
		for _, value := range hsMap {
			zeroStarts /= F(value)
		}
	}

	return allTimes - zeroStarts
}

func F(num int) int {
	if num == 0 {
		panic("zero")
	}
	if num == 1 {
		return 1
	} else {
		return num * F(num-1)
	}
}
