package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr1))
}

// func runningSum(nums []int) []int {

// }

func sum(nums []int) int {
	var res int = 0
	for _, num := range nums {
		res += num
	}
	return res
}
