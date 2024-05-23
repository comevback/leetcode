# Bubble Sort 
### O(n^2) - O(1)

![Bubble-Sort](Bubble-Sort.gif)

冒泡排序每次都从头开始，和下一位比较大小，如果比对方大（小），则交换位置，一直到最里面，然后第二次循环，就从头到倒数第二个位置，以此类推

```go
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
```