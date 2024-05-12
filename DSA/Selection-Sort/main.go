package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 5, 6, 2, 7, 9, 5, 3, 1}
	SelectionSort(nums)
	fmt.Println(nums)
}

// SelectionSort 选择排序 O(n^2) - O(1)
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
