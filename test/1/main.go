package main

import "fmt"

func main() {
	arr := []int{1, 4, -1, 3, 2}
	fmt.Println(Solution(arr))
}

func Solution(A []int) int {
	length := 1
	index := 0
	for A[index] != -1 {
		index = A[index]
		length += 1
	}

	return length
}
