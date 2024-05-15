package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{6, 10, 6, 9, 7, 8}
	res := Solution(arr)
	fmt.Println(res)
}

func Solution(A []int) int {
	sort.Ints(A)
	maxLength := 1
	length := 0
	base := A[0]

	for i := 0; i < len(A)-1; i++ {
		if A[i]-base <= 1 {
			length += 1
			if length > maxLength {
				maxLength = length
			}
		} else {
			base = A[i]
			length = 1
		}
	}

	return maxLength
}
