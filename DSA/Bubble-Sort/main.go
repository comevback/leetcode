package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 5, 6, 2, 7, 9, 5, 3, 1}
	BubbleSort(nums)
	fmt.Println(nums)
}

// BubbleSort 冒泡排序 O(n^2) - O(1)

func BubbleSort(nums []int) {
	// 两层循环，每次都把最大的数放到最后
	for i := 0; i < len(nums)-1; i++ {
		// 每次都把最大的数放到最后
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
