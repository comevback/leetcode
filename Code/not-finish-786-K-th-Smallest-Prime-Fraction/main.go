package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5}
	fmt.Println(kthSmallestPrimeFraction(arr, 3))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	if k == 1 {
		return []int{arr[0], arr[len(arr)-1]}
	}

	for i := len(arr) - 1; i < 0; i-- {
		for j := 0; i < i; i++ {

		}
	}

}
